package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

/* /// VOTRE MISSION ///

   => Ajouter une route ViewCocktail : /cocktail pour accéder à un cocktail
   => Ajouter ces routes au routeur Gorilla avec une boucle for

   => Ajouter un paramètre d'URL à la route ViewCocktail avec mux.Vars :

      * le pattern devient : /cocktail/{cocktail_id}

      * Dans votre fonction réponse (handler) :
        vars := mux.Vars(r)
	    id := vars["cocktail_id"]

*/

type Route struct {
	Name    string
	Pattern string
	Handler http.HandlerFunc
}

func ListCocktailsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<html><h2>Nos cocktails</h2></html>")
}

func main() {
	router := mux.NewRouter()

	r := Route{"ListCocktails", "/cocktails", ListCocktailsHandler}

	router.HandleFunc(r.Pattern, r.Handler)

	log.Println("Serveur en écoute sur le port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
