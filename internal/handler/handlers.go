package handler

import (
	"Chat/views/layouts"
	"fmt"
	"log"
	"net/http"
	"sort"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

const (
	pongWait       = 60 * time.Second
	maxMessageSize = 512
	rateLimit      = 1000 * time.Millisecond
)

var (
	wsChan            = make(chan WsPayload)
	clients           = make(map[WebSocketConnection]string)
	clientsMutex      sync.Mutex
	rateLimiterMutex  sync.Mutex
	lastMessageTime   = make(map[WebSocketConnection]time.Time)
	upgradeConnection = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func Home(w http.ResponseWriter, r *http.Request) {
	if err := layouts.App().Render(r.Context(), w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Chat(w http.ResponseWriter, r *http.Request) {
	if err := layouts.Chat().Render(r.Context(), w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Message(w http.ResponseWriter, r *http.Request) {
	sendMessage(w, r)
}

type WebSocketConnection struct {
	*websocket.Conn
}

type WsJsonResponse struct {
	Action         string   `json:"action"`
	Message        string   `json:"message"`
	MessageType    string   `json:"message_type"`
	ConnectedUsers []string `json:"connected_users"`
}

type WsPayload struct {
	Action   string              `json:"action"`
	Username string              `json:"username"`
	Message  string              `json:"message"`
	Conn     WebSocketConnection `json:"-"`
}

// WsEndpoint upgrades connection to websocket
func WsEndpoint(w http.ResponseWriter, r *http.Request) {
	ws, err := upgradeConnection.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}

	log.Println("Client Connected...")
	response := WsJsonResponse{Message: "YOU ARE IN WS CONNECTION"}
	conn := WebSocketConnection{Conn: ws}

	clientsMutex.Lock()
	clients[conn] = ""
	clientsMutex.Unlock()

	if err := ws.WriteJSON(response); err != nil {
		log.Println("WriteJSON error:", err)
	}

	go ListenForWs(&conn)
}

func ListenForWs(conn *WebSocketConnection) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Error:", fmt.Sprintf("%v", r))
		}
		conn.Close()
		clientsMutex.Lock()
		delete(clients, *conn)
		clientsMutex.Unlock()
	}()

	conn.SetReadLimit(maxMessageSize)
	conn.SetReadDeadline(time.Now().Add(pongWait))
	conn.SetPongHandler(func(string) error {
		conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	var payload WsPayload
	for {
		if err := conn.ReadJSON(&payload); err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		payload.Conn = *conn
		wsChan <- payload
	}
}

func ListenToWsChannel() {
	for {
		e := <-wsChan
		response := WsJsonResponse{}

		if canSendMessage(e.Conn) {
			switch e.Action {
			case "username":
				clientsMutex.Lock()
				clients[e.Conn] = e.Username
				clientsMutex.Unlock()
				users := GetUserList()

				response.Action = "list_users"
				response.ConnectedUsers = users
				broadcastToAll(&response)

			case "left":
				clientsMutex.Lock()
				delete(clients, e.Conn)
				clientsMutex.Unlock()
				users := GetUserList()
				response.Action = "list_users"
				response.ConnectedUsers = users

				broadcastToAll(&response)

			case "broadcast":
				response.Action = "broadcast"
				response.Message = fmt.Sprintf("%s: %s", e.Username, e.Message)
				broadcastToAll(&response)
			}
		} else {
			log.Println("Rate limit exceeded for user:", e.Username)
		}
	}
}

func canSendMessage(conn WebSocketConnection) bool {
	rateLimiterMutex.Lock()
	defer rateLimiterMutex.Unlock()

	now := time.Now()
	lastTime, exists := lastMessageTime[conn]

	if !exists || now.Sub(lastTime) > rateLimit {
		lastMessageTime[conn] = now
		return true
	}
	return false
}

func sendMessage(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Println("ParseForm error:", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	message := r.FormValue("message")
	username := r.FormValue("username")
	if message == "" || username == "" {
		http.Error(w, "Message and username cannot be empty", http.StatusBadRequest)
		return
	}

	wsPayload := WsPayload{
		Action:   "broadcast",
		Username: username,
		Message:  message,
	}

	wsChan <- wsPayload
}

func GetUserList() []string {
	clientsMutex.Lock()
	defer clientsMutex.Unlock()

	var userList []string
	for _, username := range clients {
		if username != "" {
			userList = append(userList, username)
		}
	}
	sort.Strings(userList)
	return userList
}

func broadcastToAll(response *WsJsonResponse) {
	clientsMutex.Lock()
	defer clientsMutex.Unlock()

	for client := range clients {
		go func(c WebSocketConnection) {
			if err := c.WriteJSON(response); err != nil {
				log.Println("WebSocket error:", err)
				_ = c.Close()
				clientsMutex.Lock()
				delete(clients, c)
				clientsMutex.Unlock()
			}
		}(client)
	}
}
