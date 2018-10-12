package model

type User struct {
	// Tên bảng
	TableName struct{} `json:"table_name" sql:"shop.users"`

	// Mã User (chuỗi ngẫu nhiên duy nhất)
	Id string `json:"id"`

	// Tên hiển thị
	UserName string `json:"user_name" valid:"required~Họ tên không được để trống,runelength(6|100)~Họ tên không hợp lệ (từ 6 - 100 ký tự)"`

	// Mật khẩu phải có tối thiểu 8 ký tự
	Password string `json:"password" valid:"required~Mật khẩu không được để trống,runelength(8|32)~Mật khẩu phải từ 8 - 32 ký tự"`
}

type CreateUser struct {
	// Tên hiển thị
	UserName string `json:"user_name" valid:"required~Họ tên không được để trống,runelength(6|100)~Họ tên không hợp lệ (từ 6 - 100 ký tự)"`

	// Mật khẩu phải có tối thiểu 8 ký tự
	Password string `json:"password" valid:"required~Mật khẩu không được để trống,runelength(8|32)~Mật khẩu phải từ 8 - 32 ký tự"`

}

type UserLogin struct {
	// Tên hiển thị
	UserName string `json:"user_name" valid:"required~Họ tên không được để trống,runelength(6|100)~Họ tên không hợp lệ (từ 6 - 100 ký tự)"`

	// Mật khẩu phải có tối thiểu 8 ký tự
	Password string `json:"password" valid:"required~Mật khẩu không được để trống,runelength(8|32)~Mật khẩu phải từ 8 - 32 ký tự"`

}


