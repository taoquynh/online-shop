package router

import (
	"github.com/gin-gonic/gin"
	"online-shop/controller"
	_ "online-shop/docs"
)

func setupAdminRoutes(c *controller.Controller, api *gin.RouterGroup) {
	api.GET("/get-users", c.GetUsers)
	api.POST("/create-user", c.CreateUser)
	api.GET("/get-login/:id", c.GetLogInUserById)


	api.GET("/get-products", c.GetProducts)
	api.POST("/create-product", c.CreateProduct)
	api.GET("/get-products/:id", c.GetProductById)
	api.PUT("/update-products/:id", c.UpdateProductById)
	api.DELETE("/delete-product/:id", c.DeleteProductById)
}
