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
        "contact": {},
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/discord/music/enqueue": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Play the song from YouTube by name or url",
                "parameters": [
                    {
                        "description": "Song name or url",
                        "name": "query",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/rest.songQuery"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "The song that was added to the queue",
                        "schema": {
                            "$ref": "#/definitions/pkg.SongRequest"
                        }
                    },
                    "400": {
                        "description": "Incorrect input",
                        "schema": {
                            "$ref": "#/definitions/rest.Response"
                        }
                    }
                }
            }
        },
        "/discord/music/loopstatus": {
            "get": {
                "produces": [
                    "text/plain"
                ],
                "summary": "Is loop mode enabled",
                "responses": {
                    "200": {
                        "description": "Returns true or false as string",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/discord/music/now": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Song that is playing now",
                "responses": {
                    "200": {
                        "description": "The song that is playing right now",
                        "schema": {
                            "$ref": "#/definitions/pkg.SongRequest"
                        }
                    }
                }
            }
        },
        "/discord/music/setloop": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Set loop mode",
                "parameters": [
                    {
                        "description": "Song name or url",
                        "name": "query",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/rest.loopQuery"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Incorrect input",
                        "schema": {
                            "$ref": "#/definitions/rest.Response"
                        }
                    }
                }
            }
        },
        "/discord/music/skip": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Skip the current song and play next from the queue",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/discord/music/stats": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Stats of player on current song",
                "responses": {
                    "200": {
                        "description": "The song that is playing right now",
                        "schema": {
                            "$ref": "#/definitions/audio.SessionStats"
                        }
                    }
                }
            }
        },
        "/discord/music/stop": {
            "get": {
                "produces": [
                    "text/plain"
                ],
                "summary": "Skip the current song and play next from the queue",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "audio.SessionStats": {
            "type": "object",
            "properties": {
                "duration": {
                    "description": "seconds",
                    "type": "number"
                },
                "position": {
                    "description": "seconds",
                    "type": "number"
                }
            }
        },
        "pkg.Metadata": {
            "type": "object",
            "properties": {
                "artists": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/pkg.MetadataArtist"
                    }
                },
                "artwork_url": {
                    "type": "string"
                },
                "display_url": {
                    "type": "string"
                },
                "duration": {
                    "type": "number"
                },
                "thumbnail_url": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "pkg.MetadataArtist": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "pkg.SongRequest": {
            "type": "object",
            "properties": {
                "metadata": {
                    "$ref": "#/definitions/pkg.Metadata"
                },
                "service_color": {
                    "description": "Color of service used for this queue entry",
                    "type": "integer"
                },
                "service_name": {
                    "description": "Name of service used for this queue entry",
                    "type": "string"
                }
            }
        },
        "rest.Response": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "rest.loopQuery": {
            "type": "object",
            "required": [
                "enable"
            ],
            "properties": {
                "enable": {
                    "type": "boolean"
                }
            }
        },
        "rest.songQuery": {
            "type": "object",
            "required": [
                "song"
            ],
            "properties": {
                "song": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:9091",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "HalvaBot for Discord",
	Description:      "A music discord bot.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}