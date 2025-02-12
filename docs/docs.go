// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "https://github.com/barganakukuhraditya",
            "email": "714220013@std.ulbi.ac.id"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/delete/{id}": {
            "delete": {
                "description": "Hapus data parfume.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Parfume"
                ],
                "summary": "Delete data parfume.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Masukan ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/hapus/{id}": {
            "delete": {
                "description": "Hapus data user.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Delete data user.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Masukan ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/insert": {
            "post": {
                "description": "Input data parfume.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Parfume"
                ],
                "summary": "Insert data parfume.",
                "parameters": [
                    {
                        "description": "Payload Body [RAW]",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.Parfume"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Parfume"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/parfume": {
            "get": {
                "description": "Mengambil semua data parfume.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Parfume"
                ],
                "summary": "Get All Data Parfume.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Parfume"
                        }
                    }
                }
            }
        },
        "/parfume/{id}": {
            "get": {
                "description": "Ambil per ID data parfume.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Parfume"
                ],
                "summary": "Get By ID Data Parfume.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Masukan ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Parfume"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/post": {
            "post": {
                "description": "Input data user.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Insert data user.",
                "parameters": [
                    {
                        "description": "Payload Body [RAW]",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.User"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/put/{id}": {
            "put": {
                "description": "Ubah data user.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Update data user.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Masukan ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Payload Body [RAW]",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.User"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/update/{id}": {
            "put": {
                "description": "Ubah data parfume.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Parfume"
                ],
                "summary": "Update data parfume.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Masukan ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Payload Body [RAW]",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.ReqParfume"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Parfume"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/user": {
            "get": {
                "description": "Mengambil semua data user.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get All Data User.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.User"
                        }
                    }
                }
            }
        },
        "/user/{id}": {
            "get": {
                "description": "Ambil per ID data user.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get By ID Data User.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Masukan ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.User"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        }
    },
    "definitions": {
        "controller.Parfume": {
            "type": "object",
            "properties": {
                "_id": {
                    "type": "string",
                    "example": "123456789"
                },
                "deskripsi": {
                    "type": "string",
                    "example": "Parfum yang sangat wangi"
                },
                "harga": {
                    "type": "integer",
                    "example": 1000000
                },
                "jenis_parfume": {
                    "type": "string",
                    "example": "Eau de Parfum"
                },
                "merk": {
                    "type": "string",
                    "example": "Dior"
                },
                "nama_parfume": {
                    "type": "string",
                    "example": "Chirstian Dior"
                },
                "stok": {
                    "type": "integer",
                    "example": 100
                },
                "tahun_peluncuran": {
                    "type": "integer",
                    "example": 2021
                },
                "ukuran": {
                    "type": "string",
                    "example": "100ml"
                }
            }
        },
        "controller.ReqParfume": {
            "type": "object",
            "properties": {
                "deskripsi": {
                    "type": "string",
                    "example": "Parfum yang sangat wangi"
                },
                "harga": {
                    "type": "integer",
                    "example": 1000000
                },
                "jenis_parfume": {
                    "type": "string",
                    "example": "Eau de Parfum"
                },
                "merk": {
                    "type": "string",
                    "example": "Dior"
                },
                "nama_parfume": {
                    "type": "string",
                    "example": "Chirstian Dior"
                },
                "stok": {
                    "type": "integer",
                    "example": 100
                },
                "tahun_peluncuran": {
                    "type": "integer",
                    "example": 2021
                },
                "ukuran": {
                    "type": "string",
                    "example": "100ml"
                }
            }
        },
        "controller.User": {
            "type": "object",
            "properties": {
                "_id": {
                    "type": "string",
                    "example": "123456789"
                },
                "address": {
                    "type": "string",
                    "example": "Jl. Jalan"
                },
                "email": {
                    "type": "string",
                    "example": "user"
                },
                "idrole": {
                    "type": "string",
                    "example": "123456789"
                },
                "password": {
                    "type": "string",
                    "example": "user"
                },
                "phone": {
                    "type": "string",
                    "example": "08123456789"
                },
                "username": {
                    "type": "string",
                    "example": "user"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "tb-parfume2024-34a7b650de40.herokuapp.com",
	BasePath:         "/",
	Schemes:          []string{"https", "http"},
	Title:            "SWAGGER TUGAS BESAR",
	Description:      "This is a sample swagger for Fiber",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
