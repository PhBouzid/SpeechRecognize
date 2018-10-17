package main

import (
	"encoding/json"
	"net/http"
	"encoding/binary"
	"io/ioutil"
	"log"
	"fmt"
)

func audiohandler(w http.ResponseWriter, r *http.Request) {
	flusher, ok := w.(http.Flusher)
	if !ok {
		panic("expected http.ResponseWriter to be an http.Flusher")
	}
	w.Header().Set("Connection", "Keep-Alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("Transfer-Encoding", "chunked")
	w.Header().Set("Content-Type", "audio/wave")
	for true {
		files, err := ioutil.ReadDir("./music")
		if err != nil {
			log.Fatal(err)
		}

		for _, file := range files {
			fmt.Println(file.Name(), file.IsDir())
			ioutil.ReadFile(file.Name())
			buffer := make([]float32, 44100 * 4)
			binary.Write(w, binary.BigEndian,buffer)
			flusher.Flush() // Trigger "chunked" encoding
		}

		return
	}
}

func ResponceErrorCreate(err error, StatusCode int) (RespWR []byte) {
	rWrong := ResponceJson{
		Status:     err.Error(),
		StatusCode: StatusCode,
	}

	RespWR, err = json.Marshal(rWrong)
	if err != nil {
		logerror(err)
	}

	return RespWR
}
