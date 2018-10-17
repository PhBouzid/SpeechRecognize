package main

import (
	"net/http"
	"encoding/binary"
	"io/ioutil"
	"fmt"
	"bufio"
	"os"
)

const (
	chunksize int =  44100 * 2
)


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

	for true {
		files, err := ioutil.ReadDir("./music")
		if err != nil {
			logerror(err)
		}

		for _, file := range files {
			fmt.Println(file.Name())
			data, err := os.Open("./music/"+file.Name())
			if err != nil {
				logerror(err)
			}
			buffer := make([]byte, 44100 * 2)
			reader := bufio.NewReader(data)
			for  {
				count, err := reader.Read(buffer)
				fmt.Println(count)
				if err != nil {
					logerror(err)
				}
				if count == 0 {
					break
				}
				// if err != io.EOF {
				//  break
				// } else {
				//  err = nil
				// }
				binary.Write(w, binary.BigEndian, buffer)
				flusher.Flush()
			}

			data.Close()
			advData, err := os.Open("./web/audio/advert.m4a")
			advReader := bufio.NewReader(advData)
			advBuffer := make([]byte, 44100 * 2)
			for  {
				count, err := advReader.Read(advBuffer)
				fmt.Println(count)
				if err != nil {
					logerror(err)
				}
				if count == 0 {
					break
				}
				// if err != io.EOF {
				//  break
				// } else {
				//  err = nil
				// }
				binary.Write(w, binary.BigEndian, advBuffer)
				flusher.Flush()
			}
			playerlocal()
			/*if(file.Name()=="advert.m4a"){
				playerlocal()
				binary.Write(w, binary.BigEndian, make([]byte,1000))
				flusher.Flush()
			}*/
		}

		return
	}
}





