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
        "/disputegames": {
            "get": {
                "description": "Get all dispute game by page",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get dispute games",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "page num",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "page size",
                        "name": "size",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/disputegames/:address/claimdatas": {
            "get": {
                "description": "Get all claim data by address",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get claim data",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "dispute game address",
                        "name": "address",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/disputegames/:address/credit": {
            "get": {
                "description": "Get credit details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get credit details",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "account address",
                        "name": "address",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/disputegames/count": {
            "get": {
                "description": "Get dispute games count group by status and per day",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "GetCountDisputeGameGroupByStatus",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "today before ? days",
                        "name": "days",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/api.CountGames"
                            }
                        }
                    }
                }
            }
        },
        "/disputegames/credit/rank": {
            "get": {
                "description": "Get credit rank",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get credit rank",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "rank length limit number",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/disputegames/overview": {
            "get": {
                "description": "Get overview",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "overview",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.Overview"
                        }
                    }
                }
            }
        },
        "/disputegames/overview/amountperdays": {
            "get": {
                "description": "Get amount per day",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "GetAmountPerDays",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "today before ? days",
                        "name": "days",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/api.AmountPerDay"
                            }
                        }
                    }
                }
            }
        },
        "/disputegames/statistics/bond/inprogress": {
            "get": {
                "description": "Get bond in progress per days",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "GetBondInProgressPerDays",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/api.AmountPerDay"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.AmountPerDay": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "string"
                },
                "date": {
                    "type": "string"
                }
            }
        },
        "api.CountGames": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "string"
                },
                "date": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "api.Overview": {
            "type": "object",
            "properties": {
                "challengerWinGamesCount": {
                    "type": "integer"
                },
                "defenderWinWinGamesCount": {
                    "type": "integer"
                },
                "disputeGameProxy": {
                    "type": "string"
                },
                "inProgressGamesCount": {
                    "type": "integer"
                },
                "totalCredit": {
                    "type": "string"
                },
                "totalGames": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
