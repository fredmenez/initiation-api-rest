package main

import (
	"fmt"
	"net/http"
)

/* VOTRE MISSION :

   => Ajouter une route /recherche/ingredient/ avec son propre handler (fonction de réponse HTTP)

   => Renseigner les routes non plus manuellement mais avec une boucle for

   => Trouver un moyen d'afficher le nom de l'ingrédient dans la réponse produite
   	  ex.: => quand on accède à /recherche/ingredient/coco, la page affiche : "Cocktails contenant de la coco"

	  Piste: cf. fonction filepath.Base
*/

type Route struct {
	Name    string
	Pattern string
	Handler http.HandlerFunc
}

var (
	RoutesCoktails = []Route{
		Route{"Accès par nom",
			"/cocktail/",
			handlerCocktail},
	}
)

func handlerCocktail(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<html><h2>Recette du \"Panashake\"</h2></html>")
}

func main() {
	http.HandleFunc(RoutesCoktails[0].Pattern, RoutesCoktails[0].Handler)

	fmt.Println("Serveur à l'écoute sur le port 8080")
	http.ListenAndServe(":8080", nil)
}
