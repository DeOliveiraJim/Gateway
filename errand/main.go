package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
)

func errandGetHandler(w http.ResponseWriter, r *http.Request) {
	servant, err := os.Hostname()

	first := r.URL.Query().Get("first")
	second := r.URL.Query().Get("second")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte("User  " + servant + " is accessing MS Errand ! first = " + first + "second = " + second))
}

func errandPatchHandler(w http.ResponseWriter, r *http.Request) {
	servant, err := os.Hostname()

	first := r.URL.Query().Get("first")
	second := r.URL.Query().Get("second")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte("User  " + servant + " is accessing MS Errand ! first = " + first + "second = " + second))
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	// Récupération de l'URL
	url := r.URL.Path

	// Expression régulière pour extraire l'id
	re := regexp.MustCompile(`^/carrier/(\d+)/itinerary$`)
	matches := re.FindStringSubmatch(url)

	// Si l'URL correspond au modèle attendu, on extrait l'id et on l'affiche
	if len(matches) == 2 {
		id := matches[1]
		fmt.Fprintf(w, "L'id est : %s", id)
	} else {
		http.NotFound(w, r)
	}
}

const keyServerAddr = "server 127.0.0.1"

func String(any any) {
	panic("unimplemented")
}

func main() {
	http.HandleFunc("/get", errandGetHandler)
	http.HandleFunc("/patch", errandPatchHandler)
	http.HandleFunc("/", handleRequest)
	log.Fatal(http.ListenAndServe(":8282", nil))
}
