package gcp

import (
	run "cloud.google.com/go/run/apiv2"
	"context"
	"google.golang.org/api/option"
	"microservice.com/appfyt/config"
	"microservice.com/appfyt/util"
)

/*
* @param ctx : context.Context
* @param fileToContextJson : string
* @return *run.ServicesClient
* Se connecter à son compte GCP via le fichier de clé de compte de service
 */
func ConnectToGCP(ctx context.Context) *run.ServicesClient {
	cfg := config.LoadConfig()
	client, err := run.NewServicesClient(ctx, option.WithCredentialsFile(cfg.AccountJSON))
	if err != nil {
		util.GestionErrorConnectToGCP(err)
	}
	//defer client.Close()
	return client
}
