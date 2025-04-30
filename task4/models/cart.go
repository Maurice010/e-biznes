package models

import "gorm.io/gorm"

type Cart struct {
    gorm.Model
    UserId	uint	`json:"user_id"`
    cartItems	[]CartItem	`json:"cart_items" gorm:"foreignKey:CartId"`
}