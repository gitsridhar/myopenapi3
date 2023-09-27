// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.14.0 DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	. "github.com/gitsridhar/myopenapi3/echo/Impl/api/models"
	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Returns all fruits
	// (GET /fruits)
	FindFruits(ctx echo.Context, params FindFruitsParams) error
	// Creates a new fruit
	// (POST /fruits)
	AddFruit(ctx echo.Context) error
	// Delete a fruit by ID
	// (DELETE /pets/{id})
	DeleteFruit(ctx echo.Context, id int64) error
	// Returns a fruit by ID
	// (GET /pets/{id})
	FindFruitByID(ctx echo.Context, id int64) error
	// Update a fruit by ID
	// (PUT /pets/{id})
	UpdateFruitByID(ctx echo.Context, id int64) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// FindFruits converts echo context to params.
func (w *ServerInterfaceWrapper) FindFruits(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params FindFruitsParams
	// ------------- Optional query parameter "tags" -------------

	err = runtime.BindQueryParameter("form", true, false, "tags", ctx.QueryParams(), &params.Tags)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter tags: %s", err))
	}

	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.FindFruits(ctx, params)
	return err
}

// AddFruit converts echo context to params.
func (w *ServerInterfaceWrapper) AddFruit(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.AddFruit(ctx)
	return err
}

// DeleteFruit converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteFruit(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int64

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.DeleteFruit(ctx, id)
	return err
}

// FindFruitByID converts echo context to params.
func (w *ServerInterfaceWrapper) FindFruitByID(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int64

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.FindFruitByID(ctx, id)
	return err
}

// UpdateFruitByID converts echo context to params.
func (w *ServerInterfaceWrapper) UpdateFruitByID(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int64

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.UpdateFruitByID(ctx, id)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/fruits", wrapper.FindFruits)
	router.POST(baseURL+"/fruits", wrapper.AddFruit)
	router.DELETE(baseURL+"/pets/:id", wrapper.DeleteFruit)
	router.GET(baseURL+"/pets/:id", wrapper.FindFruitByID)
	router.PUT(baseURL+"/pets/:id", wrapper.UpdateFruitByID)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+RWTXPbNhD9K5htj6yoOpkeeKoT2TOaaZ1O3fSS6gATSxEt8WFgaYWj4X/vACAlW6Sc",
	"pO100uYkCtwF3r73dsE9lEZZo1GTh2IPvqxR8fh45Zxx4cE6Y9GRxLhcGoHhV6AvnbQkjYYiBbP4LoPK",
	"OMUJCpCaXlxABtRZTH9xiw76DBR6z7dnNxpfH1I9Oam30PcZOLxvpUMBxTsYDhzDN30G166VFPblTfOm",
	"guLdHr52WEEBX+XHWvOh0PwGdymjz04rlWIK762W9y0yKZipGNXIqpj8tObvXs7UfIJcCtj0Ae8BwIRp",
	"zdUMQTdc4enhJyRlQHw7zfylsx/KPAEZEWz6sCx1ZZL8mngZ0aLisgnJToqyMa34fhtWFqVRkA3o4dZJ",
	"UXPHfkX9Bw8nti7k1ETWF3l+yI1ZfXYC+ZJd/rRmVHNiApXRnhwn9Ak+82QcskAZD/EeMmhkidrjkT64",
	"tLyskV0slpPDd7vdgsfXC+O2+ZDr8x/Wr69ubq++uVgsFzWpJjKKTvk31S26B1niuQryGJYHYiU1IezH",
	"jkV52W0ACxk8oPOpuG8Xy8Uy7G0sam4lFPAiLmVgOdXRA3msND5ukaaa/ozUOu0Zb5pESuDGqKhxpOc3",
	"DfGARNFaQAHXUovrtG04yXGFhM7HVnm6O/GtZ2RYJRtCx+46CEaAAu5bdN1R5BAH2TA8YusQqgh66sy0",
	"wJ3jXfjvqYs8hf6JPfgUgeLvpWoV0626Qxfc69C3DUVYLhZ/BlMjVXT4EdQHx1K/Cfb31gQXhIyL5XL0",
	"POo0VKxtZBm5zH/3AeJ+ruznJs44bk6o6CfmTyYfAaXmqHjb0Cdheg5KGvEzR7ca31ssCQXDISYD3yrF",
	"XTfruoDOGj/j0NcOY8typnE3NK7UR4Mu2KpN+EOQw7Cp2aGY2PZSJNdCmlHo6ZUR3T/GxfEimNKROpgM",
	"40KEnwN4eDwvybXY/00HfYRx/ktGmRE/RuQWyed7KfpkmAZp5qpbxXXGmZd62wyXFrvr2HrFrDMPUszY",
	"JCWNTnl2vK1XYaBUo7gDjGGchBF8nCZSTKQ+N1rmb//paHk5LThBSTjE56TjQYlHEgR8z99JYzT3KJjR",
	"RyHXq/OX0qsuvv0U4Sqksv7XdPuiG3qi7egE28444a0V/LFtPsYIKeevWqGN2f8/Lwx1fUZOONU2GSHG",
	"oHsYxXr6wd2Ykje18ZSHz91+0/8ZAAD//6yT0DZ/DgAA",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}