package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "online-shop/model"
)

//// @Tags admin
//// @Description Lấy danh sách User
//// @Success 200 {string} string
//// @Failure 500 {object} model.HTTPError
//// @Router /auth/get-users [get]
//func (c *Controller) GetUsers(ctx *gin.Context) {
//	ctx.String(http.StatusOK, "GetUsers")
//}
//
//// @Tags admin
//// @Description Tạo User
//// @Param user body model.CreateUser true "Thông tin tạo User"
//// @Success 200 {string} string
//// @Failure 500 {object} model.HTTPError
//// @Router /auth/create-user [post]
//func (c *Controller) CreateUser(ctx *gin.Context) {
//}
//
//// @Tags admin
//// @Description Lấy thông tin User theo ID
//// @Success 200 {string} string
//// @Failure 404 {object} model.HTTPError
//// @Failure 500 {object} model.HTTPError
//// @Router /auth/get-user/{id} [get]
//func (c *Controller) GetUserById(ctx *gin.Context) {
//}
//
//// @Tags admin
//// @Description Cập nhật User theo ID
//// @Success 200 {string} string
//// @Failure 404 {object} model.HTTPError
//// @Failure 500 {object} model.HTTPError
//// @Router /auth/update-user/{id} [put]
//func (c *Controller) UpdateUserById(ctx *gin.Context) {
//}
//
//// @Tags admin
//// @Description Xóa User theo ID
//// @Success 200 {string} string
//// @Failure 404 {object} model.HTTPError
//// @Failure 500 {object} model.HTTPError
//// @Router /auth/delete-user/{id} [delete]
//func (c *Controller) DeleteUserById(ctx *gin.Context) {
//	ctx.String(http.StatusOK, "DeleteUserById")
//}

//-----

// @Tags admin
// @Description Lấy danh sách Products
// @Success 200 {string} string
// @Failure 500 {object} model.HTTPError
// @Router /auth/get-products [get]
func (c *Controller) GetProducts(ctx *gin.Context) {
	ctx.String(http.StatusOK, "GetProducts")
}

// @Tags admin
// @Description Tạo Product
// @Param user body model.CreateProduct true "Thông tin tạo sản phẩm"
// @Success 200 {string} string
// @Failure 500 {object} model.HTTPError
// @Router /auth/create-product [post]
func (c *Controller) CreateProduct(ctx *gin.Context) {
}

// @Tags admin
// @Description Lấy thông tin Product theo ID
// @Success 200 {string} string
// @Failure 404 {object} model.HTTPError
// @Failure 500 {object} model.HTTPError
// @Router /auth/get-product/{id} [get]
func (c *Controller) GetProductById(ctx *gin.Context) {
}

// @Tags admin
// @Description Cập nhật Product theo ID
// @Success 200 {string} string
// @Failure 404 {object} model.HTTPError
// @Failure 500 {object} model.HTTPError
// @Router /auth/update-product/{id} [put]
func (c *Controller) UpdateProductById(ctx *gin.Context) {
}

// @Tags admin
// @Description Xóa Product theo ID
// @Success 200 {string} string
// @Failure 404 {object} model.HTTPError
// @Failure 500 {object} model.HTTPError
// @Router /auth/delete-product/{id} [delete]
func (c *Controller) DeleteProductById(ctx *gin.Context) {
	ctx.String(http.StatusOK, "DeleteProductById")
}