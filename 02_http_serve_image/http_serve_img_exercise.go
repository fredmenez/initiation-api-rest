package main

import (
	"fmt"
)

func main() {
	fmt.Println(`/// MISSION /// Un serveur d'image basique

    => Placer une image PNG à côté de votre programme

    => Chargez ce fichier image avec ioutil.ReadFile

    => Ajoutez un request handler pour la route /img/

    => Régler le type de contenu à image/png en ajoutant un en-tête, cf. http.ResponseWriter.Header()

    => Afin de pouvoir finalement servir les images à un client demandant l'url : http//localhost/img/


Proposition d'approfondissement
    => Charger plusieurs images
    => Les référencer dans une map
    => Décoder leur nom de fichier depuis l'URL, ex: /img/zebre
    => Servir les octets de l'image correspondante
`)

}
