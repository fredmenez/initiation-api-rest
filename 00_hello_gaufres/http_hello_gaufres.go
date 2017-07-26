package main

import (
	"fmt"
	"log"
	"net/http"
)

/* VOTRE MISSION : faites en sort que ce programme affiche le temps écoulé depuis le démarrage du serveur

   => importez la bibliothèque time
   => gardez une copie du temps au démarrage du programme avec time.Now()
   => calculez une durée avec time.Time.Sub
   => et imprimer cette durée dans la fonction de génération de réponse handler, cf. Fprintf

*/

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<html><h2>Hello gopher</h2></html>")
}

func main() {
	http.HandleFunc("/", handler)

	fmt.Println("Serveur à l'écoute sur le port 8080")
	log.Fatalln(http.ListenAndServe(":8080", nil))
	// A NOTER
	//
	// l'omission d'une adresse ip avant le port veut dire
	// qu'on écoute sur toutes les interfaces réseau
}
