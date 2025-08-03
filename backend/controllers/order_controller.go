package controllers

import (
    "ecommerce-backend/config"
    "ecommerce-backend/models"
    "github.com/gin-gonic/gin"
    "net/http"
)

func CreateOrder(c *gin.Context) {
    user := c.MustGet("user").(models.User)

    var cart models.Cart
    if err := config.DB.Preload("Items").Where("user_id = ?", user.ID).First(&cart).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Cart not found"})
        return
    }

    order := models.Order{UserID: user.ID, Items: cart.Items}
    config.DB.Create(&order)

    config.DB.Model(&cart).Association("Items").Clear()

    c.JSON(http.StatusOK, order)
}

func ListOrders(c *gin.Context) {
    var orders []models.Order
    config.DB.Preload("Items").Find(&orders)
    c.JSON(http.StatusOK, orders)
}
