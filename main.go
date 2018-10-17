package main

import (
	"net/http"
)

func main(){
	http.HandleFunc("/audio",audiohandler)
	http.ListenAndServe(":7008",nil)
}

