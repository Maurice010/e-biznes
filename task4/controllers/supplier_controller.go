package controllers

import (
    "net/http"
    "task4/database"
    "task4/models"
    "github.com/labstack/echo/v4"
)

func CreateSupplier(c echo.Context) error {
    var supplier models.Supplier
    if err := c.Bind(&supplier); err != nil {
        return c.String(http.StatusBadRequest, "Invalid input")
    }
    if supplier.Name == "" {
        return c.String(http.StatusBadRequest, "Name is required")
    }
    database.DB.Create(&supplier)
    return c.JSON(http.StatusCreated, supplier)
}

func GetSuppliers(c echo.Context) error {
    var suppliers []models.Supplier
    database.DB.Find(&suppliers)
    return c.JSON(http.StatusOK, suppliers)
}

func GetSupplier(c echo.Context) error {
    id := c.Param("id")
    var supplier models.Supplier
    if err := database.DB.First(&supplier, id).Error; err != nil {
        return c.String(http.StatusNotFound, "Supplier not found")
    }
    return c.JSON(http.StatusOK, supplier)
}

func UpdateSupplier(c echo.Context) error {
    id := c.Param("id")
    var supplier models.Supplier
    if err := database.DB.First(&supplier, id).Error; err != nil {
        return c.String(http.StatusNotFound, "Supplier not found")
    }
    if err := c.Bind(&supplier); err != nil {
        return c.String(http.StatusBadRequest, "Invalid input")
    }
    if supplier.Name == "" {
        return c.String(http.StatusBadRequest, "Name is required")
    }
    database.DB.Save(&supplier)
    return c.JSON(http.StatusOK, supplier)
}

func DeleteSupplier(c echo.Context) error {
    id := c.Param("id")
    var supplier models.Supplier
    if err := database.DB.First(&supplier, id).Error; err != nil {
        return c.String(http.StatusNotFound, "Supplier not found")
    }
    database.DB.Delete(&supplier)
    return c.String(http.StatusOK, "Supplier deleted")
}