package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

/* /// VOTRE MISSION /// Associer des méthodes aux routes

   => Ajouter une route CreateCocktail avec méthode POST et lui faire utiliser le handler CocktailHandler

   => Dans CocktailHandler utiliser un if pour traiter séparément les scénarios des méthodes GET et POST

   => Pour la méthode POST récupérer le contenu du corps de la requête en lisant http.Request.Body
      Suggestion : cf. la fonction util io.ReadAll pour lire tout le body d'un coup

   => Créer un type struct Cocktail comportant 2 champs : Name et Ingredients

   => Décoder le octets du corps de la requête dans une valeur struct Cocktail

   => Afficher le contenu de cette valeur Cocktail avec un simple fmt.Println

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
