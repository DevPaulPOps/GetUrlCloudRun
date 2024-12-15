package gcp

import (
	run "cloud.google.com/go/run/apiv2"
	"cloud.google.com/go/run/apiv2/runpb"
	"context"
	"fmt"
	"google.golang.org/api/iterator"
	"microservice.com/appfyt/config"
	"microservice.com/appfyt/model"
	"microservice.com/appfyt/util"
)

/*
* @param parent : string
* @param client : *run.ServicesClient
* @param ctx : context.Context
* @return *run.ServiceIterator
* Lister chaque service cloud Run
 */
func ListServicesRequest(client *run.ServicesClient, ctx context.Context) *run.ServiceIterator {
	cfg := config.LoadConfig()

	allService := &runpb.ListServicesRequest{
		Parent: cfg.Parent,
	}
	service := client.ListServices(ctx, allService)
	return service
}

/*
* @param client : *run.ServicesClient
* @return []string, []string
* Récupérer les Urls et Noms de chaque service
 */
func GetAllUrl(client *run.ServiceIterator) []model.Service {
	var service []model.Service
	for {
		allService, err := client.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			util.GestionErrorGetAllUrlAndName(err)
		}
		service = append(service, model.Service{Name: allService.Name, Url: allService.Uri})
	}
	return service

}

/*
* @param service : ServiceIterator
* @return void
* Afficher les informations de chaque service
 */
func PrintAllUri(service []model.Service) {
	for _, allService := range service {
		println("-------------------Service Name-------------------")
		fmt.Println(allService.Name)

		println("-------------------Service URL-------------------")
		fmt.Println(allService.Url)
	}

}

/*
* @param service : ServiceIterator
* @param name : string
* @param parent : string
* @return string
* Renvoyer une URL d'un service par rapport à son nom
 */
func GetURLByName(service *run.ServiceIterator, name string) string {
	cfg := config.LoadConfig()
	nameToVerify := nameToAbsolutName(name, cfg.Parent)

	for {
		allService, err := service.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Println(err)
		}
		if allService.Name == nameToVerify {
			return allService.Uri
		}
	}
	return ""
}

/*
* @param name : string
* @param parent : string
* @return string
* Renvoyer le nom complet d'un service
 */
func nameToAbsolutName(name string, parent string) string {
	return parent + "/services/" + name
}
