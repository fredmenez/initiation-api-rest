package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var (
	Img []byte
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("Requête d'accès reçue à:", time.Now())

	w.Header().Add("Content Type", "image/png")
	w.Write(Img)
}

func main() {
	Img, _ = ioutil.ReadFile("./logo.png")

	http.HandleFunc("/", handler)

	fmt.Println("Serveur en écoute sur le port 8080")
	http.ListenAndServe(":8080", nil)
}
