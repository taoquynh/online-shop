package model

type Product struct {
	// Tên bảng
	TableName struct{} `json:"table_name" sql:"shop.products"`

	// Mã Product (chuỗi ngẫu nhiên duy nhất)
	Id string `json:"id"`

	// Tên hiển thị
	ProductName string `json:"name" valid:"required~Tên sản phẩm không được để trống,runelength(6|100)~Tên không hợp lệ (từ 6 - 100 ký tự)"`

	// Hãng sản xuất
	Manufacture string `json:"manufacture"`

	// Giá sản phẩm
	Price int32 `json:"price"`

	// Số lượng
	Quantity int16 `json:"quantity" valid:"required~Số lượng sản phẩm không được để trống"`

	// Ảnh sản phẩm
	Image string `json:"image"`

	// Mô tả sản phẩm
	Description string `json:"description"`
}

type CreateProduct struct {
	// Tên hiển thị
	ProductName string `json:"name" valid:"required~Tên sản phẩm không được để trống,runelength(6|100)~Tên không hợp lệ (từ 6 - 100 ký tự)"`

	// Hãng sản xuất
	Manufacture string `json:"manufacture"`

	// Giá sản phẩm
	Price int32 `json:"price"`

	// Số lượng
	Quantity int16 `json:"quantity" valid:"required~Số lượng sản phẩm không được để trống"`

	// Ảnh sản phẩm
	Image string `json:"image"`

	// Mô tả sản phẩm
	Description string `json:"description"`
}

type UpdateProduct struct {
	// Mã User (chuỗi ngẫu nhiên duy nhất)
	Id string `json:"id"`

	// Tên hiển thị
	// Tên hiển thị
	ProductName string `json:"name" valid:"required~Tên sản phẩm không được để trống,runelength(6|100)~Tên không hợp lệ (từ 6 - 100 ký tự)"`

	// Hãng sản xuất
	Manufacture string `json:"manufacture"`

	// Giá sản phẩm
	Price int32 `json:"price"`

	// Số lượng
	Quantity int16 `json:"quantity" valid:"required~Số lượng sản phẩm không được để trống"`

	// Ảnh sản phẩm
	Image string `json:"image"`

	// Mô tả sản phẩm
	Description string `json:"description"`
}