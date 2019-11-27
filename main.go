package main

// Hannes de Waal  2019

import (
	"fmt"
	"hannesdw/sap/saprest/sap"
	"log"
	"net/http"

	"encoding/json"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var rfc *sap.Rfc

func main() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	sap, err := sap.New()
	if err != nil {
		log.Fatal(err)
	}

	server := echo.New()
	server.Use(
		middleware.Recover(),   // Recover from all panics to always have your server up
		middleware.Logger(),    // Log everything to stdout
		middleware.RequestID(), // Generate a request id on the HTTP response headers for identification
	)
	server.Debug = true
	//server.Debug = false
	server.HideBanner = true

	server.HTTPErrorHandler = func(err error, c echo.Context) {
		// pass on error if required
		log.Println(c.Path(), c.QueryParams(), err.Error())

		server.DefaultHTTPErrorHandler(err, c)
		//return echo.NewHTTPError(http.StatusUnauthorized, fmt.Sprintf("%s-%s-%s", c.Path(), c.QueryParams(), err.Error()))

	}

	server.GET("/health", healthHandler)
	server.POST("/rfc/:function", sapRfcHandler)

	log.Fatal(server.Start(":8088"))
}

//healthHandler
func healthHandler(c echo.Context) error {

	//"STFC_STRUCTURE"

	err := rfc.Conn.Ping()

	if err != nil {
		return echo.NewHTTPError(401, err)
	}

	return c.JSON(http.StatusOK, "")
}

//sapRfcHandler
func sapRfcHandler(c echo.Context) error {

	function := c.Param("function")

	params := make(map[string]interface{})
	log.Print(function)

	err := json.NewDecoder(c.Request().Body).Decode(&params)

	if err != nil {
		return echo.NewHTTPError(401, err)
	}
	r, err := rfc.Call(function, params)

	if err != nil {
		return echo.NewHTTPError(401, err)
	}

	return c.JSON(http.StatusOK, r)
}
