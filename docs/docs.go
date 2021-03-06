// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2018-10-15 17:09:45.773351 +0700 +07 m=+0.032042081

package docs

import (
	"bytes"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/shop/create-product": {
            "post": {
                "description": "Tạo Product",
                "tags": [
                    "admin"
                ],
                "parameters": [
                    {
                        "description": "Thông tin tạo sản phẩm",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.CreateProduct"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.HTTPError"
                        }
                    }
                }
            }
        },
        "/shop/create-user": {
            "post": {
                "description": "Register User",
                "tags": [
                    "admin"
                ],
                "parameters": [
                    {
                        "description": "Thông tin Register",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.CreateUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.HTTPError"
                        }
                    }
                }
            }
        },
        "/shop/delete-product/{id}": {
            "delete": {
                "description": "Xóa Product theo ID",
                "tags": [
                    "admin"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.HTTPError"
                        }
                    }
                }
            }
        },
        "/shop/get-product/{id}": {
            "get": {
                "description": "Lấy thông tin Product theo ID",
                "tags": [
                    "admin"
                ],
                "parameters": [
                    {
                        "description": "Sản phẩm",
                        "name": "product",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.GetProductById"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.HTTPError"
                        }
                    }
                }
            }
        },
        "/shop/get-products": {
            "get": {
                "description": "Lấy danh sách Products",
                "tags": [
                    "admin"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.HTTPError"
                        }
                    }
                }
            }
        },
        "/shop/get-users": {
            "get": {
                "description": "Lấy danh sách User",
                "tags": [
                    "admin"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.HTTPError"
                        }
                    }
                }
            }
        },
        "/shop/login-user": {
            "post": {
                "description": "Login User theo username",
                "tags": [
                    "admin"
                ],
                "parameters": [
                    {
                        "description": "Đăng nhập",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.UserLogIn"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.HTTPError"
                        }
                    }
                }
            }
        },
        "/shop/update-product/{id}": {
            "put": {
                "description": "Cập nhật Product theo ID",
                "tags": [
                    "admin"
                ],
                "parameters": [
                    {
                        "description": "Thông tin sản phẩm",
                        "name": "product",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.UpdateProductById"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.HTTPError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.CreateProduct": {
            "type": "object",
            "properties": {
                "description": {
                    "description": "Mô tả sản phẩm",
                    "type": "string"
                },
                "image": {
                    "description": "Ảnh sản phẩm",
                    "type": "string"
                },
                "manufacturer": {
                    "description": "Hãng sản xuất",
                    "type": "string"
                },
                "name": {
                    "description": "Tên hiển thị",
                    "type": "string"
                },
                "price": {
                    "description": "Giá sản phẩm",
                    "type": "number"
                },
                "quantity": {
                    "description": "Số lượng",
                    "type": "integer"
                }
            }
        },
        "model.CreateUser": {
            "type": "object",
            "properties": {
                "password": {
                    "description": "Mật khẩu phải có tối thiểu 8 ký tự",
                    "type": "string"
                },
                "username": {
                    "description": "Tên hiển thị",
                    "type": "string"
                }
            }
        },
        "model.GetProductById": {
            "type": "object",
            "properties": {
                "description": {
                    "description": "Mô tả sản phẩm",
                    "type": "string"
                },
                "id": {
                    "description": "Mã User (chuỗi ngẫu nhiên duy nhất)",
                    "type": "string"
                },
                "image": {
                    "description": "Ảnh sản phẩm",
                    "type": "string"
                },
                "manufacturer": {
                    "description": "Hãng sản xuất",
                    "type": "string"
                },
                "name": {
                    "description": "Tên hiển thị",
                    "type": "string"
                },
                "price": {
                    "description": "Giá sản phẩm",
                    "type": "number"
                },
                "quantity": {
                    "description": "Số lượng",
                    "type": "integer"
                }
            }
        },
        "model.HTTPError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "model.UpdateProductById": {
            "type": "object",
            "properties": {
                "description": {
                    "description": "Mô tả sản phẩm",
                    "type": "string"
                },
                "id": {
                    "description": "Mã User (chuỗi ngẫu nhiên duy nhất)",
                    "type": "string"
                },
                "image": {
                    "description": "Ảnh sản phẩm",
                    "type": "string"
                },
                "manufacturer": {
                    "description": "Hãng sản xuất",
                    "type": "string"
                },
                "name": {
                    "description": "Tên hiển thị",
                    "type": "string"
                },
                "price": {
                    "description": "Giá sản phẩm",
                    "type": "number"
                },
                "quantity": {
                    "description": "Số lượng",
                    "type": "integer"
                }
            }
        },
        "model.UserLogIn": {
            "type": "object",
            "properties": {
                "password": {
                    "description": "Mật khẩu phải có tối thiểu 8 ký tự",
                    "type": "string"
                },
                "username": {
                    "description": "Tên hiển thị",
                    "type": "string"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo swaggerInfo

type s struct{}

func (s *s) ReadDoc() string {
	t, err := template.New("swagger_info").Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, SwaggerInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
