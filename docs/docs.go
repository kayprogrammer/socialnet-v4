// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "consumes": [
        "application/json"
    ],
    "produces": [
        "application/json"
    ],
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
        "/auth/register": {
            "post": {
                "description": "This endpoint registers new users into our application.",
                "tags": [
                    "Auth"
                ],
                "summary": "Register a new user",
                "parameters": [
                    {
                        "description": "User object",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schemas.RegisterUser"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/schemas.RegisterResponseSchema"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/auth/resend-verification-email": {
            "post": {
                "description": "This endpoint resends new otp to the user's email.",
                "tags": [
                    "Auth"
                ],
                "summary": "Resend Verification Email",
                "parameters": [
                    {
                        "description": "Email object",
                        "name": "email",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schemas.EmailRequestSchema"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schemas.ResponseSchema"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/auth/send-password-reset-otp": {
            "post": {
                "description": "This endpoint sends new password reset otp to the user's email.",
                "tags": [
                    "Auth"
                ],
                "summary": "Send Password Reset Otp",
                "parameters": [
                    {
                        "description": "Email object",
                        "name": "email",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schemas.EmailRequestSchema"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schemas.ResponseSchema"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorResponse"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/auth/set-new-password": {
            "post": {
                "description": "This endpoint verifies the password reset otp.",
                "tags": [
                    "Auth"
                ],
                "summary": "Set New Password",
                "parameters": [
                    {
                        "description": "Password reset object",
                        "name": "email",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schemas.SetNewPasswordSchema"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schemas.ResponseSchema"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorResponse"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/auth/verify-email": {
            "post": {
                "description": "This endpoint verifies a user's email.",
                "tags": [
                    "Auth"
                ],
                "summary": "Verify a user's email",
                "parameters": [
                    {
                        "description": "Verify Email object",
                        "name": "verify_email",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schemas.VerifyEmailRequestSchema"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schemas.ResponseSchema"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/general/site-detail": {
            "get": {
                "description": "This endpoint retrieves few details of the site/application.",
                "tags": [
                    "General"
                ],
                "summary": "Retrieve site details",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schemas.SiteDetailResponseSchema"
                        }
                    }
                }
            }
        },
        "/healthcheck": {
            "get": {
                "description": "This endpoint checks the health of our application.",
                "tags": [
                    "HealthCheck"
                ],
                "summary": "HealthCheck",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/routes.HealthCheckSchema"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "routes.HealthCheckSchema": {
            "type": "object",
            "properties": {
                "success": {
                    "type": "string",
                    "example": "pong"
                }
            }
        },
        "schemas.EmailRequestSchema": {
            "type": "object",
            "required": [
                "email"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "minLength": 5,
                    "example": "johndoe@email.com"
                }
            }
        },
        "schemas.RegisterResponseSchema": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/schemas.EmailRequestSchema"
                },
                "message": {
                    "type": "string",
                    "example": "Data fetched/created/updated/deleted"
                },
                "status": {
                    "type": "string",
                    "example": "success"
                }
            }
        },
        "schemas.RegisterUser": {
            "type": "object",
            "required": [
                "email",
                "first_name",
                "last_name",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "minLength": 5,
                    "example": "johndoe@email.com"
                },
                "first_name": {
                    "type": "string",
                    "maxLength": 50,
                    "example": "John"
                },
                "last_name": {
                    "type": "string",
                    "maxLength": 50,
                    "example": "Doe"
                },
                "password": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 8,
                    "example": "strongpassword"
                },
                "terms_agreement": {
                    "type": "boolean"
                }
            }
        },
        "schemas.ResponseSchema": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "Data fetched/created/updated/deleted"
                },
                "status": {
                    "type": "string",
                    "example": "success"
                }
            }
        },
        "schemas.SetNewPasswordSchema": {
            "type": "object",
            "required": [
                "email",
                "otp",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "minLength": 5,
                    "example": "johndoe@example.com"
                },
                "otp": {
                    "type": "integer",
                    "example": 123456
                },
                "password": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 8,
                    "example": "newstrongpassword"
                }
            }
        },
        "schemas.SiteDetail": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string",
                    "example": "234, Lagos, Nigeria"
                },
                "email": {
                    "type": "string",
                    "example": "johndoe@email.com"
                },
                "fb": {
                    "type": "string",
                    "example": "https://facebook.com"
                },
                "ig": {
                    "type": "string",
                    "example": "https://instagram.com"
                },
                "name": {
                    "type": "string",
                    "example": "SocialNet"
                },
                "phone": {
                    "type": "string",
                    "example": "+2348133831036"
                },
                "tw": {
                    "type": "string",
                    "example": "https://twitter.com"
                },
                "wh": {
                    "type": "string",
                    "example": "https://wa.me/2348133831036"
                }
            }
        },
        "schemas.SiteDetailResponseSchema": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/schemas.SiteDetail"
                },
                "message": {
                    "type": "string",
                    "example": "Data fetched/created/updated/deleted"
                },
                "status": {
                    "type": "string",
                    "example": "success"
                }
            }
        },
        "schemas.VerifyEmailRequestSchema": {
            "type": "object",
            "required": [
                "email",
                "otp"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "minLength": 5,
                    "example": "johndoe@example.com"
                },
                "otp": {
                    "type": "integer",
                    "example": 123456
                }
            }
        },
        "utils.ErrorResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "data": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "description": "Type 'Bearer jwt_string' to correctly set the API Key",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "4.0",
	Host:             "",
	BasePath:         "/api/v4",
	Schemes:          []string{},
	Title:            "SOCIALNET API",
	Description:      "A Realtime Social Networking API built with Fiber",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
