// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/datafabric",
        "contact": {
            "name": "API Support datafabric",
            "url": "http://www.swagger.io/support.datafabric",
            "email": "support@swagger.io datafabric"
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
        "/api/sensorConfig/serviceType": {
            "get": {
                "description": "Get service type",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "sensorConfig"
                ],
                "summary": "Get service type",
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dto.GetServiceTypeDto"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.GetServiceTypeDto": {
            "type": "object",
            "properties": {
                "serviceType": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BasicAuth": {
            "type": "basic"
        }
    },
    "externalDocs": {
        "description": "OpenAPI",
        "url": "https://swagger.io/resources/open-api/"
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8000",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Swagger datafabric API",
	Description:      "datafabricThis is a sample server celler server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
