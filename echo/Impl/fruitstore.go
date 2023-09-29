package main

// cd api
// mkdir models
// oapi-codegen --config=models.cfg.yaml ../../../myfruit.yaml
// oapi-codegen --config=server.cfg.yaml ../../../myfruit.yaml

// create: curl -i http://localhost:8080/fruits  -d '{"name":"apple"}' -H 'Content-Type: application/json'
// get:    curl -i http://localhost:8080/fruits -H 'Content-Type: application/json'
// update: curl -i http://localhost:8080/fruits/1000 -H 'Content-Type: application/json' -X PUT -d '{"name":"newapple"}'
// delete: curl -i http://localhost:8080/fruits/1000 -H 'Content-Type: application/json' -X DELETE

import (
	"flag"
	"fmt"
	"net"
	"os"

	middleware "github.com/deepmap/oapi-codegen/pkg/middleware"
	"github.com/gitsridhar/myopenapi3/echo/Impl/api"
	echo "github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
)

func main() {
	port := flag.String("port", "8080", "Port for HTTP server")
	flag.Parse()

	swagger, err := api.GetSwagger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading swagger spec\n: %s", err)
		os.Exit(1)
	}

	swagger.Servers = nil

	fruitStore := api.NewFruitStore()

	e := echo.New()
	e.Use(echomiddleware.Logger())
	e.Use(middleware.OapiRequestValidator(swagger))

	api.RegisterHandlers(e, fruitStore)

	e.Logger.Fatal(e.Start(net.JoinHostPort("0.0.0.0", *port)))
}
