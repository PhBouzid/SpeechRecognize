package main

import (
	"github.com/gorilla/websocket"
	"net/http"
)

// basic gorilla/websocket upgrader
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// function handles GET request and
// upgrades it to websocket connection
func hubHandler(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logerror(err)
		return
	}

	client := CreateClient(ws)
	MusicHub.Register(client)
	go client.ReadPump()
	client.WritePump()
}
