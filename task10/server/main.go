package main

import (
	"os"
	"server/controllers"
	"server/database"
	"server/models"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	database.Connect()
	database.DB.AutoMigrate(&models.User{})

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://task10-front-c7eyc9dfhfabg4fb.polandcentral-01.azurewebsites.net"},
		AllowMethods: []string{echo.GET, echo.POST, echo.OPTIONS},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization,},
	}))

	e.POST("/api/auth/register", controllers.Register)
	e.POST("/api/auth/login", controllers.Login)
	e.GET("/api/auth/google", controllers.GoogleLogin)
	e.GET("/api/auth/google/callback", controllers.GoogleCallback)
	e.GET("/api/auth/github", controllers.GitHubLogin)
	e.GET("/api/auth/github/callback", controllers.GitHubCallback)
	e.GET("/products", controllers.GetProducts)
	e.POST("/payment", controllers.HandlePayment)
	e.POST("/cart/save", controllers.SaveCart)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	e.Logger.Fatal(e.Start(":" + port))
}
