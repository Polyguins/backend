package main

import (
	"github.com/polyguins/backend/server"
	"log"
	"net/http"
)

func main() {
	server.AllRooms.Init()

	http.HandleFunc("/create", server.CreateRoomRequestHandler)
	http.HandleFunc("/join", server.JoinRoomRequestHandler)

	log.Println("Server started on Port 8000")
  err := http.ListenAndServe(":8000", nil) 
  if err != nil {
		log.Fatal(err)
	}
}
