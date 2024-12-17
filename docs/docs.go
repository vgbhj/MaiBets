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
        "/api/bet": {
            "post": {
                "description": "Add a new bet for the current user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "bets"
                ],
                "summary": "AddBet",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Bet Details",
                        "name": "bet",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.BetInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Bet added successfully",
                        "schema": {
                            "$ref": "#/definitions/models.SuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid input or insufficient balance",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/bets": {
            "get": {
                "description": "Retrieve all bets for the current user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "bets"
                ],
                "summary": "GetBets",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "List of user bets",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Bet"
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid user ID",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Could not retrieve bets",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/event": {
            "post": {
                "description": "Add a new event",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "events"
                ],
                "summary": "AddEvent",
                "parameters": [
                    {
                        "description": "Event",
                        "name": "event",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Event"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Event added successfully",
                        "schema": {
                            "$ref": "#/definitions/models.SuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid input",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/event/{id}": {
            "get": {
                "description": "Retrieve a single event by its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "events"
                ],
                "summary": "GetEvent",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Event ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Event details",
                        "schema": {
                            "$ref": "#/definitions/models.Event"
                        }
                    },
                    "400": {
                        "description": "Invalid event ID",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Could not retrieve event",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/events": {
            "get": {
                "description": "Retrieve all events",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "events"
                ],
                "summary": "GetEvents",
                "responses": {
                    "200": {
                        "description": "List of events",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Event"
                            }
                        }
                    },
                    "500": {
                        "description": "Could not retrieve events",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/login": {
            "post": {
                "description": "Authenticate a user and return a JWT token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "description": "User credentials",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.AuthInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "JWT token",
                        "schema": {
                            "$ref": "#/definitions/api.TokenResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid credentials",
                        "schema": {
                            "$ref": "#/definitions/api.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/api.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/register": {
            "post": {
                "description": "Create a new user with a username and password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "CreateUser",
                "parameters": [
                    {
                        "description": "User details",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.AuthInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User created successfully",
                        "schema": {
                            "$ref": "#/definitions/api.SuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Username already exists or invalid input",
                        "schema": {
                            "$ref": "#/definitions/api.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/api.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.AuthInput": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string",
                    "example": "password123"
                },
                "username": {
                    "type": "string",
                    "example": "johndoe"
                }
            }
        },
        "api.BetInput": {
            "type": "object",
            "properties": {
                "bet_amount": {
                    "type": "number",
                    "example": 10.1
                },
                "bet_date": {
                    "type": "string",
                    "example": "2024-06-01T20:00:00Z"
                },
                "name": {
                    "type": "string",
                    "example": "Champions League Final"
                }
            }
        },
        "api.ErrorResponse": {
            "type": "object",
            "properties": {
                "details": {
                    "type": "string",
                    "example": "Optional detailed error message"
                },
                "error": {
                    "type": "string",
                    "example": "Description of the error"
                }
            }
        },
        "api.SuccessResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "User created successfully"
                }
            }
        },
        "api.TokenResponse": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string",
                    "example": "your_jwt_token_here"
                }
            }
        },
        "models.Bet": {
            "type": "object",
            "properties": {
                "bet_amount": {
                    "description": "Сумма ставки",
                    "type": "number"
                },
                "bet_date": {
                    "description": "Дата ставки",
                    "type": "string"
                },
                "bet_type_id": {
                    "description": "Идентификатор типа ставки",
                    "type": "integer"
                },
                "client_id": {
                    "description": "Идентификатор клиента",
                    "type": "integer"
                },
                "event_id": {
                    "description": "Идентификатор события",
                    "type": "integer"
                },
                "id": {
                    "description": "Идентификатор ставки",
                    "type": "integer"
                },
                "odd_id": {
                    "description": "Идентификатор коэффициента (ссылка на Odd)",
                    "type": "integer"
                },
                "status": {
                    "description": "Статус ставки (например, pending, won, lost)",
                    "type": "string"
                }
            }
        },
        "models.ErrorResponse": {
            "type": "object",
            "properties": {
                "details": {
                    "type": "string",
                    "example": "Error details here"
                },
                "error": {
                    "type": "string",
                    "example": "Invalid input"
                }
            }
        },
        "models.Event": {
            "type": "object",
            "properties": {
                "date": {
                    "type": "string",
                    "example": "2024-06-01T20:00:00Z"
                },
                "description": {
                    "type": "string",
                    "example": "Final match of the 2024 Champions League"
                },
                "id": {
                    "description": "юзаем только в GET запросах",
                    "type": "integer"
                },
                "name": {
                    "type": "string",
                    "example": "Champions League Final"
                },
                "status": {
                    "type": "string",
                    "example": "live"
                }
            }
        },
        "models.SuccessResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "Event added successfully"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}