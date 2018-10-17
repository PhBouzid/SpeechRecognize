package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var queue = make(chan bool)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/audio", Audiohandler)
	fs := http.FileServer(http.Dir("./web"))
	router.Handle("/web/", fs)
	router.Handle("/", fs)

	router.HandleFunc("/play", player)
	router.HandleFunc("/ws", wsHandler)

	log.Fatal(http.ListenAndServe(":8030", router))
}

func player(w http.ResponseWriter, r *http.Request) {

	queue <- true
}

func playerlocal() {
	queue <- true
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Upgrade(w, r, w.Header(), 1024, 1024)

	if err != nil {
		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
	}
	go echo(conn)
}

// 3
func echo(conn *websocket.Conn) {
	for {
		val := <-queue
		v := strconv.FormatBool(val)
		fmt.Println("sending")
		if err := conn.WriteMessage(websocket.TextMessage, []byte(v)); err != nil {
			log.Printf("Websocket error: %s", err)
		}
	}
}
