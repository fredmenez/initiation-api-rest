package main

/* /// VOTRE MISSION ///

   => Réparer le programme pour que le paramètre de port soit bien récupéré

   => Ajouter un paramètre de config pour configurer l'adresse de l'interface réseau à utiliser

*/

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world !")
}

type Config struct {
	port string `json:"port"`
}

func main() {
	fileBytes, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Fatalln("Erreur de chargement de fichier:", err)
	}

	var config Config

	if err := json.Unmarshal(fileBytes, &config); err != nil {
		log.Fatalln("Erreur de décodage JSON:", err)
	}

	http.HandleFunc("/", handler)

	fmt.Println("Serving from port :", config.port)
	log.Fatalln(http.ListenAndServe(":"+config.port, nil))
}
