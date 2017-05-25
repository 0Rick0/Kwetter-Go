package main

import (
	"log"
	"flag"
	"github.com/emicklei/go-restful"
	"net/http"
	"./types"
	"./EndPoints"
	"./service"
	"github.com/emicklei/go-restful-swagger12"
)

var swaggerPath string
var useSwagger bool

func main()  {
	parseFlags()
	//setup
	s := service.Service{Kwets:[]types.Kwet{}, Users: map[string]types.User{}}

	s.SetupDatabase()

	userServiceContainer := EndPoints.ServiceContainer{Service: &s}

	wsContainer := restful.NewContainer()
	userServiceContainer.DefineUserEndpoints(wsContainer)

	if useSwagger{
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
			SwaggerFilePath:swaggerPath,
			ApiVersion:"1.0",
		}
		swagger.RegisterSwaggerService(config, wsContainer)
	}

	//serve
	server := &http.Server{Addr:":8080", Handler:wsContainer}

	log.Print("Starting Server on port 8080")
	log.Fatal(server.ListenAndServe())
}

func parseFlags()  {
	//parse config
	flag.BoolVar(&useSwagger, "swagger", false, "If specified, swagger is enabled")
	flag.StringVar(&swaggerPath, "swagger-path", "/home/rick/swagger-ui/dist", "Specify the path of the swagger dist")

	flag.Parse()

}
