package controllers

import (
	"net/http"
    "task4/database"
    "task4/models"
    "github.com/labstack/echo/v4"
)

func AddToCart(c echo.Context) error {
	var input struct {
        UserId    uint `json:"user_id"`
        ProductId uint `json:"product_id"`
        Quantity  int  `json:"quantity"`
    }

	if err := c.Bind(&input); err != nil {
        return c.String(http.StatusBadRequest, "Invalid input")
    }

	var product models.Product
    if err := database.DB.First(&product, input.ProductId).Error; err != nil {
        return c.String(http.StatusNotFound, "Product not found")
    }

	var cart models.Cart
    if err := database.DB.Where("user_id = ?", input.UserId).First(&cart).Error; err != nil {
		cart = models.Cart{UserId: input.UserId}
        database.DB.Create(&cart)
    }

	cartItem := models.CartItem{
        CartId:    cart.ID,
        ProductId: input.ProductId,
        Quantity:  input.Quantity,
    }
    database.DB.Create(&cartItem)

	database.DB.Preload("CartItems.Product.Category").Preload("CartItems.Product.Supplier").First(&cart, cart.ID)
    return c.JSON(http.StatusOK, cart)
}

func GetCarts(c echo.Context) error {
    var carts []models.Cart
    database.DB.Preload("CartItems.Product.Category").Preload("CartItems.Product.Supplier").Find(&carts)
    return c.JSON(http.StatusOK, carts)
}

func GetCart(c echo.Context) error {
    userId := c.Param("user_id")
    var cart models.Cart
    if err := database.DB.Preload("CartItems.Product.Category").Preload("CartItems.Product.Supplier").Where("user_id = ?", userId).First(&cart).Error; err != nil {
        return c.String(http.StatusNotFound, "Cart not found")
    }
    return c.JSON(http.StatusOK, cart)
}

func UpdateCart(c echo.Context) error {
    userId := c.Param("user_id")
    var cart models.Cart
    if err := database.DB.Where("user_id = ?", userId).First(&cart).Error; err != nil {
        return c.String(http.StatusNotFound, "Cart not found")
    }

    var input struct {
        ProductId uint `json:"product_id"`
        Quantity  int  `json:"quantity"`
    }
    if err := c.Bind(&input); err != nil {
        return c.String(http.StatusBadRequest, "Invalid input")
    }

    if input.Quantity < 0 {
        return c.String(http.StatusBadRequest, "Quantity cannot be negative")
    }

    var product models.Product
    if err := database.DB.First(&product, input.ProductId).Error; err != nil {
        return c.String(http.StatusNotFound, "Product not found")
    }

    var cartItem models.CartItem
    if err := database.DB.Where("cart_id = ? AND product_id = ?", cart.ID, input.ProductId).First(&cartItem).Error; err != nil {
        if input.Quantity > 0 {
            cartItem = models.CartItem{
                CartId:    cart.ID,
                ProductId: input.ProductId,
                Quantity:  input.Quantity,
            }
            database.DB.Create(&cartItem)
        }
    } else {
        if input.Quantity == 0 {
            database.DB.Delete(&cartItem)
        } else {
            cartItem.Quantity = input.Quantity
            database.DB.Save(&cartItem)
        }
    }

    database.DB.Preload("CartItems.Product.Category").Preload("CartItems.Product.Supplier").First(&cart, cart.ID)
    return c.JSON(http.StatusOK, cart)
}

func DeleteCart(c echo.Context) error {
    userId := c.Param("user_id")
    var cart models.Cart
    if err := database.DB.Where("user_id = ?", userId).First(&cart).Error; err != nil {
        return c.String(http.StatusNotFound, "Cart not found")
    }
    database.DB.Where("cart_id = ?", cart.ID).Delete(&models.CartItem{})
    database.DB.Delete(&cart)
    return c.String(http.StatusOK, "Cart deleted")
}