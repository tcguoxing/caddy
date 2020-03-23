package integration

import (
	"testing"

	"github.com/caddyserver/caddy/v2/caddytest"
)

func TestDefaultSNI(t *testing.T) {

	// arrange
	caddytest.InitServer(t, `{
    "apps": {
      "http": {
        "http_port": 9080,
        "https_port": 9443,
        "servers": {
          "srv0": {
            "listen": [
              ":9443"
            ],
            "routes": [
              {
                "handle": [
                  {
                    "handler": "subroute",
                    "routes": [
                      {
                        "handle": [
                          {
                            "body": "hello from a.caddy.localhost",
                            "handler": "static_response",
                            "status_code": 200
                          }
                        ],
                        "match": [
                          {
                            "path": [
                              "/version"
                            ]
                          }
                        ]
                      }
                    ]
                  }
                ],
                "match": [
                  {
                    "host": [
                      "127.0.0.1"
                    ]
                  }
                ],
                "terminal": true
              }
            ],
            "tls_connection_policies": [
              {
                "certificate_selection": {
                  "policy": "custom",
                  "tag": "cert0"
                },
                "match": {
                  "sni": [
                    "127.0.0.1"
                  ]
                }
              },
              {
                "default_sni": "*.caddy.localhost"
              }
            ]
          }
        }
      },
      "tls": {
        "certificates": {
          "load_files": [
            {
              "certificate": "/caddy.localhost.crt",
              "key": "/caddy.localhost.key",
              "tags": [
                "cert0"
              ]
            }
          ]
        }
      },
      "pki": {
        "certificate_authorities" : {
          "local" : {
            "install_trust": false
          }
        }
      }
    }
  }
  `, "json")

	// act and assert
	// makes a request with no sni
	caddytest.AssertGetResponse(t, "https://127.0.0.1:9443/version", 200, "hello from a")
}

func TestDefaultSNIWithNamedHostAndExplicitIP(t *testing.T) {

	// arrange
	caddytest.InitServer(t, ` 
  {
    "apps": {
      "http": {
        "http_port": 9080,
        "https_port": 9443,
        "servers": {
          "srv0": {
            "listen": [
              ":9443"
            ],
            "routes": [
              {
                "handle": [
                  {
                    "handler": "subroute",
                    "routes": [
                      {
                        "handle": [
                          {
                            "body": "hello from a",
                            "handler": "static_response",
                            "status_code": 200
                          }
                        ],
                        "match": [
                          {
                            "path": [
                              "/version"
                            ]
                          }
                        ]
                      }
                    ]
                  }
                ],
                "match": [
                  {
                    "host": [
                      "a.caddy.localhost",
                      "127.0.0.1"
                    ]
                  }
                ],
                "terminal": true
              }
            ],
            "tls_connection_policies": [
              {
                "certificate_selection": {
                  "policy": "custom",
                  "tag": "cert0"
                },
                "default_sni": "a.caddy.localhost",
                "match": {
                  "sni": [
                    "a.caddy.localhost",
                    "127.0.0.1",
                    ""
                  ]
                }
              },
              {
                "default_sni": "a.caddy.localhost"
              }
            ]
          }
        }
      },
      "tls": {
        "certificates": {
          "load_files": [
            {
              "certificate": "/a.caddy.localhost.crt",
              "key": "/a.caddy.localhost.key",
              "tags": [
                "cert0"
              ]
            }
          ]
        }
      },
      "pki": {
        "certificate_authorities" : {
          "local" : {
            "install_trust": false
          }
        }
      }
    }
  }
  `, "json")

	// act and assert
	// makes a request with no sni
	caddytest.AssertGetResponse(t, "https://127.0.0.1:9443/version", 200, "hello from a")
}

func TestDefaultSNIWithPortMappingOnly(t *testing.T) {

	// arrange
	caddytest.InitServer(t, ` 
  {
    "apps": {
      "http": {
        "http_port": 9080,
        "https_port": 9443,
        "servers": {
          "srv0": {
            "listen": [
              ":9443"
            ],
            "routes": [
              {
                "handle": [
                  {
                    "body": "hello from a.caddy.localhost",
                    "handler": "static_response",
                    "status_code": 200
                  }
                ],
                "match": [
                  {
                    "path": [
                      "/version"
                    ]
                  }
                ]
              }
            ],
            "tls_connection_policies": [
              {
                "certificate_selection": {
                  "policy": "custom",
                  "tag": "cert0"
                },
                "default_sni": "a.caddy.localhost"
              }
            ]
          }
        }
      },
      "tls": {
        "certificates": {
          "load_files": [
            {
              "certificate": "/a.caddy.localhost.crt",
              "key": "/a.caddy.localhost.key",
              "tags": [
                "cert0"
              ]
            }
          ]
        }
      },
      "pki": {
        "certificate_authorities" : {
          "local" : {
            "install_trust": false
          }
        }
      }
    }
  }
  `, "json")

	// act and assert
	// makes a request with no sni
	caddytest.AssertGetResponse(t, "https://127.0.0.1:9443/version", 200, "hello from a")
}