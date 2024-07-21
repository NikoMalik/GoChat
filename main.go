package main

import (
	"Chat/internal/handler"
	"Chat/router"
	"log"
	"net/http"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	r := router.Setup()
	server := &http.Server{
		Addr:    ":8000",
		Handler: r,
	}

	log.Println("Channel start listening...")
	go handler.ListenToWsChannel()

	log.Println("Server started on :8000")
	log.Fatal(server.ListenAndServe())
}
