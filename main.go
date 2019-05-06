package main

import (
	"net/http"
	"time"
)

func main(){
	http.HandleFunc("/now",func(w http.ResponseWriter, r *http.Request){
		var now = time.Now().Format(" 2006-01-02 15:04:05")
		w.Write([]byte(now))
	})
	http.ListenAndServe(":6000",nil)
}