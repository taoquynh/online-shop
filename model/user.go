package model

type User struct {
	// Tên bảng
	TableName struct{} `json:"table_name" sql:"shop.users"`

	// Mã User (chuỗi ngẫu nhiên duy nhất)
	Id string `json:"id"`

	// Tên hiển thị
	FullName string `json:"full_name" valid:"required~Họ tên không được để trống,runelength(6|100)~Họ tên không hợp lệ (từ 6 - 100 ký tự)"`

	// Địa chỉ email
	Email string `json:"email" sql:",unique" valid:"email~Email không hợp lệ"`

	// Mật khẩu phải có tối thiểu 8 ký tự
	Password string `json:"password" valid:"required~Mật khẩu không được để trống,runelength(8|32)~Mật khẩu phải từ 8 - 32 ký tự"`

	// Số điện thoại
	Phone string `json:"phone" sql:",unique" valid:"numeric~Số điện thoại không hợp lệ, runelength(8|20)~Số điện thoại phải từ 8 - 20 ký tự"`
}

type CreateUser struct {
	// Tên hiển thị
	FullName string `json:"full_name" valid:"required~Họ tên không được để trống,runelength(6|100)~Họ tên không hợp lệ (từ 6 - 100 ký tự)"`

	// Địa chỉ email
	Email string `json:"email" sql:",unique"`

	// Mật khẩu phải có tối thiểu 8 ký tự
	Password string `json:"password" valid:"required~Mật khẩu không được để trống,runelength(8|32)~Mật khẩu phải từ 8 - 32 ký tự"`

	// Số điện thoại
	Phone string `json:"phone" sql:",unique" valid:"numeric~Số điện thoại không hợp lệ, runelength(8|20)~Số điện thoại phải từ 8 - 20 ký tự"`
}

type UpdateUser struct {
	// Mã User (chuỗi ngẫu nhiên duy nhất)
	Id string `json:"id"`

	// Tên hiển thị
	FullName string `json:"full_name" valid:"required~Họ tên không được để trống,runelength(6|100)~Họ tên không hợp lệ (từ 6 - 100 ký tự)"`

	// Địa chỉ email
	Email string `json:"email" sql:",unique"`

	// Số điện thoại
	Phone string `json:"phone" sql:",unique" valid:"numeric~Số điện thoại không hợp lệ, runelength(8|20)~Số điện thoại phải từ 8 - 20 ký tự"`
}

