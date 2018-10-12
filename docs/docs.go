// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2018-10-12 16:19:19.377504 +0700 +07 m=+0.028757085

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
        "/auth/create-product": {
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
        "/auth/delete-product/{id}": {
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
        "/auth/get-product/{id}": {
            "get": {
                "description": "Lấy thông tin Product theo ID",
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
        "/auth/get-products": {
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
        "/auth/update-product/{id}": {
            "put": {
                "description": "Cập nhật Product theo ID",
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
                "manufacture": {
                    "description": "Hãng sản xuất",
                    "type": "string"
                },
                "price": {
                    "description": "Giá sản phẩm",
                    "type": "integer"
                },
                "product_name": {
                    "description": "Tên hiển thị",
                    "type": "string"
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
