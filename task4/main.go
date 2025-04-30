package main

import (
    "task4/controllers"
    "task4/database"
    "github.com/labstack/echo/v4"
)

func main() {
    e := echo.New()
    database.Connect()

    e.POST("/products", controllers.CreateProduct)
    e.GET("/products", controllers.GetProducts)
    e.GET("/products/:id", controllers.GetProduct)
    e.PUT("/products/:id", controllers.UpdateProduct)
    e.DELETE("/products/:id", controllers.DeleteProduct)

	e.POST("/categories", controllers.CreateCategory)
    e.GET("/categories", controllers.GetCategories)
    e.GET("/categories/:id", controllers.GetCategory)
    e.PUT("/categories/:id", controllers.UpdateCategory)
    e.DELETE("/categories/:id", controllers.DeleteCategory)

	e.POST("/suppliers", controllers.CreateSupplier)
    e.GET("/suppliers", controllers.GetSuppliers)
    e.GET("/suppliers/:id", controllers.GetSupplier)
    e.PUT("/suppliers/:id", controllers.UpdateSupplier)
    e.DELETE("/suppliers/:id", controllers.DeleteSupplier)

	e.POST("/cart", controllers.AddToCart)
    e.GET("/carts", controllers.GetCarts)
    e.GET("/cart/:user_id", controllers.GetCart)
    e.PUT("/cart/:user_id", controllers.UpdateCart)
    e.DELETE("/cart/:user_id", controllers.DeleteCart)

    e.Logger.Fatal(e.Start(":8080"))
}