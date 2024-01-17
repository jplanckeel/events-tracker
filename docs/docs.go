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
        "/api/v1/events/": {
            "get": {
                "description": "List all Events",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "event"
                ],
                "summary": "List all Events",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/services.EventOutput"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "Create an Event",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "event"
                ],
                "summary": "Create an Event",
                "parameters": [
                    {
                        "description": "The Request body",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.CreateDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/services.EventOutput"
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
        }
    },
    "definitions": {
        "event.EventAttributes": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "deployment service lambda version v0.0.0"
                },
                "priority": {
                    "type": "string",
                    "example": "P1"
                },
                "related_id": {
                    "type": "string",
                    "example": "53aa35c8-e659-44b2-882f-f6056e443c99"
                },
                "service": {
                    "type": "string",
                    "example": "service-event"
                },
                "source": {
                    "type": "string",
                    "example": "github_action"
                },
                "status": {
                    "type": "string",
                    "example": "success"
                },
                "type": {
                    "type": "string",
                    "example": "event"
                }
            }
        },
        "event.EventLinks": {
            "type": "object",
            "properties": {
                "pull_request_link": {
                    "type": "string",
                    "example": "https://github.com/jplanckeel/events-tracker/pull/240"
                }
            }
        },
        "handlers.CreateDTO": {
            "type": "object",
            "properties": {
                "attributes": {
                    "$ref": "#/definitions/event.EventAttributes"
                },
                "links": {
                    "$ref": "#/definitions/event.EventLinks"
                },
                "title": {
                    "type": "string",
                    "example": "Deployment service lambda"
                }
            }
        },
        "services.EventOutput": {
            "type": "object"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Events-Tracker API",
	Description:      "This API was built to stock all events in infrastructure",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
