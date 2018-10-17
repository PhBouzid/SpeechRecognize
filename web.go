package main

import (
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Audiohandler(w http.ResponseWriter, r *http.Request) {
	//flusher, ok := w.(http.Flusher)
	/*if !ok {
		logwarn("expected http.ResponseWriter to be an http.Flusher")
	}*/
	w.Header().Set("Connection", "Keep-Alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("Transfer-Encoding", "chunked")
	w.Header().Set("Content-Type", "audio/mpeg")
	w.Header().Set("ice-audio-info", "bitrate=128")
	w.Header().Set("icy-br", "128")
	w.Header().Set("icy-description", "Default description")
	w.Header().Set("icy-genre", "Unspecified")
	w.Header().Set("icy-name", "RFM Demo Stream")
	w.Header().Set("icy-pub", "0")
	for true {
		files, err := ioutil.ReadDir("./music")
		if err != nil {
			logerror(err)
		}

		for _, file := range files {
			fmt.Println(file.Name())
			b, err := ioutil.ReadFile("./music/" + file.Name())
			if err != nil {
				logerror(err)
			}
			binary.Write(w, binary.BigEndian, b)
			//flusher.Flush() // Trigger "chunked" encoding
		}
		return
	}
}

func AudiNo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Connection", "Keep-Alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("Transfer-Encoding", "chunked")
	w.Header().Set("Content-Type", "audio/mpeg")
	w.Header().Set("ice-audio-info", "bitrate=128")
	w.Header().Set("icy-br", "128")
	w.Header().Set("icy-description", "Default description")
	w.Header().Set("icy-genre", "Unspecified")
	w.Header().Set("icy-name", "RFM Demo Stream")
	w.Header().Set("icy-pub", "0")
	for true {
		files, err := ioutil.ReadDir("./music")
		if err != nil {
			logerror(err)
		}

		for _, file := range files {
			fmt.Println(file.Name())
			b, err := ioutil.ReadFile("./music/" + file.Name())
			if err != nil {
				logerror(err)
			}
			binary.Write(w, binary.BigEndian, b)
			//flusher.Flush() // Trigger "chunked" encoding
		}
		return
	}
}
