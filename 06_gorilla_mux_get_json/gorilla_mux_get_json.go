package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

/* /// VOTRE MISSION /// Associer des méthodes aux routes

   => Définir un type struct Cocktail
   => Créer une variable globale Cocktail prédéfinie, ex: Daiquiri

   => Dans la fonction ViewCocktailHandler

      => Régler le type de contenu à content/json avec ResponseWriter.Header().Add("Content Type", ...)
   	  => encoder en JSON la variable globale
      => Ecrire le résultat de l'encodage dans le ResponseWriter

*/

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
	}
)

func ListCocktailsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<html><h2>Nos cocktails</h2></html>")
}

func CocktailHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["cocktail_id"]

	fmt.Fprintf(w, "<html><h1>Cocktail %s</h1>", id)
	fmt.Fprintf(w, "Méthode: %s", r.Method)
	fmt.Fprintf(w, "</html")
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
