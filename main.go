package main

import (
	_ "cloud.google.com/go"
	"microservice.com/appfyt/api"
)

/**
**********Objectif**********
 * 1. Se connecter à son compte GCP
 * 2. Lister chaque service cloud Run
 * 3. Récupérer les informations de chaque service
 * 4. Afficher les informations de chaque service
 * 5. Renvoyer une URL pour chaque service
*/

func main() {

	api.LauncServ()
}
