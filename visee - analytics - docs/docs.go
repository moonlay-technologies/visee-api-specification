// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
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
        "/analytic/buyers/{client_uuid}/count": {
            "get": {
                "description": "Get count buyers analytics",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Buyers"
                ],
                "summary": "Get count buyers analytics",
                "operationId": "count-buyers-analytics",
                "parameters": [
                    {
                        "type": "string",
                        "description": "client uuid",
                        "name": "client_uuid",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "collectionFormat": "multi",
                        "description": "zone uuid",
                        "name": "filter_zone",
                        "in": "query"
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "collectionFormat": "multi",
                        "description": "gender",
                        "name": "filter_gender",
                        "in": "query"
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "collectionFormat": "multi",
                        "description": "age",
                        "name": "filter_age",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "YYYY-MM-DD",
                        "name": "range_from",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "YYYY-MM-DD",
                        "name": "range_to",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/models.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/models.BuyersCountResponse"
                                        },
                                        "message": {
                                            "type": "array",
                                            "items": {
                                                "type": "string"
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
                            "allOf": [
                                {
                                    "$ref": "#/definitions/models.BasicResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "message": {
                                            "type": "array",
                                            "items": {
                                                "type": "string"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/models.BasicResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "message": {
                                            "type": "array",
                                            "items": {
                                                "type": "string"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/analytic/history-state/{client_uuid}/avg-dw-time-by-period": {
            "get": {
                "description": "Get avg dwell time by time period analytics",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "History State"
                ],
                "summary": "Get avg dwell time by time period analytics",
                "operationId": "avg-dwell-time-by-time-period-analytics",
                "parameters": [
                    {
                        "type": "string",
                        "description": "client uuid",
                        "name": "client_uuid",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "collectionFormat": "multi",
                        "description": "zone uuid",
                        "name": "filter_zone",
                        "in": "query"
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "collectionFormat": "multi",
                        "description": "gender",
                        "name": "filter_gender",
                        "in": "query"
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "collectionFormat": "multi",
                        "description": "age",
                        "name": "filter_age",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "YYYY-MM-DD",
                        "name": "range_from",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "YYYY-MM-DD",
                        "name": "range_to",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/models.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/models.AvgDwTimeByPeriod"
                                        },
                                        "message": {
                                            "type": "array",
                                            "items": {
                                                "type": "string"
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
                            "allOf": [
                                {
                                    "$ref": "#/definitions/models.BasicResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "message": {
                                            "type": "array",
                                            "items": {
                                                "type": "string"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/models.BasicResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "message": {
                                            "type": "array",
                                            "items": {
                                                "type": "string"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/analytic/visitor/{client_uuid}/count": {
            "get": {
                "description": "Get count visitor analytics",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Visitor"
                ],
                "summary": "Get count visitor analytics",
                "operationId": "count-visitor-analytics",
                "parameters": [
                    {
                        "type": "string",
                        "description": "client uuid",
                        "name": "client_uuid",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "collectionFormat": "multi",
                        "description": "zone uuid",
                        "name": "filter_zone",
                        "in": "query"
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "collectionFormat": "multi",
                        "description": "gender",
                        "name": "filter_gender",
                        "in": "query"
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "collectionFormat": "multi",
                        "description": "age",
                        "name": "filter_age",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "YYYY-MM-DD",
                        "name": "range_from",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "YYYY-MM-DD",
                        "name": "range_to",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/models.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        " data": {
                                            "$ref": "#/definitions/models.VisitorCountResponse"
                                        },
                                        "message": {
                                            "type": "array",
                                            "items": {
                                                "type": "string"
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
                            "allOf": [
                                {
                                    "$ref": "#/definitions/models.BasicResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "message": {
                                            "type": "array",
                                            "items": {
                                                "type": "string"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/models.BasicResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "message": {
                                            "type": "array",
                                            "items": {
                                                "type": "string"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/analytic/visitor/{client_uuid}/count-group": {
            "get": {
                "description": "Get count max, min, avg visitor analytics",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Visitor"
                ],
                "summary": "Get count max, min, avg visitor analytics",
                "operationId": "count-visitor-analytic-group",
                "parameters": [
                    {
                        "type": "string",
                        "description": "client uuid",
                        "name": "client_uuid",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "collectionFormat": "multi",
                        "description": "zone uuid",
                        "name": "filter_zone",
                        "in": "query"
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "collectionFormat": "multi",
                        "description": "gender",
                        "name": "filter_gender",
                        "in": "query"
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "collectionFormat": "multi",
                        "description": "age",
                        "name": "filter_age",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "YYYY-MM-DD",
                        "name": "range_from",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "YYYY-MM-DD",
                        "name": "range_to",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/models.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        " data": {
                                            "$ref": "#/definitions/models.VisitorCount"
                                        },
                                        "message": {
                                            "type": "array",
                                            "items": {
                                                "type": "string"
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
                            "allOf": [
                                {
                                    "$ref": "#/definitions/models.BasicResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "message": {
                                            "type": "array",
                                            "items": {
                                                "type": "string"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/models.BasicResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "message": {
                                            "type": "array",
                                            "items": {
                                                "type": "string"
                                            }
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
        "models.AvgDwTimeByPeriod": {
            "type": "object",
            "properties": {
                "avg_dw_time": {
                    "type": "number"
                },
                "time_period": {
                    "type": "string"
                }
            }
        },
        "models.BasicResponse": {
            "type": "object",
            "properties": {
                "message": {},
                "status_code": {
                    "type": "integer"
                }
            }
        },
        "models.BuyersCountResponse": {
            "type": "object",
            "properties": {
                "number_of_buyers": {
                    "type": "integer"
                }
            }
        },
        "models.Response": {
            "type": "object",
            "properties": {
                "data": {},
                "message": {},
                "status_code": {
                    "type": "integer"
                }
            }
        },
        "models.SumCountVisitor": {
            "type": "object",
            "properties": {
                "label": {
                    "type": "string"
                },
                "sum_count": {
                    "type": "integer"
                }
            }
        },
        "models.VisitorCount": {
            "type": "object",
            "properties": {
                "avg_visitor": {
                    "type": "number"
                },
                "max_visitor": {
                    "$ref": "#/definitions/models.SumCountVisitor"
                },
                "min_visitor": {
                    "$ref": "#/definitions/models.SumCountVisitor"
                }
            }
        },
        "models.VisitorCountResponse": {
            "type": "object",
            "properties": {
                "number_of_visitor": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:3009",
	BasePath:         "/v1",
	Schemes:          []string{"http"},
	Title:            "Echo Swagger Example API",
	Description:      "This is a sample server server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
