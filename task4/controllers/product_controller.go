package controllers

import (
	"net/http"
    "task4/database"
    "task4/models"
    "github.com/labstack/echo/v4"
)

func CreateProduct(c echo.Context) error {
	var product models.Product
	if err := c.Bind(&product); err != nil {
        return c.String(http.StatusBadRequest, "Invalid input")
    }
	database.DB.Create(&product)
	return c.JSON(http.StatusCreated, product)
}

func GetProduct(c echo.Context) error {
    id := c.Param("id")
    var product models.Product
    if err := database.DB.Preload("Category").First(&product, id).Error; err != nil {
        return c.String(http.StatusNotFound, "Product not found")
    }
    return c.JSON(http.StatusOK, product)
}

func GetProducts(c echo.Context) error {
    var products []models.Product
    database.DB.Preload("Category").Find(&products)
    return c.JSON(http.StatusOK, products)
}

func DeleteProduct(c echo.Context) error {
    id := c.Param("id")
    var product models.Product
    if err := database.DB.First(&product, id).Error; err != nil {
        return c.String(http.StatusNotFound, "Product not found")
    }
    database.DB.Delete(&product)
    return c.String(http.StatusOK, "Product deleted")
}

func UpdateProduct(c echo.Context) error {
    id := c.Param("id")
    var product models.Product
    if err := database.DB.First(&product, id).Error; err != nil {
        return c.String(http.StatusNotFound, "Product not found")
    }
    if err := c.Bind(&product); err != nil {
        return c.String(http.StatusBadRequest, "Invalid input")
    }
    database.DB.Save(&product)
    return c.JSON(http.StatusOK, product)
}