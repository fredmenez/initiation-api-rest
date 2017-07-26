package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Cocktail struct {
	Name        string
	Ingredients []string
}

var (
	Daiquiri = Cocktail{"Daiquiri",
		[]string{"Grenadine", "Jus d'orange", "Et des trucs..."},
	}

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

	if bytes, err := json.Marshal(Daiquiri); err != nil {
		log.Println("Erreur d'encodage JSON:", err)

	} else {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "%s", string(bytes))

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
