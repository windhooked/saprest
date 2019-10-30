package main

// Hannes de Waal  2019

import (
	"log"
	"net/http"

	"encoding/json"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)

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

	server.GET("/health", health)
	server.POST("/rfc/:function", sapRfc)

	log.Fatal(server.Start(":8088"))
}

func health(c echo.Context) error {

	//"STFC_STRUCTURE"
	params := make(map[string]interface{})
	r, err := callSAP("RFC_PING", params)

	if err != nil {
		return echo.NewHTTPError(401, err)
	}

	return c.JSON(http.StatusOK, r)
}

func sapRfc(c echo.Context) error {

	function := c.Param("function")

	params := make(map[string]interface{})
	log.Print(function)

	err := json.NewDecoder(c.Request().Body).Decode(&params)

	if err != nil {
		return echo.NewHTTPError(401, err)
	}
	r, err := callSAP(function, params)

	if err != nil {
		return echo.NewHTTPError(401, err)
	}

	return c.JSON(http.StatusOK, r)
}
