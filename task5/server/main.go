package main

import (
	"server/controllers"
	"server/database"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func main() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{http.MethodGet, http.MethodPost},
	}))

	database.Connect()

	e.GET("/products", controllers.GetProducts)
	e.POST("/payment", controllers.HandlePayment)
	e.POST("/cart/save", controllers.SaveCart)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
