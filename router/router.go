package router

import (
	"Chat/internal/handler"
	"net/http"

	"github.com/bmizerany/pat"
)

func Setup() http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handler.Home))
	mux.Get("/ws", http.HandlerFunc(handler.WsEndpoint))
	mux.Get("/chat", http.HandlerFunc(handler.Chat))
	mux.Post("/send", http.HandlerFunc(handler.Message))
	mux.Get("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	return mux

}
