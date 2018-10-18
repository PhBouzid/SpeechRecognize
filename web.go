package main

import (
	"net/http"
	"encoding/binary"
	"fmt"
	"os"
	"log"
	"github.com/tcolgate/mp3"
	"time"
)

const (
	chunksize int =  44100 * 2
)

func AudiohandlerMP3(w http.ResponseWriter, r *http.Request){
	//flusher, ok := w.(http.Flusher)
	/*if !ok {
		logwarn("expected http.ResponseWriter to be an http.Flusher")
	}*/
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
	hub :=MusicHub
	for true{
		hub.selectTrack()
		r, err := os.Open(hub.track)
		defer r.Close()

		if err != nil {
			log.Println(err)
			fmt.Println("here error")
			break
		}

		d := mp3.NewDecoder(r)
		var f mp3.Frame
		skipped := 0
		for {
			if err := d.Decode(&f,&skipped); err != nil {
				log.Println(err)
				fmt.Println("here is problem")
				break
			}
			b := make([]byte, f.Size())
			f.Reader().Read(b)
			binary.Write(w, binary.BigEndian, b)
		//	flusher.Flush()
			time.Sleep(f.Duration())

		}
	}

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
	/*
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
				*//*if(file.Name()=="advert.m4a"){
					playerlocal()
					binary.Write(w, binary.BigEndian, make([]byte,1000))
					flusher.Flush()
				}*/
		//}
		hub :=MusicHub
		//go hub.RunHttp()
		for true{
			fmt.Println("in for loop")
			hub.RunHttp()
			binary.Write(w, binary.BigEndian, hub.broadcast)
			flusher.Flush()
			/*select{
			case m :=<-hub.broadcast:
				fmt.Println("broadcast is not empty")
				binary.Write(w, binary.BigEndian, m)
				flusher.Flush()
			}*/
		}

		return
	//}
}





