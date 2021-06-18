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
    "title": "GoChat API",
    "version": "1.0.0"
  },
  "basePath": "/api/v1",
  "paths": {
    "/account": {
      "get": {
        "security": [
          {
            "jwt": []
          }
        ],
        "description": "Endpoint to get user account information.",
        "tags": [
          "stable"
        ],
        "responses": {
          "200": {
            "description": "Requester's user account returned.",
            "schema": {
              "$ref": "#/definitions/user"
            }
          },
          "401": {
            "description": "Unauthorized",
            "schema": {
              "$ref": "#/responses/Unauthorized"
            }
          },
          "500": {
            "description": "Unknown Error",
            "schema": {
              "$ref": "#/responses/Unknown"
            }
          }
        }
      },
      "put": {
        "security": [
          {
            "jwt": []
          }
        ],
        "description": "Endpoint for a user to update their profile.",
        "tags": [
          "stable"
        ],
        "parameters": [
          {
            "name": "user",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/user"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Account was updated.",
            "schema": {
              "$ref": "#/definitions/user"
            }
          },
          "401": {
            "description": "Unauthorized",
            "schema": {
              "$ref": "#/responses/Unauthorized"
            }
          },
          "500": {
            "description": "Unknown Error",
            "schema": {
              "$ref": "#/responses/Unknown"
            }
          }
        }
      },
      "post": {
        "description": "Endpoint to create a new account.",
        "tags": [
          "stable"
        ],
        "parameters": [
          {
            "name": "user",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/user"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Account was created.",
            "schema": {
              "$ref": "#/definitions/user"
            }
          },
          "500": {
            "description": "Unknown Error",
            "schema": {
              "$ref": "#/responses/Unknown"
            }
          }
        }
      }
    },
    "/auth/login": {
      "post": {
        "description": "Logs a user in and returns a JWT for secured API calls.",
        "tags": [
          "stable"
        ],
        "parameters": [
          {
            "name": "login",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/login_request"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "User was successfully logged in.",
            "schema": {
              "$ref": "#/definitions/login_response"
            }
          },
          "401": {
            "description": "Unauthorized",
            "schema": {
              "$ref": "#/responses/Unauthorized"
            }
          },
          "500": {
            "description": "Unknown Error",
            "schema": {
              "$ref": "#/responses/Unknown"
            }
          }
        }
      }
    },
    "/conversation": {
      "put": {
        "security": [
          {
            "jwt": []
          }
        ],
        "description": "Updates an existing conversation.",
        "tags": [
          "stable"
        ],
        "parameters": [
          {
            "name": "conversation",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/conversation"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Conversation was updated.",
            "schema": {
              "$ref": "#/definitions/conversation"
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/responses/BadRequest"
            }
          },
          "401": {
            "description": "Unauthorized",
            "schema": {
              "$ref": "#/responses/Unauthorized"
            }
          },
          "500": {
            "description": "Unknown Error",
            "schema": {
              "$ref": "#/responses/Unknown"
            }
          }
        }
      },
      "post": {
        "security": [
          {
            "jwt": []
          }
        ],
        "description": "Creates a new conversation.",
        "tags": [
          "stable"
        ],
        "parameters": [
          {
            "name": "conversation",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/conversation"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Conversation was created.",
            "schema": {
              "$ref": "#/definitions/conversation"
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/responses/BadRequest"
            }
          },
          "401": {
            "description": "Unauthorized",
            "schema": {
              "$ref": "#/responses/Unauthorized"
            }
          },
          "500": {
            "description": "Unknown Error",
            "schema": {
              "$ref": "#/responses/Unknown"
            }
          }
        }
      }
    },
    "/conversation/{id}": {
      "get": {
        "security": [
          {
            "jwt": []
          }
        ],
        "description": "Fetches an existing conversation.",
        "tags": [
          "stable"
        ],
        "parameters": [
          {
            "type": "string",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Conversation was retrieved.",
            "schema": {
              "$ref": "#/definitions/conversation"
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/responses/BadRequest"
            }
          },
          "401": {
            "description": "Unauthorized",
            "schema": {
              "$ref": "#/responses/Unauthorized"
            }
          },
          "404": {
            "description": "Not Found",
            "schema": {
              "$ref": "#/responses/NotFound"
            }
          },
          "500": {
            "description": "Unknown Error",
            "schema": {
              "$ref": "#/responses/Unknown"
            }
          }
        }
      }
    },
    "/message": {
      "put": {
        "security": [
          {
            "jwt": []
          }
        ],
        "description": "Updates an existing message.",
        "tags": [
          "stable"
        ],
        "parameters": [
          {
            "name": "message",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/message"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Message was updated.",
            "schema": {
              "$ref": "#/definitions/message"
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/responses/BadRequest"
            }
          },
          "401": {
            "description": "Unauthorized",
            "schema": {
              "$ref": "#/responses/Unauthorized"
            }
          },
          "500": {
            "description": "Unknown Error",
            "schema": {
              "$ref": "#/responses/Unknown"
            }
          }
        }
      },
      "post": {
        "security": [
          {
            "jwt": []
          }
        ],
        "description": "Creates a new message.",
        "tags": [
          "stable"
        ],
        "parameters": [
          {
            "name": "message",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/message"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Message was created.",
            "schema": {
              "$ref": "#/definitions/message"
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/responses/BadRequest"
            }
          },
          "401": {
            "description": "Unauthorized",
            "schema": {
              "$ref": "#/responses/Unauthorized"
            }
          },
          "500": {
            "description": "Unknown Error",
            "schema": {
              "$ref": "#/responses/Unknown"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "conversation": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "format": "uuid"
        },
        "is_public": {
          "type": "boolean"
        },
        "label": {
          "type": "string",
          "example": "#all-social"
        },
        "last_message_on": {
          "type": "string",
          "format": "date-time"
        },
        "messages": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/message"
          }
        },
        "participants": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/user"
          }
        }
      }
    },
    "error": {
      "type": "object",
      "properties": {
        "message": {
          "type": "string"
        }
      }
    },
    "login_request": {
      "type": "object",
      "properties": {
        "email": {
          "type": "string",
          "format": "email"
        },
        "password": {
          "type": "string"
        }
      }
    },
    "login_response": {
      "type": "object",
      "properties": {
        "auth_token": {
          "type": "string"
        },
        "user": {
          "type": "object",
          "$ref": "#/definitions/user"
        }
      }
    },
    "message": {
      "type": "object",
      "properties": {
        "author_id": {
          "type": "string",
          "format": "uuid"
        },
        "conversation_id": {
          "type": "string",
          "format": "uuid"
        },
        "created_at": {
          "type": "string",
          "format": "date-time"
        },
        "id": {
          "type": "string",
          "format": "uuid"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "principal": {
      "type": "string"
    },
    "user": {
      "type": "object",
      "properties": {
        "email": {
          "type": "string",
          "format": "email"
        },
        "first_name": {
          "type": "string",
          "example": "John"
        },
        "id": {
          "type": "string",
          "format": "uuid"
        },
        "last_name": {
          "type": "string",
          "example": "Smith"
        },
        "password": {
          "type": "string"
        },
        "salt": {
          "type": "string"
        }
      }
    }
  },
  "responses": {
    "BadRequest": {
      "description": "Bad Request",
      "schema": {
        "$ref": "#/definitions/error"
      }
    },
    "NotFound": {
      "description": "Not Found",
      "schema": {
        "$ref": "#/definitions/error"
      }
    },
    "Unauthorized": {
      "description": "Unauthorized",
      "schema": {
        "$ref": "#/definitions/error"
      }
    },
    "Unknown": {
      "description": "Unknown",
      "schema": {
        "$ref": "#/definitions/error"
      }
    }
  },
  "securityDefinitions": {
    "jwt": {
      "type": "apiKey",
      "name": "Authorization",
      "in": "header"
    }
  },
  "security": [
    {
      "jwt": []
    }
  ],
  "tags": [
    {
      "description": "Ready for integration",
      "name": "stable"
    }
  ]
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
    "title": "GoChat API",
    "version": "1.0.0"
  },
  "basePath": "/api/v1",
  "paths": {
    "/account": {
      "get": {
        "security": [
          {
            "jwt": []
          }
        ],
        "description": "Endpoint to get user account information.",
        "tags": [
          "stable"
        ],
        "responses": {
          "200": {
            "description": "Requester's user account returned.",
            "schema": {
              "$ref": "#/definitions/user"
            }
          },
          "401": {
            "description": "Unauthorized",
            "schema": {
              "description": "Unauthorized",
              "schema": {
                "$ref": "#/definitions/error"
              }
            }
          },
          "500": {
            "description": "Unknown Error",
            "schema": {
              "description": "Unknown",
              "schema": {
                "$ref": "#/definitions/error"
              }
            }
          }
        }
      },
      "put": {
        "security": [
          {
            "jwt": []
          }
        ],
        "description": "Endpoint for a user to update their profile.",
        "tags": [
          "stable"
        ],
        "parameters": [
          {
            "name": "user",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/user"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Account was updated.",
            "schema": {
              "$ref": "#/definitions/user"
            }
          },
          "401": {
            "description": "Unauthorized",
            "schema": {
              "description": "Unauthorized",
              "schema": {
                "$ref": "#/definitions/error"
              }
            }
          },
          "500": {
            "description": "Unknown Error",
            "schema": {
              "description": "Unknown",
              "schema": {
                "$ref": "#/definitions/error"
              }
            }
          }
        }
      },
      "post": {
        "description": "Endpoint to create a new account.",
        "tags": [
          "stable"
        ],
        "parameters": [
          {
            "name": "user",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/user"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Account was created.",
            "schema": {
              "$ref": "#/definitions/user"
            }
          },
          "500": {
            "description": "Unknown Error",
            "schema": {
              "description": "Unknown",
              "schema": {
                "$ref": "#/definitions/error"
              }
            }
          }
        }
      }
    },
    "/auth/login": {
      "post": {
        "description": "Logs a user in and returns a JWT for secured API calls.",
        "tags": [
          "stable"
        ],
        "parameters": [
          {
            "name": "login",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/login_request"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "User was successfully logged in.",
            "schema": {
              "$ref": "#/definitions/login_response"
            }
          },
          "401": {
            "description": "Unauthorized",
            "schema": {
              "description": "Unauthorized",
              "schema": {
                "$ref": "#/definitions/error"
              }
            }
          },
          "500": {
            "description": "Unknown Error",
            "schema": {
              "description": "Unknown",
              "schema": {
                "$ref": "#/definitions/error"
              }
            }
          }
        }
      }
    },
    "/conversation": {
      "put": {
        "security": [
          {
            "jwt": []
          }
        ],
        "description": "Updates an existing conversation.",
        "tags": [
          "stable"
        ],
        "parameters": [
          {
            "name": "conversation",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/conversation"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Conversation was updated.",
            "schema": {
              "$ref": "#/definitions/conversation"
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "description": "Bad Request",
              "schema": {
                "$ref": "#/definitions/error"
              }
            }
          },
          "401": {
            "description": "Unauthorized",
            "schema": {
              "description": "Unauthorized",
              "schema": {
                "$ref": "#/definitions/error"
              }
            }
          },
          "500": {
            "description": "Unknown Error",
            "schema": {
              "description": "Unknown",
              "schema": {
                "$ref": "#/definitions/error"
              }
            }
          }
        }
      },
      "post": {
        "security": [
          {
            "jwt": []
          }
        ],
        "description": "Creates a new conversation.",
        "tags": [
          "stable"
        ],
        "parameters": [
          {
            "name": "conversation",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/conversation"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Conversation was created.",
            "schema": {
              "$ref": "#/definitions/conversation"
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "description": "Bad Request",
              "schema": {
                "$ref": "#/definitions/error"
              }
            }
          },
          "401": {
            "description": "Unauthorized",
            "schema": {
              "description": "Unauthorized",
              "schema": {
                "$ref": "#/definitions/error"
              }
            }
          },
          "500": {
            "description": "Unknown Error",
            "schema": {
              "description": "Unknown",
              "schema": {
                "$ref": "#/definitions/error"
              }
            }
          }
        }
      }
    },
    "/conversation/{id}": {
      "get": {
        "security": [
          {
            "jwt": []
          }
        ],
        "description": "Fetches an existing conversation.",
        "tags": [
          "stable"
        ],
        "parameters": [
          {
            "type": "string",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Conversation was retrieved.",
            "schema": {
              "$ref": "#/definitions/conversation"
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "description": "Bad Request",
              "schema": {
                "$ref": "#/definitions/error"
              }
            }
          },
          "401": {
            "description": "Unauthorized",
            "schema": {
              "description": "Unauthorized",
              "schema": {
                "$ref": "#/definitions/error"
              }
            }
          },
          "404": {
            "description": "Not Found",
            "schema": {
              "description": "Not Found",
              "schema": {
                "$ref": "#/definitions/error"
              }
            }
          },
          "500": {
            "description": "Unknown Error",
            "schema": {
              "description": "Unknown",
              "schema": {
                "$ref": "#/definitions/error"
              }
            }
          }
        }
      }
    },
    "/message": {
      "put": {
        "security": [
          {
            "jwt": []
          }
        ],
        "description": "Updates an existing message.",
        "tags": [
          "stable"
        ],
        "parameters": [
          {
            "name": "message",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/message"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Message was updated.",
            "schema": {
              "$ref": "#/definitions/message"
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "description": "Bad Request",
              "schema": {
                "$ref": "#/definitions/error"
              }
            }
          },
          "401": {
            "description": "Unauthorized",
            "schema": {
              "description": "Unauthorized",
              "schema": {
                "$ref": "#/definitions/error"
              }
            }
          },
          "500": {
            "description": "Unknown Error",
            "schema": {
              "description": "Unknown",
              "schema": {
                "$ref": "#/definitions/error"
              }
            }
          }
        }
      },
      "post": {
        "security": [
          {
            "jwt": []
          }
        ],
        "description": "Creates a new message.",
        "tags": [
          "stable"
        ],
        "parameters": [
          {
            "name": "message",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/message"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Message was created.",
            "schema": {
              "$ref": "#/definitions/message"
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "description": "Bad Request",
              "schema": {
                "$ref": "#/definitions/error"
              }
            }
          },
          "401": {
            "description": "Unauthorized",
            "schema": {
              "description": "Unauthorized",
              "schema": {
                "$ref": "#/definitions/error"
              }
            }
          },
          "500": {
            "description": "Unknown Error",
            "schema": {
              "description": "Unknown",
              "schema": {
                "$ref": "#/definitions/error"
              }
            }
          }
        }
      }
    }
  },
  "definitions": {
    "conversation": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "format": "uuid"
        },
        "is_public": {
          "type": "boolean"
        },
        "label": {
          "type": "string",
          "example": "#all-social"
        },
        "last_message_on": {
          "type": "string",
          "format": "date-time"
        },
        "messages": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/message"
          }
        },
        "participants": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/user"
          }
        }
      }
    },
    "error": {
      "type": "object",
      "properties": {
        "message": {
          "type": "string"
        }
      }
    },
    "login_request": {
      "type": "object",
      "properties": {
        "email": {
          "type": "string",
          "format": "email"
        },
        "password": {
          "type": "string"
        }
      }
    },
    "login_response": {
      "type": "object",
      "properties": {
        "auth_token": {
          "type": "string"
        },
        "user": {
          "type": "object",
          "$ref": "#/definitions/user"
        }
      }
    },
    "message": {
      "type": "object",
      "properties": {
        "author_id": {
          "type": "string",
          "format": "uuid"
        },
        "conversation_id": {
          "type": "string",
          "format": "uuid"
        },
        "created_at": {
          "type": "string",
          "format": "date-time"
        },
        "id": {
          "type": "string",
          "format": "uuid"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "principal": {
      "type": "string"
    },
    "user": {
      "type": "object",
      "properties": {
        "email": {
          "type": "string",
          "format": "email"
        },
        "first_name": {
          "type": "string",
          "example": "John"
        },
        "id": {
          "type": "string",
          "format": "uuid"
        },
        "last_name": {
          "type": "string",
          "example": "Smith"
        },
        "password": {
          "type": "string"
        },
        "salt": {
          "type": "string"
        }
      }
    }
  },
  "responses": {
    "BadRequest": {
      "description": "Bad Request",
      "schema": {
        "$ref": "#/definitions/error"
      }
    },
    "NotFound": {
      "description": "Not Found",
      "schema": {
        "$ref": "#/definitions/error"
      }
    },
    "Unauthorized": {
      "description": "Unauthorized",
      "schema": {
        "$ref": "#/definitions/error"
      }
    },
    "Unknown": {
      "description": "Unknown",
      "schema": {
        "$ref": "#/definitions/error"
      }
    }
  },
  "securityDefinitions": {
    "jwt": {
      "type": "apiKey",
      "name": "Authorization",
      "in": "header"
    }
  },
  "security": [
    {
      "jwt": []
    }
  ],
  "tags": [
    {
      "description": "Ready for integration",
      "name": "stable"
    }
  ]
}`))
}
