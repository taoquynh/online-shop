package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	_ "online-shop/model"
)

// @Tags admin
// @Description Lấy danh sách User
// @Success 200 {string} string
// @Failure 500 {object} model.HTTPError
// @Router /shop/get-users [get]
func (c *Controller) GetUsers(ctx *gin.Context) {
	ctx.String(http.StatusOK, "GetUsers")
}

// @Tags admin
// @Description Register User
// @Param user body model.CreateUser true "Thông tin Register"
// @Success 200 {string} string
// @Failure 500 {object} model.HTTPError
// @Router /shop/create-user [post]
func (c *Controller) CreateUser(ctx *gin.Context) {

}

// @Tags admin
// @Description Login User theo ID
// @Success 200 {string} string
// @Failure 404 {object} model.HTTPError
// @Failure 500 {object} model.HTTPError
// @Router /shop/login-user/{id} [get]
func (c *Controller) GetLogInUserById(ctx *gin.Context) {
}

//-----

// @Tags admin
// @Description Lấy danh sách Products
// @Success 200 {string} string
// @Failure 500 {object} model.HTTPError
// @Router /shop/get-products [get]
func (c *Controller) GetProducts(ctx *gin.Context) {
	ctx.String(http.StatusOK, "GetProducts")
}

// @Tags admin
// @Description Tạo Product
// @Param user body model.CreateProduct true "Thông tin tạo sản phẩm"
// @Success 200 {string} string
// @Failure 500 {object} model.HTTPError
// @Router /shop/create-product [post]
func (c *Controller) CreateProduct(ctx *gin.Context) {
}

// @Tags admin
// @Description Lấy thông tin Product theo ID
// @Success 200 {string} string
// @Failure 404 {object} model.HTTPError
// @Failure 500 {object} model.HTTPError
// @Router /shop/get-product/{id} [get]
func (c *Controller) GetProductById(ctx *gin.Context) {
}

// @Tags admin
// @Description Cập nhật Product theo ID
// @Success 200 {string} string
// @Failure 404 {object} model.HTTPError
// @Failure 500 {object} model.HTTPError
// @Router /shop/update-product/{id} [put]
func (c *Controller) UpdateProductById(ctx *gin.Context) {
}

// @Tags admin
// @Description Xóa Product theo ID
// @Success 200 {string} string
// @Failure 404 {object} model.HTTPError
// @Failure 500 {object} model.HTTPError
// @Router /shop/delete-product/{id} [delete]
func (c *Controller) DeleteProductById(ctx *gin.Context) {
	ctx.String(http.StatusOK, "DeleteProductById")
}