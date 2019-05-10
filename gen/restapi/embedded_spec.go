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
  "consumes": [
    "application/json",
    "application/xml"
  ],
  "produces": [
    "application/json",
    "application/xml"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Generates update feeds for different types of media",
    "title": "Feed Generator"
  },
  "paths": {
    "/api/feed/manga": {
      "post": {
        "description": "Creates a URL containing the current feed for the requested manga titles",
        "schemes": [
          "http"
        ],
        "summary": "Create feed from manga titles",
        "operationId": "feedgen#Manga",
        "parameters": [
          {
            "name": "MangaRequestBody",
            "in": "body",
            "required": true,
            "schema": {
              "required": [
                "titles"
              ],
              "$ref": "#/definitions/FeedgenMangaRequestBody"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK response.",
            "schema": {
              "type": "string"
            }
          },
          "404": {
            "description": "Not Found response."
          },
          "500": {
            "description": "Internal Server Error response."
          },
          "502": {
            "description": "Bad Gateway response."
          }
        }
      }
    },
    "/api/feed/manga/{hash}": {
      "get": {
        "description": "Returns an RSS Feed of the manga titles. Can return atom and json feeds as well.",
        "produces": [
          "application/xml",
          "application/json"
        ],
        "schemes": [
          "http"
        ],
        "summary": "Get rss/atom/json feed of manga updates",
        "operationId": "feedgen#viewManga",
        "parameters": [
          {
            "type": "string",
            "description": "Identifier of previously created manga feed",
            "name": "hash",
            "in": "path",
            "required": true
          },
          {
            "enum": [
              "rss",
              "atom",
              "json"
            ],
            "type": "string",
            "default": "rss",
            "description": "RSS, Atom, or JSON Feed",
            "name": "feedType",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "OK response.",
            "schema": {
              "type": "string"
            }
          },
          "404": {
            "description": "Not Found response."
          },
          "500": {
            "description": "Internal Server Error response."
          },
          "502": {
            "description": "Bad Gateway response."
          }
        }
      }
    }
  },
  "definitions": {
    "FeedgenMangaRequestBody": {
      "type": "object",
      "title": "FeedgenMangaRequestBody",
      "required": [
        "titles"
      ],
      "properties": {
        "titles": {
          "description": "List of manga titles to subscribe to",
          "type": "array",
          "maxItems": 2048,
          "minItems": 1,
          "items": {
            "type": "string",
            "example": "Quas impedit ratione esse."
          },
          "example": [
            "Ut labore quis atque nobis debitis."
          ]
        }
      },
      "example": {
        "titles": [
          "Culpa atque et."
        ]
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json",
    "application/xml"
  ],
  "produces": [
    "application/json",
    "application/xml"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Generates update feeds for different types of media",
    "title": "Feed Generator"
  },
  "paths": {
    "/api/feed/manga": {
      "post": {
        "description": "Creates a URL containing the current feed for the requested manga titles",
        "schemes": [
          "http"
        ],
        "summary": "Create feed from manga titles",
        "operationId": "feedgen#Manga",
        "parameters": [
          {
            "name": "MangaRequestBody",
            "in": "body",
            "required": true,
            "schema": {
              "required": [
                "titles"
              ],
              "$ref": "#/definitions/FeedgenMangaRequestBody"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK response.",
            "schema": {
              "type": "string"
            }
          },
          "404": {
            "description": "Not Found response."
          },
          "500": {
            "description": "Internal Server Error response."
          },
          "502": {
            "description": "Bad Gateway response."
          }
        }
      }
    },
    "/api/feed/manga/{hash}": {
      "get": {
        "description": "Returns an RSS Feed of the manga titles. Can return atom and json feeds as well.",
        "produces": [
          "application/xml",
          "application/json"
        ],
        "schemes": [
          "http"
        ],
        "summary": "Get rss/atom/json feed of manga updates",
        "operationId": "feedgen#viewManga",
        "parameters": [
          {
            "type": "string",
            "description": "Identifier of previously created manga feed",
            "name": "hash",
            "in": "path",
            "required": true
          },
          {
            "enum": [
              "rss",
              "atom",
              "json"
            ],
            "type": "string",
            "default": "rss",
            "description": "RSS, Atom, or JSON Feed",
            "name": "feedType",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "OK response.",
            "schema": {
              "type": "string"
            }
          },
          "404": {
            "description": "Not Found response."
          },
          "500": {
            "description": "Internal Server Error response."
          },
          "502": {
            "description": "Bad Gateway response."
          }
        }
      }
    }
  },
  "definitions": {
    "FeedgenMangaRequestBody": {
      "type": "object",
      "title": "FeedgenMangaRequestBody",
      "required": [
        "titles"
      ],
      "properties": {
        "titles": {
          "description": "List of manga titles to subscribe to",
          "type": "array",
          "maxItems": 2048,
          "minItems": 1,
          "items": {
            "type": "string",
            "example": "Quas impedit ratione esse."
          },
          "example": [
            "Ut labore quis atque nobis debitis."
          ]
        }
      },
      "example": {
        "titles": [
          "Culpa atque et."
        ]
      }
    }
  }
}`))
}