// Package api Code generated by swaggo/swag. DO NOT EDIT
package api

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
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
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
        "/auth": {
            "post": {
                "description": "Авторизует пользователя и возвращает JWT токен",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Авторизация пользователя",
                "parameters": [
                    {
                        "description": "Пользователь",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/hello": {
            "get": {
                "description": "Responds with a \"Hello, World!\" message",
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "hello"
                ],
                "summary": "Returns a greeting message",
                "responses": {
                    "200": {
                        "description": "Hello, World!",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/pdfReport": {
            "get": {
                "description": "Генерирует и возвращает PDF отчет",
                "produces": [
                    "application/pdf"
                ],
                "tags": [
                    "reports"
                ],
                "summary": "Получить PDF отчет",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "file"
                        }
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/register": {
            "post": {
                "description": "Handles user registration by binding and validating the user input",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Registers a new user",
                "parameters": [
                    {
                        "description": "User to register",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/sales": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Создает новую запись о продаже в системе",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "sales"
                ],
                "summary": "Создание новой продажи",
                "parameters": [
                    {
                        "description": "Информация о продаже",
                        "name": "sale",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Sale"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Успешное создание продажи",
                        "schema": {
                            "$ref": "#/definitions/model.Sale"
                        }
                    },
                    "400": {
                        "description": "Ошибка в запросе",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/welcome": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Возвращает приветственное сообщение для авторизованного пользователя",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Приветствие пользователя",
                "responses": {
                    "200": {
                        "description": "Welcome {email}!",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Sale": {
            "type": "object",
            "required": [
                "company",
                "count",
                "name_prod",
                "price"
            ],
            "properties": {
                "company": {
                    "type": "string"
                },
                "count": {
                    "type": "integer"
                },
                "name_prod": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                }
            }
        },
        "model.User": {
            "type": "object",
            "required": [
                "email",
                "name",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "maxLength": 32,
                    "minLength": 8
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "description": "Provide your JWT token as: Bearer {token}",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Swagger Example API",
	Description:      "This is a sample server for a pet store.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}