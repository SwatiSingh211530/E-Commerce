
package routes

import (
    "github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
    r.POST("/users", controllers.CreateUser)
    r.POST("/users/login", controllers.LoginUser)
    r.GET("/users", controllers.ListUsers)

    r.POST("/items", controllers.CreateItem)
    r.GET("/items", controllers.ListItems)

    r.POST("/carts", middleware.AuthMiddleware(), controllers.AddToCart)
    r.GET("/carts", controllers.ListCarts)

    r.POST("/orders", middleware.AuthMiddleware(), controllers.CreateOrder)
    r.GET("/orders", controllers.ListOrders)
}
