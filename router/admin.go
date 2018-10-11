package router

import (
	"github.com/gin-gonic/gin"
	"online-shop/controller"
	_ "online-shop/docs"
)

func setupAdminRoutes(c *controller.Controller, api *gin.RouterGroup) {
	api.GET("/get-users", c.GetUsers)
	api.POST("/create-user", c.CreateUser)
	api.GET("/get-user/:id", c.GetUserById)
	api.PUT("/update-user/:id", c.UpdateUserById)
	api.DELETE("/delete-user/:id", c.DeleteUserById)
}
