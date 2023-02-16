package main

import (
	"fmt"
	"io/ioutil"
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

func getBodyAndHeader(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("server: %s /\n", r.Method)
	fmt.Printf("server: query id: %s\n", r.URL.Query().Get("id"))
	fmt.Printf("server: content-type: %s\n", r.Header.Get("content-type"))
	fmt.Printf("server: headers:\n")
	for headerName, headerValue := range r.Header {
		fmt.Printf("\t%s = %s\n", headerName, strings.Join(headerValue, ", "))
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("server: could not read request body: %s\n", err)
	}
	fmt.Printf("server: request body: %s\n", reqBody)

	fmt.Fprintf(w, `{"message": "hello!"}`)
}

const keyServerAddr = "server 127.0.0.1"

func String(any any) {
	panic("unimplemented")
}

func main() {
	http.HandleFunc("/get/", plannerGetHandler)
	http.HandleFunc("/post/", plannerPostHandler)
	http.HandleFunc("/patch/", plannerPatchHandler)
	http.HandleFunc("/", getBodyAndHeader)
	log.Fatal(http.ListenAndServe(":8585", nil))
}
