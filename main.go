package main

import (
	"log"
	"github.com/emicklei/go-restful"
	"net/http"
	"./types"
	"./EndPoints"
	"./service"
	"github.com/emicklei/go-restful-swagger12"
	"github.com/droundy/goopt"
)

var swaggerPath *string = goopt.String([]string{"--swagger-path","--sp", "-w"}, "/home/rick/swagger-ui/dist",
	"The path to the swagger installation")
var useSwagger *bool = goopt.Flag([]string{"--swagger","-s"}, []string{}, "Enable swagger", "")
var ipport *string = goopt.String([]string{"--http", "--hostport", "-p"}, ":8080", "Set the port to use")

func main()  {
	parseFlags()
	//setup
	s := service.Service{Kwets:[]types.Kwet{}, Users: map[string]types.User{}}

	//initiate the database and create the service container
	s.SetupDatabase()
	serviceContainer := EndPoints.ServiceContainer{Service: &s}

	//start building the wsContainer
	wsContainer := restful.NewContainer()

	// Add container filter to enable CORS
	cors := restful.CrossOriginResourceSharing{
		AllowedHeaders: []string{"Content-Type", "Accept", "Authorization"},
		AllowedMethods: []string{"GET", "POST"},
		AllowedDomains: []string{"localhost:4200", "localhost:8080", "localhost"},
		CookiesAllowed: false,
		Container:      wsContainer}
	serviceContainer.DefineUserEndpoints(wsContainer)
	serviceContainer.DefineKwetEndpoints(wsContainer)
	//Install the CORS filter
	wsContainer.Filter(cors.Filter)

	if *useSwagger{
		//if swagger is enabled, install the swagger service
		config := swagger.Config{
			WebServices:wsContainer.RegisteredWebServices(),
			WebServicesUrl:"http://localhost:8080",
			ApiPath: "/apidocs.json",
			DisableCORS: true,
			Info: swagger.Info{
				Title: "Kwetter-Go",
				Description:"A Go implementation of the kwetter application",
			},
			SwaggerPath: "/apidocs/",
			SwaggerFilePath:*swaggerPath, //use the configured swagger path
			ApiVersion:"1.0",
		}
		swagger.RegisterSwaggerService(config, wsContainer)
	}

	//serve
	server := &http.Server{Addr:*ipport, Handler:wsContainer}

	log.Printf("Starting Server on %s\n", *ipport)
	log.Fatal(server.ListenAndServe())
}

func parseFlags()  {
	goopt.Description = func() string {
		return "A go implementation of the Kwetter application"
	}
	goopt.Version = "0.1"
	goopt.Parse(nil)
}
