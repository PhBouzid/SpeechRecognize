package main

import (
	"net/http"
)

func main(){
	http.HandleFunc("/audio",Audiohandler)
	http.ListenAndServe(":7008",nil)
}

