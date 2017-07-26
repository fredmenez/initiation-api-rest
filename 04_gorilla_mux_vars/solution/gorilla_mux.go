package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	Routes = []Route{
		Route{"ViewCocktail", "/cocktail/{cocktail_id}", ViewCocktailHandler},
		Route{"ListCocktails", "/cocktails", ListCocktailsHandler},
	}
)

func ListCocktailsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<html><h2>Nos cocktails</h2></html>")
}

func ViewCocktailHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["cocktail_id"]

	fmt.Fprintf(w, "<html><h1>Cocktail %s</h1></html>", id)
}

func main() {
	router := mux.NewRouter()

	for _, r := range Routes {
		router.HandleFunc(r.Pattern, r.Handler)
	}

	log.Println("Serveur en écoute sur le port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
