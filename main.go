package main

import (
	"context"
	"flag"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/rafaelherik/terraform-provider-aznamingtool/internal/provider"
)

var (
	version string = "1.0.0"
)

func main() {

	// myClient := apiclient.NewAPIClient("http://localhost:8081", "603a01da-b170-4a0f-8d55-f809332faacd", "", nil)
	// svc := apiclient.NewResourceNamingService(myClient)

	// request := models.ResourceNameRequest{
	// 	ResourceType:        "vm",
	// 	ResourceEnvironment: "dev",
	// 	ResourceFunction:    "func",
	// 	ResourceInstance:    "88",
	// 	ResourceId:          85,
	// 	ResourceLocation:    "aec",
	// 	ResourceProjAppSvc:  "tnp",
	// }

	// response, err := svc.RequestName(request)

	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(response)
	// }

	var debug bool

	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	opts := providerserver.ServeOpts{
		Address: "registry.terrafrom.io/rafaelherik/aznamingtool",
		Debug:   debug,
	}

	err := providerserver.Serve(context.Background(), provider.NewProvider(version), opts)

	if err != nil {
		log.Fatal(err.Error())
	}
}
