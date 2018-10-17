package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func logerror(err error) {
	CreateDirIfNotExist("./logs/errors")
	t := time.Now()
	date := t.Format("2006-01-02")
	f, err := os.OpenFile("./logs/errors/log_"+date+".txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		return
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println(t, err)
}

func logevents(event string) {
	CreateDirIfNotExist("./logs/events")
	fmt.Println(event)
	t := time.Now()
	date := t.Format("2006-01-02")
	f, err := os.OpenFile("./logs/events/events_"+date+".txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		return
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println(t, `	`+event)
}

func logwarn(warn string){
	CreateDirIfNotExist("./logs/events")
	fmt.Println(warn)
	t := time.Now()
	date := t.Format("2006-01-02")
	f, err := os.OpenFile("./logs/events/events_"+date+".txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		return
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println(t, `	`+warn)
}

func CreateDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			fmt.Println("error on create folder " + err.Error())
		} else {
			logevents("Folder created: " + dir)
		}
	}
}
