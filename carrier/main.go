package main

import (
	"log"
	"net/http"
	"os"
)

func carrierHandler(w http.ResponseWriter, r *http.Request) {
	servant, err := os.Hostname()

	first := r.URL.Query().Get("first")
	second := r.URL.Query().Get("second")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte("User  " + servant + " is accessing MS Carrier ! first = " + first + "second = " + second))
}

const keyServerAddr = "server 127.0.0.1"

func String(any any) {
	panic("unimplemented")
}

func main() {
	http.HandleFunc("/test", carrierHandler)
	log.Fatal(http.ListenAndServe(":8181", nil))
}