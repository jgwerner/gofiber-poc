// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "email": "fiber@swagger.io"
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
        "/v1/media": {
            "post": {
                "description": "get object media",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Lists item",
                "parameters": [
                    {
                        "description": "Lists Media id",
                        "name": "ids",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entities.RequestFindByListID"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/handlers.HTTPSuccess"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/entities.Post"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handlers.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handlers.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handlers.HTTPError"
                        }
                    }
                }
            }
        },
        "/v1/media/:id": {
            "get": {
                "description": "get media by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "object media",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Media id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/handlers.HTTPSuccess"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/entities.Post"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "entities.Post": {
            "type": "object",
            "properties": {
                "_id": {
                    "description": "The identifier string of the account in the database.",
                    "type": "string",
                    "example": "61605ba183de6c07967b102d"
                },
                "author": {
                    "description": "The user identifier string.",
                    "type": "string",
                    "example": "aof"
                },
                "created_at": {
                    "description": "Last this document was created.",
                    "type": "string",
                    "example": "2021-01-27T05:44:15.337Z"
                },
                "data_version": {
                    "description": "The data version identifier string of the post in the database.",
                    "type": "string",
                    "example": "v1.0"
                },
                "meta": {
                    "$ref": "#/definitions/entities.metaPost"
                },
                "permalink": {
                    "description": "The URL in the format Media by the user.",
                    "type": "string",
                    "example": "https://example.com/222"
                },
                "post_time": {
                    "description": "Creation time of the Media.",
                    "type": "string",
                    "example": "2021-02-14T09:20:32.000Z"
                },
                "text": {
                    "description": "The content of the Media.",
                    "type": "string",
                    "example": "hello world"
                },
                "updated_at": {
                    "description": "Last this document was updated.",
                    "type": "string",
                    "example": "2021-01-28T05:44:15.337Z"
                }
            }
        },
        "entities.RequestFindByListID": {
            "type": "object",
            "required": [
                "ids"
            ],
            "properties": {
                "ids": {
                    "description": "Request body ` + "`" + `IDS` + "`" + ` are array",
                    "type": "array",
                    "items": {
                        "type": "string",
                        "maxLength": 100,
                        "minLength": 1
                    },
                    "example": [
                        "6929046813118876930",
                        "6929047870637441025"
                    ]
                }
            }
        },
        "entities.metaPost": {
            "type": "object",
            "properties": {
                "comment_count": {
                    "description": "Number of Comments of this media.",
                    "type": "integer",
                    "example": 100
                },
                "like_count": {
                    "description": "Number of Likes of this media.",
                    "type": "integer",
                    "example": 2000
                },
                "share_count": {
                    "description": "Number of Shares of this media.",
                    "type": "integer",
                    "example": 100
                },
                "type": {
                    "description": "Indicates whether this media represents a type post.",
                    "type": "string",
                    "example": "post"
                }
            }
        },
        "handlers.HTTPError": {
            "type": "object",
            "properties": {
                "error": {
                    "description": "The message error description of string.",
                    "type": "string",
                    "example": "a message error"
                },
                "status": {
                    "description": "The status error of string.",
                    "type": "string",
                    "example": "error"
                }
            }
        },
        "handlers.HTTPSuccess": {
            "type": "object",
            "properties": {
                "data": {
                    "description": "The result of json."
                },
                "status": {
                    "description": "The status success of string.",
                    "type": "string",
                    "example": "success"
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
	Host:        "localhost:3000",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "Fiber Versioning Boilerplate",
	Description: "This is a sample swagger for Fiber",
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
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
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