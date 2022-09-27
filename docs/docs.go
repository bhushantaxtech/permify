// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "API Support",
            "url": "https://github.com/Permify/permify/issues",
            "email": "hello@permify.co"
        },
        "license": {
            "name": "GNU Affero General Public License v3.0"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/permissions/check": {
            "post": {
                "description": "check subject is authorized",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Permission"
                ],
                "summary": "Permission",
                "operationId": "permissions.check",
                "parameters": [
                    {
                        "description": "check subject is authorized",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/permission.CheckRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/permission.CheckResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/common.HTTPErrorResponse"
                        }
                    }
                }
            }
        },
        "/permissions/expand": {
            "post": {
                "description": "expand relationships according to schema",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Permission"
                ],
                "summary": "Permission",
                "operationId": "permissions.expand",
                "parameters": [
                    {
                        "description": "expand relationships according to schema",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/permission.ExpandRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/permission.ExpandResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/common.HTTPErrorResponse"
                        }
                    }
                }
            }
        },
        "/relationships/delete": {
            "post": {
                "description": "delete relation tuple",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Relationship"
                ],
                "summary": "Relationship",
                "operationId": "relationships.delete",
                "parameters": [
                    {
                        "description": "delete relation tuple",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/relationship.DeleteRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/tuple.Tuple"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/common.HTTPErrorResponse"
                        }
                    }
                }
            }
        },
        "/relationships/read": {
            "post": {
                "description": "read relation tuple(s)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Relationship"
                ],
                "summary": "Relationship",
                "operationId": "relationships.read",
                "parameters": [
                    {
                        "description": "read relation tuple(s)",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/relationship.ReadRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/tuple.Tuple"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/common.HTTPErrorResponse"
                        }
                    }
                }
            }
        },
        "/relationships/write": {
            "post": {
                "description": "create new relation tuple",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Relationship"
                ],
                "summary": "Relationship",
                "operationId": "relationships.write",
                "parameters": [
                    {
                        "description": "create new relation tuple",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/relationship.WriteRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/tuple.Tuple"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/common.HTTPErrorResponse"
                        }
                    }
                }
            }
        },
        "/schemas/lookup": {
            "post": {
                "description": "lookup your authorization model",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Schema"
                ],
                "summary": "Schema",
                "operationId": "schemas.lookup",
                "parameters": [
                    {
                        "description": "lookup your authorization model",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schema.LookupRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schema.LookupResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/common.HTTPErrorResponse"
                        }
                    }
                }
            }
        },
        "/schemas/read": {
            "post": {
                "description": "read your authorization model",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Schema"
                ],
                "summary": "Schema",
                "operationId": "schemas.read",
                "parameters": [
                    {
                        "description": "read your authorization model",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schema.ReadRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schema.ReadResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/common.HTTPErrorResponse"
                        }
                    }
                }
            }
        },
        "/schemas/write": {
            "post": {
                "description": "write your authorization model",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Schema"
                ],
                "summary": "Schema",
                "operationId": "schemas.write",
                "parameters": [
                    {
                        "description": "schema file (expected extension .perm)",
                        "name": "schema",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "integer"
                            }
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schema.WriteResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/common.HTTPErrorResponse"
                        }
                    }
                }
            }
        },
        "/status/ping": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Server"
                ],
                "summary": "Server",
                "operationId": "ping",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.Message"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/common.HTTPErrorResponse"
                        }
                    }
                }
            }
        },
        "/status/version": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Server"
                ],
                "summary": "Server",
                "operationId": "version",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.Message"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/common.HTTPErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "common.HTTPErrorResponse": {
            "type": "object",
            "properties": {
                "errors": {}
            }
        },
        "common.Message": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "filters.EntityFilter": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "filters.RelationTupleFilter": {
            "type": "object",
            "properties": {
                "entity": {
                    "$ref": "#/definitions/filters.EntityFilter"
                },
                "relation": {
                    "type": "string"
                },
                "subject": {
                    "$ref": "#/definitions/filters.SubjectFilter"
                }
            }
        },
        "filters.SubjectFilter": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "relation": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "permission.CheckRequest": {
            "type": "object",
            "properties": {
                "action": {
                    "type": "string"
                },
                "depth": {
                    "type": "integer"
                },
                "entity": {
                    "$ref": "#/definitions/tuple.Entity"
                },
                "schema_version": {
                    "type": "string"
                },
                "subject": {
                    "$ref": "#/definitions/tuple.Subject"
                }
            }
        },
        "permission.CheckResponse": {
            "type": "object",
            "properties": {
                "can": {
                    "type": "boolean"
                },
                "decisions": {},
                "remaining_depth": {
                    "type": "integer"
                }
            }
        },
        "permission.ExpandRequest": {
            "type": "object",
            "properties": {
                "action": {
                    "type": "string"
                },
                "entity": {
                    "$ref": "#/definitions/tuple.Entity"
                },
                "schema_version": {
                    "type": "string"
                }
            }
        },
        "permission.ExpandResponse": {
            "type": "object",
            "properties": {
                "tree": {}
            }
        },
        "relationship.DeleteRequest": {
            "type": "object",
            "properties": {
                "entity": {
                    "$ref": "#/definitions/tuple.Entity"
                },
                "relation": {
                    "type": "string"
                },
                "subject": {
                    "$ref": "#/definitions/tuple.Subject"
                }
            }
        },
        "relationship.ReadRequest": {
            "type": "object",
            "properties": {
                "filter": {
                    "$ref": "#/definitions/filters.RelationTupleFilter"
                }
            }
        },
        "relationship.WriteRequest": {
            "type": "object",
            "properties": {
                "entity": {
                    "$ref": "#/definitions/tuple.Entity"
                },
                "relation": {
                    "type": "string"
                },
                "schema_version": {
                    "type": "string"
                },
                "subject": {
                    "$ref": "#/definitions/tuple.Subject"
                }
            }
        },
        "schema.Action": {
            "type": "object",
            "properties": {
                "child": {},
                "name": {
                    "type": "string"
                }
            }
        },
        "schema.Entity": {
            "type": "object",
            "properties": {
                "actions": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/schema.Action"
                    }
                },
                "name": {
                    "type": "string"
                },
                "option": {
                    "type": "object",
                    "additionalProperties": true
                },
                "relations": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/schema.Relation"
                    }
                }
            }
        },
        "schema.LookupRequest": {
            "type": "object",
            "properties": {
                "entity_type": {
                    "type": "string"
                },
                "relation_names": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "schema_version": {
                    "type": "string"
                }
            }
        },
        "schema.LookupResponse": {
            "type": "object",
            "properties": {
                "action_names": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "schema.ReadRequest": {
            "type": "object",
            "properties": {
                "schema_version": {
                    "type": "string"
                }
            }
        },
        "schema.ReadResponse": {
            "type": "object",
            "properties": {
                "entities": {
                    "type": "object",
                    "additionalProperties": {
                        "$ref": "#/definitions/schema.Entity"
                    }
                }
            }
        },
        "schema.Relation": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "option": {
                    "type": "object",
                    "additionalProperties": true
                },
                "type": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "schema.WriteResponse": {
            "type": "object",
            "properties": {
                "version": {
                    "type": "string"
                }
            }
        },
        "tuple.Entity": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "tuple.Subject": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "relation": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "tuple.Tuple": {
            "type": "object",
            "properties": {
                "entity": {
                    "$ref": "#/definitions/tuple.Entity"
                },
                "relation": {
                    "type": "string"
                },
                "subject": {
                    "$ref": "#/definitions/tuple.Subject"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "api-key-auth",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "v0.0.0-alpha4",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{"http", "https"},
	Title:            "Permify API",
	Description:      "Permify is an open-source authorization service for creating and maintaining fine-grained authorizations accross your individual applications and services.\nPermify converts authorization data as relational tuples into a database you point at. We called that database a Write Database (WriteDB) and it behaves as a centralized data source for your authorization system. You can model of your authorization with Permify's DSL - Permify Schema - and perform access checks with a single API call anywhere on your stack. Access decisions made according to stored relational tuples.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
