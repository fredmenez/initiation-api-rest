package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

/* /// VOTRE MISSION /// Associer des méthodes aux routes

   => Ajouter un membre Method à la struct Route
   => Initialiser vos routes avec une valeur de méthode (GET, PUT, ...)
   => Localiser l'endroit où l'on créé les routes Gorilla et ajouter un appel à mux.Route.Methods
   => Afficher la méthode utilisée dans la réponse renvoyée au client

*/

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
		gorillaRoute := router.Path(r.Pattern)
		gorillaRoute.Name(r.Name)
		gorillaRoute.Handler(r.Handler)

		// écriture équivalente chaînée à la JS:
		// router.Path(r.Pattern).
		// 	Name(r.Name).
		// 	Handler(r.Handler)
	}

	log.Println("Serveur en écoute sur le port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
