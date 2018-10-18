package main

import (
	"github.com/gorilla/websocket"
	"net/http"
	"encoding/json"
)

// basic gorilla/websocket upgrader
var upgrader = websocket.Upgrader{
	ReadBufferSize:  2048,
	WriteBufferSize: 2048,
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

func currentTrackHandler(w http.ResponseWriter, r *http.Request) {
	trackInfo, _ := MusicHub.TrackInfo()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(trackInfo.Duration())
	w.WriteHeader(http.StatusOK)
}