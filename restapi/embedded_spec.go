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
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Mp3tag service. You can find out more at [http://mp3tags.org](http://mp3tags.org)",
    "title": "mp3tags",
    "version": "0.0.0"
  },
  "host": "mp3tags.org",
  "basePath": "/api/v1",
  "paths": {
    "/api/v1/tag": {
      "post": {
        "summary": "change tags",
        "parameters": [
          {
            "name": "tags",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "requests/tag.yaml#/ChangeTagsRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Succesfull",
            "schema": {
              "type": "file"
            }
          },
          "400": {
            "description": "Bad Request"
          },
          "500": {
            "description": "Internal error"
          }
        }
      }
    },
    "/api/v1/upload": {
      "post": {
        "summary": "upload file to server",
        "parameters": [
          {
            "type": "file",
            "name": "file",
            "in": "formData",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Succesful",
            "schema": {
              "$ref": "responses/upload.yaml#/UploadResponse"
            }
          },
          "400": {
            "description": "Bad Request"
          },
          "500": {
            "description": "Internal error"
          }
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Mp3tag service. You can find out more at [http://mp3tags.org](http://mp3tags.org)",
    "title": "mp3tags",
    "version": "0.0.0"
  },
  "host": "mp3tags.org",
  "basePath": "/api/v1",
  "paths": {
    "/api/v1/tag": {
      "post": {
        "summary": "change tags",
        "parameters": [
          {
            "name": "tags",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/changeTagsRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Succesfull",
            "schema": {
              "type": "file"
            }
          },
          "400": {
            "description": "Bad Request"
          },
          "500": {
            "description": "Internal error"
          }
        }
      }
    },
    "/api/v1/upload": {
      "post": {
        "summary": "upload file to server",
        "parameters": [
          {
            "type": "file",
            "name": "file",
            "in": "formData",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Succesful",
            "schema": {
              "$ref": "#/definitions/uploadResponse"
            }
          },
          "400": {
            "description": "Bad Request"
          },
          "500": {
            "description": "Internal error"
          }
        }
      }
    }
  },
  "definitions": {
    "changeTag": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "x-nullable": false
        },
        "value": {
          "type": "string",
          "x-nullable": "string"
        }
      },
      "x-go-name": "tag"
    },
    "changeTagsRequest": {
      "type": "object",
      "required": [
        "version",
        "tags"
      ],
      "properties": {
        "tags": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/changeTag"
          }
        },
        "version": {
          "type": "string",
          "x-nullable": false
        }
      }
    },
    "tagResponse": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "x-nullable": false
        },
        "value": {
          "type": "object",
          "x-nullable": false
        }
      },
      "x-go-name": "tag"
    },
    "uploadResponse": {
      "type": "object",
      "required": [
        "version",
        "tags"
      ],
      "properties": {
        "tags": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/tagResponse"
          }
        },
        "version": {
          "type": "string",
          "x-nullable": false
        }
      }
    }
  }
}`))
}
