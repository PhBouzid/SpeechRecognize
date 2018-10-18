package main

import (
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/tcolgate/mp3"
	"fmt"
	"net/http"
	"encoding/binary"
)

// simple hub that manages all the socket connections as a clients
type Hub struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
	track      string
	trackIndex int
	tracks []string
	w http.ResponseWriter
}

// creates a hub
func newHub() Hub {
	return Hub{
		clients:    make(map[*Client]bool),
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		trackIndex: 0,
		tracks:     readTracksAt("./music"),
	}
}

func readTracksAt(dir string) []string {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	songs := make([]string, 0)
	for _, f := range files {
		if f.Name() != ".gitkeep" {
			songs = append(songs, dir+"/"+f.Name())
			fmt.Println(dir+"/"+f.Name())
		}
	}
	return songs
}

// main hub of this app. To send something to a client use this hub
var MusicHub = newHub()

// function registers new client to a hub
func (hub *Hub) Register(client *Client) {
	hub.register <- client
}

// function removes a client from a hub
func (hub *Hub) unregisterClient(c *Client) {
	if _, ok := hub.clients[c]; ok {
		delete(hub.clients, c)
		close(c.send)
	}
}

// function checks is hub empty
func (hub *Hub) isEmpty() bool {
	return len(hub.clients) == 0
}

// function returns number of clients in a hub
func (hub *Hub) count() int {
	return len(hub.clients)
}

// function makes preparations to broadcast message to all clients
// if there are clients in a hub
func (hub *Hub) SendMessage(message []byte) {
	if hub.isEmpty() {
		return
	}

	hub.broadcast <- message
}

func (hub *Hub) RunHttp(){
	go hub.stream()
}

// main hub process
func (hub *Hub) Run() {
	go hub.stream()

	for {
		select {
		case c := <-hub.register:
			hub.clients[c] = true
		case c := <-hub.unregister:
			hub.unregisterClient(c)
		case m := <-hub.broadcast:
			hub.broadcastMessage(m)
		}
	}
}

// function broadcasts prepared message to all the clients in the hub
func (hub *Hub) broadcastMessage(m []byte) {
	for c := range hub.clients {
		select {
		case c.send <- m:
		default:
			close(c.send)
			delete(hub.clients, c)
		}
	}
}

func (hub *Hub) broadcastMessageViaHTTP(m []byte){
	flusher, ok := hub.w.(http.Flusher)
	if !ok {
		logwarn("expected http.ResponseWriter to be an http.Flusher")
	}
	hub.w.Header().Set("Connection", "Keep-Alive")
	hub.w.Header().Set("Access-Control-Allow-Origin", "*")
	hub.w.Header().Set("X-Content-Type-Options", "nosniff")
	hub.w.Header().Set("Transfer-Encoding", "chunked")
	hub.w.Header().Set("Content-Type", "audio/mpeg")
	hub.w.Header().Set("ice-audio-info","bitrate=128")
	hub.w.Header().Set("icy-br","128")
	hub.w.Header().Set("icy-description","Default description")
	hub.w.Header().Set("icy-genre","Unspecified")
	hub.w.Header().Set("icy-name","RFM Demo Stream")
	hub.w.Header().Set("icy-pub","0")
	binary.Write(hub.w, binary.BigEndian, m)
	flusher.Flush()
}

func (hub *Hub) stream() {
	for {
		if err := hub.selectTrack(); err == nil {
			hub.streamStep()
			fmt.Println(hub.track)
		}

		return
	}
}

func (hub *Hub) selectTrack() error {
	if hub.trackIndex >= len(hub.tracks) {
		hub.trackIndex=0
	}
	hub.track = hub.tracks[hub.trackIndex]
	hub.trackIndex++
	return nil
}

func (hub *Hub) streamStep() {
	r, err := os.Open(hub.track)
	defer r.Close()

	if err != nil {
		log.Println(err)
		fmt.Println("here error")
		return
	}

	d := mp3.NewDecoder(r)
	var f mp3.Frame
	skipped := 0
	for {
		if err := d.Decode(&f,&skipped); err != nil {
			log.Println(err)
			fmt.Println("here is problem")
			return
		}
		b := make([]byte, f.Size())

		f.Reader().Read(b)
		hub.SendMessage(b)

		time.Sleep(f.Duration())
	}
}

func (hub *Hub) TrackInfo() (mp3.Frame, error) {
	var f mp3.Frame

	r, err := os.Open(hub.track)
	if err != nil {
		log.Println(err)
		fmt.Println("on track info error")
		return f, err
	}
	skipped := 0
	d := mp3.NewDecoder(r)
	d.Decode(&f,&skipped)
	fmt.Println(&f)
	return f, nil
}