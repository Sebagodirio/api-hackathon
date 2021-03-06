// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
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
            "name": "Sebastian Godirio",
            "email": "sebastianlgodirio@gmail.com"
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
        "/hackathons/": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "It returns all the hackathons with its delopments, which contains its developers",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Token for entering to the endpoint",
                        "name": "BearerAuthorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.HackathonResponse"
                            }
                        }
                    },
                    "503": {
                        "description": "Service Unavailable",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/hackathons/top-ten/": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "It returns the top 10 developments of all the hackathons",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Token for entering to the endpoint",
                        "name": "BearerAuthorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.DevelopmentResponse"
                            }
                        }
                    },
                    "503": {
                        "description": "Service Unavailable",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "It returns all the hackathons with its delopments, which contains its developers",
                "parameters": [
                    {
                        "description": "Email and password of the user",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.User"
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
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "503": {
                        "description": "Service Unavailable",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.DevelopersResponse": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                }
            }
        },
        "models.DevelopmentResponse": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "developers": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.DevelopersResponse"
                    }
                },
                "name": {
                    "type": "string"
                },
                "position": {
                    "type": "string"
                },
                "score": {
                    "type": "number"
                }
            }
        },
        "models.HackathonResponse": {
            "type": "object",
            "properties": {
                "country": {
                    "type": "string"
                },
                "developments": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.DevelopmentResponse"
                    }
                },
                "hackathon_year": {
                    "type": "string"
                },
                "state": {
                    "type": "string"
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
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
	Version:          "2.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "api-hackathon",
	Description:      "API Restful for getting the best developments in the hackathons",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
