// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/categories": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Categories"
                ],
                "summary": "Creates a new category",
                "parameters": [
                    {
                        "description": "Category object to be created",
                        "name": "model.Category",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CategoryCreate"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Category created successfully",
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/helper.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/helper.ErrorResponse"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/helper.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Server Error",
                        "schema": {
                            "$ref": "#/definitions/helper.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/categories/{categoryId}": {
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Categories"
                ],
                "summary": "Update a Category.",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Category ID",
                        "name": "categoryId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Category object to be updated",
                        "name": "model.Category",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CategoryUpdate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Category updated successfully",
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/helper.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/helper.ErrorResponse"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/helper.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Server Error",
                        "schema": {
                            "$ref": "#/definitions/helper.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/users/delete-account": {
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Delete an Account.",
                "responses": {
                    "200": {
                        "description": "User deleted successfully",
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/helper.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/helper.ErrorResponse"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/helper.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Server Error",
                        "schema": {
                            "$ref": "#/definitions/helper.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/users/login": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Login an Account.",
                "parameters": [
                    {
                        "description": "User login credentials",
                        "name": "model.User",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.LoginCredential"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User logged in successfully",
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/helper.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/helper.ErrorResponse"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/helper.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Server Error",
                        "schema": {
                            "$ref": "#/definitions/helper.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/users/register": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Register an Account.",
                "parameters": [
                    {
                        "description": "User object to be created",
                        "name": "model.User",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UserCreate"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "User created successfully",
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/helper.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/helper.ErrorResponse"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/helper.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Server Error",
                        "schema": {
                            "$ref": "#/definitions/helper.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/users/update-account": {
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Update an Account.",
                "parameters": [
                    {
                        "description": "User object to be updated",
                        "name": "model.User",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UserUpdate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User updated successfully",
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/helper.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/helper.ErrorResponse"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/helper.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Server Error",
                        "schema": {
                            "$ref": "#/definitions/helper.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "helper.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "model.CategoryCreate": {
            "type": "object",
            "properties": {
                "type": {
                    "type": "string"
                }
            }
        },
        "model.CategoryUpdate": {
            "type": "object",
            "properties": {
                "type": {
                    "type": "string"
                }
            }
        },
        "model.LoginCredential": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "model.Task": {
            "type": "object",
            "properties": {
                "category_id": {
                    "type": "integer"
                },
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "status": {
                    "type": "boolean"
                },
                "title": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/model.User"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "model.User": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "full_name": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "password": {
                    "type": "string"
                },
                "photos": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Task"
                    }
                },
                "role": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "model.UserCreate": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "full_name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "model.UserUpdate": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "full_name": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Kanban Board",
	Description:      "This is a simple server mimicking Instagram server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}