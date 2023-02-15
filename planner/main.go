package main

import (
	"log"
	"net/http"
	"os"
	"strings"
)

func plannerGetHandler(w http.ResponseWriter, r *http.Request) {
	servant, err := os.Hostname()

	first := r.URL.Query().Get("first")
	second := r.URL.Query().Get("second")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	param := strings.TrimPrefix(r.URL.Path, "/get/")
	w.Write([]byte("User  " + servant + " is accessing MS planner ! first = " + first + "second = " + second + "param = {" + param + "}"))
}

func plannerPostHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		// in case of any error
		return
	}

	// Use the r.PostForm.Get() method to retrieve the relevant data fields
	// from the r.PostForm map.
	value := r.PostForm.Get("name")
	w.Write([]byte("POST REQUEST , name = " + value))
}

func plannerPatchHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		// in case of any error
		return
	}

	// Use the r.PostForm.Get() method to retrieve the relevant data fields
	// from the r.PostForm map.
	value := r.PostForm.Get("name")
	w.Write([]byte("PATCH REQUEST , name = " + value))
}

const keyServerAddr = "server 127.0.0.1"

func String(any any) {
	panic("unimplemented")
}

func main() {
	http.HandleFunc("/get/", plannerGetHandler)
	http.HandleFunc("/post/", plannerPostHandler)
	http.HandleFunc("/patch/", plannerPatchHandler)
	log.Fatal(http.ListenAndServe(":8585", nil))
}
