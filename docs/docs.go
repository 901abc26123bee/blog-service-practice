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
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/articles": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Get Article",
                "parameters": [
                    {
                        "type": "string",
                        "description": "article name",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Tag ID",
                        "name": "tag_id",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "state",
                        "name": "state",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "page",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "page size",
                        "name": "page_size",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/model.ArticleSwagger"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errcode.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errcode.Error"
                        }
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "Create Article",
                "parameters": [
                    {
                        "description": "Article ID",
                        "name": "tag_id",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "article title",
                        "name": "title",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "article description",
                        "name": "desc",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "cover image url",
                        "name": "cover_image_url",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "content",
                        "name": "content",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "created by",
                        "name": "created_by",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "integer"
                        }
                    },
                    {
                        "description": "state",
                        "name": "state",
                        "in": "body",
                        "schema": {
                            "type": "integer"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/model.Article"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errcode.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errcode.Error"
                        }
                    }
                }
            }
        },
        "/api/v1/articles/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Get an article",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Article ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/model.Article"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errcode.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errcode.Error"
                        }
                    }
                }
            },
            "put": {
                "produces": [
                    "application/json"
                ],
                "summary": "Update Article",
                "parameters": [
                    {
                        "description": "Tag ID",
                        "name": "tag_id",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "article title",
                        "name": "title",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "article description",
                        "name": "desc",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "cover image url",
                        "name": "cover_image_url",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "article content",
                        "name": "content",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "modified by",
                        "name": "modified_by",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/model.Article"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errcode.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errcode.Error"
                        }
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "summary": "Delete Article",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Article ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errcode.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errcode.Error"
                        }
                    }
                }
            }
        },
        "/api/v1/tags": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Get MultiTags",
                "parameters": [
                    {
                        "maxLength": 100,
                        "type": "string",
                        "description": "tag name",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "enum": [
                            0,
                            1
                        ],
                        "type": "integer",
                        "default": 1,
                        "description": "state",
                        "name": "state",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "page",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "page size",
                        "name": "page_size",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/model.Tag"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errcode.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errcode.Error"
                        }
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "Add Tag",
                "parameters": [
                    {
                        "maxLength": 100,
                        "minLength": 3,
                        "description": "tag name",
                        "name": "name",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "enum": [
                            0,
                            1
                        ],
                        "default": 1,
                        "description": "state",
                        "name": "state",
                        "in": "body",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    {
                        "maxLength": 100,
                        "minLength": 3,
                        "description": "created by",
                        "name": "created_by",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/model.Tag"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errcode.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errcode.Error"
                        }
                    }
                }
            }
        },
        "/api/v1/tags/{id}": {
            "put": {
                "produces": [
                    "application/json"
                ],
                "summary": "Update Tag",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Tag ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "maxLength": 100,
                        "minLength": 3,
                        "description": "tag name",
                        "name": "name",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "enum": [
                            0,
                            1
                        ],
                        "default": 1,
                        "description": "state",
                        "name": "state",
                        "in": "body",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    {
                        "maxLength": 100,
                        "minLength": 3,
                        "description": "modified by",
                        "name": "modified_by",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Tag"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errcode.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errcode.Error"
                        }
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "summary": "Delete tag",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Tag ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errcode.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errcode.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "app.Pager": {
            "type": "object",
            "properties": {
                "page": {
                    "type": "integer"
                },
                "page_size": {
                    "type": "integer"
                },
                "total_rows": {
                    "type": "integer"
                }
            }
        },
        "errcode.Error": {
            "type": "object"
        },
        "model.Article": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "cover_image_url": {
                    "type": "string"
                },
                "created_by": {
                    "type": "string"
                },
                "created_on": {
                    "type": "integer"
                },
                "deleted_on": {
                    "type": "integer"
                },
                "desc": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "is_del": {
                    "type": "integer"
                },
                "modified_by": {
                    "type": "string"
                },
                "modified_on": {
                    "type": "integer"
                },
                "state": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "model.ArticleSwagger": {
            "type": "object",
            "properties": {
                "list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Article"
                    }
                },
                "pager": {
                    "$ref": "#/definitions/app.Pager"
                }
            }
        },
        "model.Tag": {
            "type": "object",
            "properties": {
                "created_by": {
                    "type": "string"
                },
                "created_on": {
                    "type": "integer"
                },
                "deleted_on": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "is_del": {
                    "type": "integer"
                },
                "modified_by": {
                    "type": "string"
                },
                "modified_on": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "state": {
                    "type": "integer"
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
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "Blog Service",
	Description: "Go Blog Service",
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
	swag.Register("swagger", &s{})
}
