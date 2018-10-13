package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/rs/xid"
	"log"
	"net/http"
	"online-shop/model"
	_ "online-shop/model"
)

// @Tags admin
// @Description Register User
// @Param user body model.CreateUser true "Thông tin Register"
// @Success 200 {string} string
// @Failure 500 {object} model.HTTPError
// @Router /shop/create-user [post]
func (c *Controller) CreateUser(ctx *gin.Context) {
	var user model.CreateUser
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		model.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	var newUser model.User
	copier.Copy(&newUser, &user)
	newUser.Id = xid.New().String()
	err = c.DB.Insert(&newUser)
	if err != nil {
		log.Println((err.Error()))
		model.NewError(ctx, http.StatusBadRequest, errors.New("Khong the tao user"))
		return

	}

	ctx.String(http.StatusOK, "tao thanh cong")
}
// @Tags admin
// @Description Lấy danh sách User
// @Success 200 {string} string
// @Failure 500 {object} model.HTTPError
// @Router /shop/get-users [get]
func (c *Controller) GetUsers(ctx *gin.Context) {
	var users []model.GetUsers
	_, err := c.DB.Query(&users, `SELECT * FROM shop.users`)
	if err != nil {
		log.Println(err)
		model.NewError(ctx, http.StatusNotFound, errors.New("Khong co user"))
		return
	}
	ctx.JSON(http.StatusOK, users)
}

// @Tags admin
// @Description Login User theo username
// @Param user body model.UserLogIn true "Dang nhap"
// @Success 200 {string} string
// @Failure 404 {object} model.HTTPError
// @Failure 500 {object} model.HTTPError
// @Router /shop/login-user/ [post]
func (c *Controller) UserLogIn(ctx *gin.Context) {
	var login model.UserLogIn
	err := ctx.ShouldBindJSON(&login)
	if err != nil {
		model.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	username := ctx.Param("username")
	password := ctx.Param("password")

	var user model.UserLogIn
	_, erro := c.DB.Query(&user, `SELECT FROM shop.users WHERE username = ? AND password = ?`, username, password)
	if erro != nil {
		model.NewError(ctx, http.StatusNotFound, errors.New("Tai khoan khong dung"))
		return
	}

	ctx.String(http.StatusOK, "Dang nhap thanh cong")

}

//-----



// @Tags admin
// @Description Tạo Product
// @Param user body model.CreateProduct true "Thông tin tạo sản phẩm"
// @Success 200 {string} string
// @Failure 500 {object} model.HTTPError
// @Router /shop/create-product [post]
func (c *Controller) CreateProduct(ctx *gin.Context) {
	var product model.CreateProduct
	err := ctx.ShouldBindJSON(&product)
	if err != nil {
		model.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	var newProduct model.Product
	/*newProduct.Description = product.Description
	newProduct.Image = product.Image
	newProduct.Manufacture = product.Manufacture
	newProduct.Price = product.Price
	newProduct.ProductName = product.ProductName
	newProduct.Quantity = product.Quantity*/
	copier.Copy(&newProduct, &product)
	newProduct.Id = xid.New().String()
	err = c.DB.Insert(&newProduct)
	if err != nil {
		model.NewError(ctx, http.StatusBadRequest, errors.New("Khong the tao product"))
			return

	}

	ctx.String(http.StatusOK, "thanh cong")
}

// @Tags admin
// @Description Cập nhật Product theo ID
// @Param product body model.UpdateProductById true "Thông tin san pham"
// @Success 200 {string} string
// @Failure 404 {object} model.HTTPError
// @Failure 500 {object} model.HTTPError
// @Router /shop/update-product/{id} [put]
func (c *Controller) UpdateProductById(ctx *gin.Context) {
	var product model.UpdateProductById
	err := ctx.ShouldBindJSON(&product)
	if err != nil {
		model.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	var updateProduct model.Product
	copier.Copy(&updateProduct, &product)
	err = c.DB.Update(&updateProduct)
	if err != nil {
		model.NewError(ctx, http.StatusBadRequest, errors.New("Khong the cap nhat san pham"))
		return

	}

	ctx.String(http.StatusOK, "cap nhat thanh cong")
}

// @Tags admin
// @Description Lấy danh sách Products
// @Success 200 {string} string
// @Failure 500 {object} model.HTTPError
// @Router /shop/get-products [get]
func (c *Controller) GetProducts(ctx *gin.Context) {
	var products []model.GetProducts
	_, err := c.DB.Query(&products, `SELECT * FROM shop.products`)
	if err != nil {
		model.NewError(ctx, http.StatusNotFound, errors.New("Khong co product"))
		return
	}
	ctx.JSON(http.StatusOK, products)
}


// @Tags admin
// @Description Lấy thông tin Product theo ID
// @Param product body model.GetProductById true "San pham"
// @Success 200 {string} string
// @Failure 404 {object} model.HTTPError
// @Failure 500 {object} model.HTTPError
// @Router /shop/get-product/{id} [get]
func (c *Controller) GetProductById(ctx *gin.Context) {
	id := ctx.Param("id")

	var products []model.GetProductById
	_, err := c.DB.Query(&products, `SELECT * FROM shop.products WHERE id = ?`, id)
	if err != nil {
		model.NewError(ctx, http.StatusNotFound, errors.New("Khong co product"))
		return
	}
	ctx.JSON(http.StatusOK, products)
}


// @Tags admin
// @Description Xóa Product theo ID
// @Success 200 {string} string
// @Failure 404 {object} model.HTTPError
// @Failure 500 {object} model.HTTPError
// @Router /shop/delete-product/{id} [delete]
func (c *Controller) DeleteProductById(ctx *gin.Context) {

	id := ctx.Param("id")

	_, err := c.DB.Exec(`DELETE FROM shop.products WHERE id = ?`, id)
	if err != nil {
		model.NewError(ctx, http.StatusNotFound, errors.New("Khong co product"))
		return
	}

	ctx.String(http.StatusOK, "Xoa thanh cong")
}

// c.DB.Exec(`SELECT FROM shop.users WHERE`)