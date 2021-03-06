// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "swagger": "2.0",
  "info": {
    "title": "GatewayUser",
    "version": "1.0.0"
  },
  "basePath": "/",
  "paths": {
    "/login": {
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "auth"
        ],
        "operationId": "LoginUser",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/LoginRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successful response",
            "schema": {
              "$ref": "#/definitions/LoginResponse"
            }
          },
          "401": {
            "description": "Unauthorized"
          }
        }
      }
    },
    "/movies": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "movies"
        ],
        "operationId": "ListMovies",
        "parameters": [
          {
            "enum": [
              0,
              1
            ],
            "type": "integer",
            "default": 0,
            "description": "Movie category (0 - action; 1 - horror)",
            "name": "category",
            "in": "query"
          },
          {
            "type": "integer",
            "description": "Movie release year start",
            "name": "year_start",
            "in": "query"
          },
          {
            "type": "integer",
            "description": "Movie release year end",
            "name": "year_end",
            "in": "query"
          },
          {
            "type": "integer",
            "name": "limit",
            "in": "query",
            "required": true
          },
          {
            "type": "integer",
            "name": "offset",
            "in": "query"
          },
          {
            "type": "boolean",
            "name": "only_rented",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "Successful response",
            "schema": {
              "$ref": "#/definitions/ListMoviesResponse"
            }
          }
        }
      }
    },
    "/movies/{movieId}": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "movies"
        ],
        "operationId": "GetMovie",
        "parameters": [
          {
            "type": "integer",
            "format": "int32",
            "name": "movieId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Successful response",
            "schema": {
              "$ref": "#/definitions/Movie"
            }
          }
        }
      }
    },
    "/movies/{movieId}/rent": {
      "post": {
        "security": [
          {
            "bearerAuth": []
          }
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "movies"
        ],
        "operationId": "RentMovie",
        "parameters": [
          {
            "type": "integer",
            "format": "int32",
            "name": "movieId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Successful response",
            "schema": {
              "$ref": "#/definitions/RentMovieResponse"
            }
          },
          "400": {
            "description": "Not enough balance"
          }
        }
      }
    },
    "/profile": {
      "get": {
        "security": [
          {
            "bearerAuth": []
          }
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "profile"
        ],
        "operationId": "GetProfile",
        "responses": {
          "200": {
            "description": "Successful response",
            "schema": {
              "$ref": "#/definitions/Profile"
            }
          }
        }
      }
    },
    "/profile/orders": {
      "get": {
        "security": [
          {
            "bearerAuth": []
          }
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "profile"
        ],
        "operationId": "ListOrders",
        "responses": {
          "200": {
            "description": "Successful response",
            "schema": {
              "$ref": "#/definitions/ListOrdersResponse"
            }
          }
        }
      }
    },
    "/profile/payments": {
      "get": {
        "security": [
          {
            "bearerAuth": []
          }
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "profile"
        ],
        "operationId": "ListPayments",
        "responses": {
          "200": {
            "description": "Successful response",
            "schema": {
              "$ref": "#/definitions/ListPaymentsResponse"
            }
          }
        }
      }
    },
    "/register": {
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "auth"
        ],
        "operationId": "RegisterUser",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/RegisterRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successful response",
            "schema": {
              "$ref": "#/definitions/RegisterResponse"
            }
          },
          "400": {
            "description": "UserExists"
          }
        }
      }
    }
  },
  "definitions": {
    "ListMoviesResponse": {
      "type": "object",
      "properties": {
        "PageNum": {
          "type": "integer"
        },
        "PageSize": {
          "type": "integer"
        },
        "Payload": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Movie"
          }
        }
      }
    },
    "ListOrdersResponse": {
      "type": "object",
      "properties": {
        "PageNum": {
          "type": "integer"
        },
        "PageSize": {
          "type": "integer"
        },
        "Payload": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Order"
          }
        }
      }
    },
    "ListPaymentsResponse": {
      "type": "object",
      "properties": {
        "PageNum": {
          "type": "integer"
        },
        "PageSize": {
          "type": "integer"
        },
        "Payload": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Payment"
          }
        }
      }
    },
    "LoginRequest": {
      "type": "object",
      "properties": {
        "Email": {
          "type": "string"
        },
        "Password": {
          "type": "string"
        }
      }
    },
    "LoginResponse": {
      "type": "object",
      "properties": {
        "JWT": {
          "type": "string"
        }
      }
    },
    "Movie": {
      "type": "object",
      "properties": {
        "Category": {
          "type": "integer"
        },
        "ID": {
          "type": "integer",
          "format": "int32"
        },
        "ReleaseYear": {
          "type": "string"
        },
        "Title": {
          "type": "string"
        }
      }
    },
    "Order": {
      "type": "object",
      "properties": {
        "Amount": {
          "type": "number",
          "format": "float64"
        },
        "CreatedAt": {
          "type": "integer",
          "format": "int64"
        },
        "ID": {
          "type": "integer",
          "format": "int32"
        },
        "MovieID": {
          "type": "integer"
        }
      }
    },
    "Payment": {
      "type": "object",
      "properties": {
        "Amount": {
          "type": "number",
          "format": "float64"
        },
        "CreatedAt": {
          "type": "integer",
          "format": "int64"
        },
        "ID": {
          "type": "integer",
          "format": "int32"
        },
        "Status": {
          "type": "integer"
        },
        "TransactionID": {
          "type": "integer",
          "format": "int32"
        },
        "UpdatedAt": {
          "type": "integer",
          "format": "int64"
        }
      }
    },
    "Profile": {
      "type": "object",
      "properties": {
        "Age": {
          "type": "integer"
        },
        "Balance": {
          "type": "number",
          "format": "float64"
        },
        "DisplayName": {
          "type": "string"
        },
        "Email": {
          "type": "string"
        },
        "ID": {
          "type": "integer"
        },
        "Phone": {
          "type": "string"
        },
        "Role": {
          "type": "string"
        },
        "Status": {
          "type": "integer"
        }
      }
    },
    "RegisterRequest": {
      "type": "object",
      "properties": {
        "Age": {
          "type": "integer"
        },
        "DisplayName": {
          "type": "string"
        },
        "Email": {
          "type": "string"
        },
        "Password": {
          "type": "string"
        },
        "Phone": {
          "type": "string"
        }
      }
    },
    "RegisterResponse": {
      "type": "object",
      "properties": {
        "JWT": {
          "type": "string"
        }
      }
    },
    "RentMovieResponse": {
      "type": "object",
      "properties": {
        "Payload": {
          "type": "object",
          "properties": {
            "MovieID": {
              "type": "integer",
              "format": "int32"
            },
            "RentEndTime": {
              "type": "integer",
              "format": "int64"
            }
          }
        }
      }
    }
  },
  "securityDefinitions": {
    "bearerAuth": {
      "type": "apiKey",
      "name": "Authorization",
      "in": "header"
    }
  },
  "x-components": {}
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "swagger": "2.0",
  "info": {
    "title": "GatewayUser",
    "version": "1.0.0"
  },
  "basePath": "/",
  "paths": {
    "/login": {
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "auth"
        ],
        "operationId": "LoginUser",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/LoginRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successful response",
            "schema": {
              "$ref": "#/definitions/LoginResponse"
            }
          },
          "401": {
            "description": "Unauthorized"
          }
        }
      }
    },
    "/movies": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "movies"
        ],
        "operationId": "ListMovies",
        "parameters": [
          {
            "enum": [
              0,
              1
            ],
            "type": "integer",
            "default": 0,
            "description": "Movie category (0 - action; 1 - horror)",
            "name": "category",
            "in": "query"
          },
          {
            "type": "integer",
            "description": "Movie release year start",
            "name": "year_start",
            "in": "query"
          },
          {
            "type": "integer",
            "description": "Movie release year end",
            "name": "year_end",
            "in": "query"
          },
          {
            "type": "integer",
            "name": "limit",
            "in": "query",
            "required": true
          },
          {
            "type": "integer",
            "name": "offset",
            "in": "query"
          },
          {
            "type": "boolean",
            "name": "only_rented",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "Successful response",
            "schema": {
              "$ref": "#/definitions/ListMoviesResponse"
            }
          }
        }
      }
    },
    "/movies/{movieId}": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "movies"
        ],
        "operationId": "GetMovie",
        "parameters": [
          {
            "type": "integer",
            "format": "int32",
            "name": "movieId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Successful response",
            "schema": {
              "$ref": "#/definitions/Movie"
            }
          }
        }
      }
    },
    "/movies/{movieId}/rent": {
      "post": {
        "security": [
          {
            "bearerAuth": []
          }
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "movies"
        ],
        "operationId": "RentMovie",
        "parameters": [
          {
            "type": "integer",
            "format": "int32",
            "name": "movieId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Successful response",
            "schema": {
              "$ref": "#/definitions/RentMovieResponse"
            }
          },
          "400": {
            "description": "Not enough balance"
          }
        }
      }
    },
    "/profile": {
      "get": {
        "security": [
          {
            "bearerAuth": []
          }
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "profile"
        ],
        "operationId": "GetProfile",
        "responses": {
          "200": {
            "description": "Successful response",
            "schema": {
              "$ref": "#/definitions/Profile"
            }
          }
        }
      }
    },
    "/profile/orders": {
      "get": {
        "security": [
          {
            "bearerAuth": []
          }
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "profile"
        ],
        "operationId": "ListOrders",
        "responses": {
          "200": {
            "description": "Successful response",
            "schema": {
              "$ref": "#/definitions/ListOrdersResponse"
            }
          }
        }
      }
    },
    "/profile/payments": {
      "get": {
        "security": [
          {
            "bearerAuth": []
          }
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "profile"
        ],
        "operationId": "ListPayments",
        "responses": {
          "200": {
            "description": "Successful response",
            "schema": {
              "$ref": "#/definitions/ListPaymentsResponse"
            }
          }
        }
      }
    },
    "/register": {
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "auth"
        ],
        "operationId": "RegisterUser",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/RegisterRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successful response",
            "schema": {
              "$ref": "#/definitions/RegisterResponse"
            }
          },
          "400": {
            "description": "UserExists"
          }
        }
      }
    }
  },
  "definitions": {
    "ListMoviesResponse": {
      "type": "object",
      "properties": {
        "PageNum": {
          "type": "integer"
        },
        "PageSize": {
          "type": "integer"
        },
        "Payload": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Movie"
          }
        }
      }
    },
    "ListOrdersResponse": {
      "type": "object",
      "properties": {
        "PageNum": {
          "type": "integer"
        },
        "PageSize": {
          "type": "integer"
        },
        "Payload": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Order"
          }
        }
      }
    },
    "ListPaymentsResponse": {
      "type": "object",
      "properties": {
        "PageNum": {
          "type": "integer"
        },
        "PageSize": {
          "type": "integer"
        },
        "Payload": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Payment"
          }
        }
      }
    },
    "LoginRequest": {
      "type": "object",
      "properties": {
        "Email": {
          "type": "string"
        },
        "Password": {
          "type": "string"
        }
      }
    },
    "LoginResponse": {
      "type": "object",
      "properties": {
        "JWT": {
          "type": "string"
        }
      }
    },
    "Movie": {
      "type": "object",
      "properties": {
        "Category": {
          "type": "integer"
        },
        "ID": {
          "type": "integer",
          "format": "int32"
        },
        "ReleaseYear": {
          "type": "string"
        },
        "Title": {
          "type": "string"
        }
      }
    },
    "Order": {
      "type": "object",
      "properties": {
        "Amount": {
          "type": "number",
          "format": "float64"
        },
        "CreatedAt": {
          "type": "integer",
          "format": "int64"
        },
        "ID": {
          "type": "integer",
          "format": "int32"
        },
        "MovieID": {
          "type": "integer"
        }
      }
    },
    "Payment": {
      "type": "object",
      "properties": {
        "Amount": {
          "type": "number",
          "format": "float64"
        },
        "CreatedAt": {
          "type": "integer",
          "format": "int64"
        },
        "ID": {
          "type": "integer",
          "format": "int32"
        },
        "Status": {
          "type": "integer"
        },
        "TransactionID": {
          "type": "integer",
          "format": "int32"
        },
        "UpdatedAt": {
          "type": "integer",
          "format": "int64"
        }
      }
    },
    "Profile": {
      "type": "object",
      "properties": {
        "Age": {
          "type": "integer"
        },
        "Balance": {
          "type": "number",
          "format": "float64"
        },
        "DisplayName": {
          "type": "string"
        },
        "Email": {
          "type": "string"
        },
        "ID": {
          "type": "integer"
        },
        "Phone": {
          "type": "string"
        },
        "Role": {
          "type": "string"
        },
        "Status": {
          "type": "integer"
        }
      }
    },
    "RegisterRequest": {
      "type": "object",
      "properties": {
        "Age": {
          "type": "integer"
        },
        "DisplayName": {
          "type": "string"
        },
        "Email": {
          "type": "string"
        },
        "Password": {
          "type": "string"
        },
        "Phone": {
          "type": "string"
        }
      }
    },
    "RegisterResponse": {
      "type": "object",
      "properties": {
        "JWT": {
          "type": "string"
        }
      }
    },
    "RentMovieResponse": {
      "type": "object",
      "properties": {
        "Payload": {
          "type": "object",
          "properties": {
            "MovieID": {
              "type": "integer",
              "format": "int32"
            },
            "RentEndTime": {
              "type": "integer",
              "format": "int64"
            }
          }
        }
      }
    },
    "RentMovieResponsePayload": {
      "type": "object",
      "properties": {
        "MovieID": {
          "type": "integer",
          "format": "int32"
        },
        "RentEndTime": {
          "type": "integer",
          "format": "int64"
        }
      }
    }
  },
  "securityDefinitions": {
    "bearerAuth": {
      "type": "apiKey",
      "name": "Authorization",
      "in": "header"
    }
  },
  "x-components": {}
}`))
}
