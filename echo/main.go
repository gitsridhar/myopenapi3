package main

import (
    "net/http"
    
    "github.com/labstack/echo/v4"
)

func main() {
    abc := func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello, Func Ptr!")
    }

    e := echo.New()
    e.GET("/", abc)
    e.Logger.Fatal(e.Start(":1323"))
}

func sridhar(c echo.Context) error {
    return c.String(http.StatusOK, "Hello, Sridhar!")
}

