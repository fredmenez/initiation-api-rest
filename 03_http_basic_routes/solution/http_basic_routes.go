package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
)

type Route struct {
	Name    string
	Pattern string
	Handler http.HandlerFunc
}

var (
	RoutesCocktails = []Route{
		Route{"Accès par nom", "/cocktail/", handlerCocktail},
		Route{"Recherche par ingrédient", "/recherche/ingredient/", handlerSearch},
	}
)

func handlerCocktail(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<html><h2>Recette du \"Panashake\"</h2></html>")
}

func handlerSearch(w http.ResponseWriter, r *http.Request) {
	param := filepath.Base(r.URL.Path)

	fmt.Println("cocktail:", param)

	fmt.Fprintf(w, "<html><h2>Cocktails à base de %s</h2></html>", param)
}

func main() {
	for _, r := range RoutesCocktails {
		http.HandleFunc(r.Pattern, r.Handler)
	}

	fmt.Println("Serveur à l'écoute sur le port 8080")
	log.Fatalln(http.ListenAndServe(":8080", nil))
}
