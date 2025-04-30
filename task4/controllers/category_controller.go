package controllers

import (
    "net/http"
    "task4/database"
    "task4/models"
    "github.com/labstack/echo/v4"
)

func CreateCategory(c echo.Context) error {
    var category models.Category

    if err := c.Bind(&category); err != nil {
        return c.String(http.StatusBadRequest, "Invalid input")
    }
	
    if category.Name == "" {
        return c.String(http.StatusBadRequest, "Name is required")
    }

    database.DB.Create(&category)
    return c.JSON(http.StatusCreated, category)
}

func GetCategories(c echo.Context) error {
    var categories []models.Category
    database.DB.Find(&categories)
    return c.JSON(http.StatusOK, categories)
}

func GetCategory(c echo.Context) error {
    id := c.Param("id")
    var category models.Category
	
    if err := database.DB.First(&category, id).Error; err != nil {
        return c.String(http.StatusNotFound, "Category not found")
    }

    return c.JSON(http.StatusOK, category)
}

func UpdateCategory(c echo.Context) error {
    id := c.Param("id")
    var category models.Category

    if err := database.DB.First(&category, id).Error; err != nil {
        return c.String(http.StatusNotFound, "Category not found")
    }
    if err := c.Bind(&category); err != nil {
        return c.String(http.StatusBadRequest, "Invalid input")
    }
    if category.Name == "" {
        return c.String(http.StatusBadRequest, "Name is required")
    }

    database.DB.Save(&category)
    return c.JSON(http.StatusOK, category)
}

func DeleteCategory(c echo.Context) error {
    id := c.Param("id")
    var category models.Category

    if err := database.DB.First(&category, id).Error; err != nil {
        return c.String(http.StatusNotFound, "Category not found")
    }

    database.DB.Delete(&category)
    return c.String(http.StatusOK, "Category deleted")
}