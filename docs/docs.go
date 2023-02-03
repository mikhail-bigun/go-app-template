// Code generated by swaggo/swag. DO NOT EDIT
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
        "/users": {
            "get": {
                "description": "Get users with options",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get users",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "offset",
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "less than unix timestamp",
                        "name": "lt_datetime",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "greater than unix timestamp",
                        "name": "gt_datetime",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Users"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/fiber.Error"
                        }
                    }
                }
            },
            "post": {
                "description": "Create new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Create new user",
                "parameters": [
                    {
                        "description": "create User request",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UserBase"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "headers": {
                            "Location": {
                                "type": "string",
                                "description": "/api/users/:id"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/validator.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/fiber.Error"
                        }
                    }
                }
            }
        },
        "/users/{id}": {
            "get": {
                "description": "Get user by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user id",
                        "name": "id",
                        "in": "path",
                        "required": true
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
                            "$ref": "#/definitions/fiber.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/fiber.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/fiber.Error"
                        }
                    }
                }
            },
            "put": {
                "description": "Update existing user or create a new one",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Update user",
                "parameters": [
                    {
                        "description": "Update User request",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    },
                    {
                        "type": "string",
                        "description": "user id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "headers": {
                            "Location": {
                                "type": "string",
                                "description": "/api/users/:id"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/validator.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/fiber.Error"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete user by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Delete user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/fiber.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/fiber.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/fiber.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "app_pkg_validator.FieldError": {
            "type": "object",
            "properties": {
                "failed_field": {
                    "type": "string"
                },
                "tag": {
                    "type": "string"
                },
                "value": {
                    "type": "string"
                }
            }
        },
        "fiber.Error": {
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
        "model.User": {
            "type": "object",
            "required": [
                "email",
                "first_name"
            ],
            "properties": {
                "created_at": {
                    "type": "string",
                    "example": "2023-01-07T18:09:15.237672Z"
                },
                "email": {
                    "type": "string",
                    "example": "example@email.com"
                },
                "first_name": {
                    "type": "string",
                    "maxLength": 255,
                    "minLength": 3,
                    "example": "John"
                },
                "id": {
                    "type": "string",
                    "example": "31afbd8d-0312-4f18-87ee-24d5881a619e"
                },
                "last_name": {
                    "type": "string",
                    "maxLength": 255,
                    "minLength": 3,
                    "example": "Doe"
                },
                "updated_at": {
                    "type": "string",
                    "example": "2023-01-07T18:09:15.237672Z"
                }
            }
        },
        "model.UserBase": {
            "type": "object",
            "required": [
                "email",
                "first_name"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "example": "example@email.com"
                },
                "first_name": {
                    "type": "string",
                    "maxLength": 255,
                    "minLength": 3,
                    "example": "John"
                },
                "last_name": {
                    "type": "string",
                    "maxLength": 255,
                    "minLength": 3,
                    "example": "Doe"
                }
            }
        },
        "model.Users": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.User"
                    }
                },
                "limit": {
                    "type": "integer",
                    "example": 100
                },
                "offset": {
                    "type": "integer",
                    "example": 100
                },
                "total": {
                    "type": "integer",
                    "example": 1000
                }
            }
        },
        "validator.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "boolean"
                },
                "field_error": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/app_pkg_validator.FieldError"
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "127.0.0.1:8000",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Swagger Example API",
	Description:      "This is a sample service template.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
