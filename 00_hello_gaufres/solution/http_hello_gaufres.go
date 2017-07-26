package main

import (
	"fmt"
	"net/http"
	"time"
)

var (
	startTime time.Time
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<html><h2>Hello gopher</h2> Ce programme tourne depuis : %v</html>", time.Now().Sub(startTime))
}

func main() {
	startTime = time.Now()

	http.HandleFunc("/", handler)

	fmt.Println("Serveur à l'écoute sur le port 8080")
	http.ListenAndServe(":8080", nil)
	// A NOTER
	//
	// l'omission d'une adresse ip avant le port veut dire
	// qu'on écoute sur toutes les interfaces réseau
}
