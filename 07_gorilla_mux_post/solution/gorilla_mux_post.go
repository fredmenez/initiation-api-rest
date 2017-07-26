package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Cocktail struct {
	Name        string
	Ingredients []string
}

var (
	Routes = []Route{
		Route{
			Name:    "ListCocktails",
			Method:  "GET",
			Pattern: "/cocktails",
			Handler: ListCocktailsHandler},

		Route{
			Name:    "ViewCocktail",
			Method:  "GET",
			Pattern: "/cocktail/{cocktail_id}",
			Handler: CocktailHandler},

		Route{
			Name:    "CreateCocktail",
			Method:  "POST",
			Pattern: "/cocktail",
			Handler: CocktailHandler},
	}
)

func ListCocktailsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<html><h2>Nos cocktails</h2>")

	fmt.Fprintf(w, "Méthode: %s", r.Method)

	fmt.Fprintf(w, "</html>")
}

func CocktailHandler(w http.ResponseWriter, r *http.Request) {
	var err error

	if r.Method == "GET" {
		vars := mux.Vars(r)
		id := vars["cocktail_id"]

		fmt.Fprintf(w, "<html><h1>Cocktail %s</h1>", id)

		fmt.Fprintf(w, "Méthode: %s", r.Method)

		fmt.Fprintf(w, "</html>")

	} else if r.Method == "POST" {
		var body []byte

		if body, err = ioutil.ReadAll(r.Body); err != nil {
			log.Println("Erreur de lecture du corps d'une requête POST:", err)
			return
		}

		var ckt Cocktail

		if err := json.Unmarshal(body, &ckt); err != nil {
			log.Println("Erreur décodage JSON:", err)
			return
		}

		fmt.Println("Requête POST de création de cocktail:")
		fmt.Println(ckt)

	}
}

func main() {
	router := mux.NewRouter()

	for _, r := range Routes {
		router.Methods(r.Method).
			Name(r.Name).
			Path(r.Pattern).
			Handler(r.Handler)
	}

	log.Println("Serveur en écoute sur le port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
