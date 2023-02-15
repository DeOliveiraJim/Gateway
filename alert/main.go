package main

import (
	"log"
	"net/http"
	"os"
	"strings"
)

const keyServerAddr = "server 127.0.0.1"

func alertGetHandler(w http.ResponseWriter, r *http.Request) {
	servant, err := os.Hostname()

	first := r.URL.Query().Get("first")
	second := r.URL.Query().Get("second")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	param := strings.TrimPrefix(r.URL.Path, "/get/")
	w.Write([]byte("User  " + servant + " is accessing MS Alert ! first = " + first + "second = " + second + "param = {" + param + "}"))
}

func String(any any) {
	panic("unimplemented")
}

func main() {
	http.HandleFunc("/get/", alertGetHandler)

	log.Fatal(http.ListenAndServe(":8383", nil))
}
