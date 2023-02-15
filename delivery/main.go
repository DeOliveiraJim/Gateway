package main

import (
	"log"
	"net/http"
	"os"
	"strings"
)

func errandHandler(w http.ResponseWriter, r *http.Request) {
	servant, err := os.Hostname()

	first := r.URL.Query().Get("first")
	second := r.URL.Query().Get("second")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	param := strings.TrimPrefix(r.URL.Path, "/test/")
	w.Write([]byte("User  " + servant + " is accessing MS Delivery ! first = " + first + "second = " + second + "param = {" + param + "}"))
}

const keyServerAddr = "server 127.0.0.1"

func String(any any) {
	panic("unimplemented")
}

func main() {
	http.HandleFunc("/test/", errandHandler)
	log.Fatal(http.ListenAndServe(":8484", nil))
}
