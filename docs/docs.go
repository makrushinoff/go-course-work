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
        "/authenticate": {
            "post": {
                "description": "authenticate user in the system",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "get token",
                "parameters": [
                    {
                        "description": "Submit",
                        "name": "authenticationForm",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UserDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Token"
                        }
                    }
                }
            }
        },
        "/boards": {
            "get": {
                "description": "retrieves all saved board entities from database",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "boards"
                ],
                "summary": "get all possible boards",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.BoardDto"
                            }
                        }
                    }
                }
            }
        },
        "/boards/{boardId}": {
            "get": {
                "description": "respond board dto by provided id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "boards"
                ],
                "summary": "get possible board by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "use id",
                        "name": "boardId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.BoardDto"
                        }
                    }
                }
            }
        },
        "/calculations": {
            "post": {
                "description": "receives a 2D array from request body and search on it a circle loop, depends on default conditions",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "calculations"
                ],
                "summary": "find circle loop",
                "parameters": [
                    {
                        "description": "Submit",
                        "name": "calculations",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.BoardDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.BoardDto"
                        }
                    }
                }
            }
        },
        "/register": {
            "post": {
                "description": "registers new user with provided date",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "register as new user",
                "parameters": [
                    {
                        "description": "Submit",
                        "name": "registrationForm",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UserDto"
                        }
                    }
                ],
                "responses": {}
            }
        }
    },
    "definitions": {
        "model.BoardDto": {
            "type": "object",
            "properties": {
                "board": {
                    "type": "array",
                    "items": {
                        "type": "array",
                        "items": {
                            "type": "integer"
                        }
                    }
                },
                "id": {
                    "type": "integer"
                }
            }
        },
        "model.Token": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "model.UserDto": {
            "type": "object",
            "properties": {
                "login": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "localhost:8080",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Lab3 Rest API",
	Description:      "Rest API documentation, generated based on annotations and swag library",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}