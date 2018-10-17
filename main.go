package main

import (
	"net/http"
)

func main(){
	http.HandleFunc("/audio",Audiohandler)

	fs := http.FileServer(http.Dir("./web"))
	http.Handle("/web/", fs)
	http.Handle("/", fs)
	http.ListenAndServe(":7008",nil)
}

