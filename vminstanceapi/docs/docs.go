// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "email": "parameswaribala@gmail.com"
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
        "/vminstances/v1.0": {
            "get": {
                "description": "Get details of all vmInstances",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "vmInstances"
                ],
                "summary": "Get details of all vmInstances",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/stores.VMConfiguration"
                            }
                        }
                    }
                }
            },
            "put": {
                "description": "Update existing vmInstance with the input vmInstance",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "vminstances"
                ],
                "summary": "Update existing vmInstance",
                "parameters": [
                    {
                        "description": "Create customer",
                        "name": "vmInstance",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/stores.VMConfiguration"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/stores.VMConfiguration"
                        }
                    }
                }
            }
        },
        "/vminstances/v1.0/{vmName}": {
            "delete": {
                "description": "Delete vmInstance by name",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "vmInstances"
                ],
                "summary": "Delete vmInstance",
                "parameters": [
                    {
                        "type": "string",
                        "description": "VMName of the VMConfiguration",
                        "name": "vmName",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/stores.VMConfiguration"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "stores.VMConfiguration": {
            "type": "object",
            "properties": {
                "access_key": {
                    "type": "string"
                },
                "cpu_count": {
                    "type": "integer"
                },
                "instance_type": {
                    "type": "string"
                },
                "provider": {
                    "type": "string"
                },
                "ram": {
                    "type": "string"
                },
                "regions": {
                    "type": "string"
                },
                "secret_key": {
                    "type": "string"
                },
                "vm_name gorm:": {
                    "type": "string"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "localhost:7072",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "VMConfiguration API",
	Description: "This is api service for managing VM Configuration",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
