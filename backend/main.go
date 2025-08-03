package main

import (
    "ecommerce-backend/config"
    "ecommerce-backend/routes"
    "ecommerce-backend/models"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    config.ConnectDB()

    // Auto migrate models
    config.DB.AutoMigrate(&models.User{}, &models.Item{}, &models.Cart{}, &models.Order{})

    routes.RegisterRoutes(r)
    r.Run(":8080")
}

