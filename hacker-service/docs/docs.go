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
        "/GetReports": {
            "get": {
                "description": "Check reports of hacked matrixes",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Additional methods"
                ],
                "summary": "Reports",
                "responses": {}
            }
        },
        "/HelpHack": {
            "post": {
                "description": "Get ready answers how to hack matrixes of matrix_service",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Main tools"
                ],
                "summary": "HelpHack",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id of the matrix desired to be hacked",
                        "name": "matrix_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {}
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "localhost:8081",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Utility to solve matrices of matrix_service",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
