package router

import (
	"net/http"
	"server/api/resource/users"

	"github.com/danielgtaylor/huma/v2"
	"github.com/uptrace/bunrouter"

	"github.com/danielgtaylor/huma/v2/adapters/humabunrouter"
	_ "github.com/danielgtaylor/huma/v2/formats/cbor"
)

// CreateRouter initializes the API router and sets up the routes.
func CreateRouter() (*bunrouter.Router, error) {
	// Create a new router & API.
	router := bunrouter.New()
	humaConfig := huma.DefaultConfig("BuckIT API", "1.0.0")
	humaConfig.DocsPath = ""
	humaConfig.Formats = map[string]huma.Format{
		"application/json": huma.DefaultJSONFormat,
		"json":             huma.DefaultJSONFormat,
	}
	api := humabunrouter.New(router, humaConfig)

	router.GET("/api/v1/-/docs", func(w http.ResponseWriter, req bunrouter.Request) error {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(`<!doctype html>
	<html>
	<head>
		<title>API Reference</title>
		<meta charset="utf-8" />
		<meta
		name="viewport"
		content="width=device-width, initial-scale=1" />
	</head>
	<body>
		<script
		id="api-reference"
		data-url="/openapi.json"></script>
		<script src="https://cdn.jsdelivr.net/npm/@scalar/api-reference"></script>
	</body>
	</html>`))
		return nil
	})

	// Create the default API group.
	grp := huma.NewGroup(api, "/api/v1")

	// Create the user group
	userGroup := huma.NewGroup(grp, "/users")
	// Create a new user handler instance.
	userHandler := users.NewUserHandler()
	// Create routes for the user group.
	huma.Get(userGroup, "/", userHandler.GetUsers)

	return router, nil
}
