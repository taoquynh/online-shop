package router

import (
	"github.com/gin-gonic/gin"
	"github.com/taoquynh/online-shop/controller"
	_ "github.com/taoquynh/online-shop/docs"
)

func setupAdminRoutes(c *controller.Controller, api *gin.RouterGroup) {
	api.GET("/get-users", c.GetUsers)
	api.POST("/create-user", c.CreateUser)
	api.GET("/get-user/:id", c.GetUserById)
	api.PUT("/update-user/:id", c.UpdateUserById)
	api.DELETE("/delete-user/:id", c.DeleteUserById)
}
