package controllers

import (
    "ecommerce-backend/config"
    "ecommerce-backend/models"
    "github.com/gin-gonic/gin"
    "net/http"
)

type AddToCartInput struct {
    ItemID uint `json:"item_id"`
}

func AddToCart(c *gin.Context) {
    var input AddToCartInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user := c.MustGet("user").(models.User)

    var cart models.Cart
    if err := config.DB.Where("user_id = ?", user.ID).First(&cart).Error; err != nil {
        cart = models.Cart{UserID: user.ID}
        config.DB.Create(&cart)
    }

    var item models.Item
    config.DB.First(&item, input.ItemID)

    config.DB.Model(&cart).Association("Items").Append(&item)

    c.JSON(http.StatusOK, cart)
}

func ListCarts(c *gin.Context) {
    var carts []models.Cart
    config.DB.Preload("Items").Find(&carts)
    c.JSON(http.StatusOK, carts)
}
