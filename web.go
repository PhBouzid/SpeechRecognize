package main

import (
	"net/http"
	"encoding/binary"
	"io/ioutil"
	"fmt"
	"bufio"
	"io"
	"log"
	"os"
)

const (
	chunksize int =  44100 * 2
)

func openFile(name string, buffer []byte) (byteCount int) {
	data, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()

	reader := bufio.NewReader(data)
	for {
		if _, err = reader.Read(buffer); err != nil {
			break
		}
	}
	if err != io.EOF {
		log.Fatal("Error Reading ", name, ": ", err)
	} else {
		err = nil
	}

	return
}

func Audiohandler(w http.ResponseWriter, r *http.Request) {
	flusher, ok := w.(http.Flusher)
	if !ok {
		logwarn("expected http.ResponseWriter to be an http.Flusher")
	}
	w.Header().Set("Connection", "Keep-Alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("Transfer-Encoding", "chunked")
	w.Header().Set("Content-Type", "audio/mpeg")
	w.Header().Set("ice-audio-info","bitrate=128")
	w.Header().Set("icy-br","128")
	w.Header().Set("icy-description","Default description")
	w.Header().Set("icy-genre","Unspecified")
	w.Header().Set("icy-name","RFM Demo Stream")
	w.Header().Set("icy-pub","0")
	buffer := make([]byte, 44100 * 2)
	go func(){
		for true {
			files, err := ioutil.ReadDir("./music")
			if err != nil {
				logerror(err)
			}
			for _, file := range files {
				fmt.Println(file.Name())
				openFile("./music/"+file.Name(),buffer)
				if err!=nil{
					logerror(err)
				}
				//
			}
		}
	}()
	for true {
		binary.Write(w, binary.BigEndian, &buffer)
		flusher.Flush() // Trigger "chunked" encoding
		return
	}
}

func AudiNo(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Connection", "Keep-Alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("Transfer-Encoding", "chunked")
	w.Header().Set("Content-Type", "audio/mpeg")
	w.Header().Set("ice-audio-info","bitrate=128")
	w.Header().Set("icy-br","128")
	w.Header().Set("icy-description","Default description")
	w.Header().Set("icy-genre","Unspecified")
	w.Header().Set("icy-name","RFM Demo Stream")
	w.Header().Set("icy-pub","0")
	for true {
		files, err := ioutil.ReadDir("./music")
		if err != nil {
			logerror(err)
		}

		for _, file := range files {
			fmt.Println("play advert")
			adv,err :=ioutil.ReadFile("./web/audio/advert.m4a")
			if err!=nil{
				logerror(err)
			}
			binary.Write(w, binary.BigEndian,adv)
			fmt.Println(file.Name())
			b,err:=ioutil.ReadFile("./music/"+file.Name())
			if err!=nil{
				logerror(err)
			}
			binary.Write(w, binary.BigEndian,b)
			//flusher.Flush() // Trigger "chunked" encoding
		}
		return
	}
}




