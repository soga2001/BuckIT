package router

import (
	"net/http"
	"server/api/resource/database"
	"server/api/resource/users"

	"github.com/danielgtaylor/huma/v2"
	"github.com/rs/cors"
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
	db, err := database.GetClient()
	if err != nil {
		return nil, err
	}

	// Create the user group
	userGroup := huma.NewGroup(grp, "/users")
	// Create a new user handler instance.
	userHandler := users.NewUserHandler(db)
	// Create routes for the user group
	huma.Register(
		userGroup,
		huma.Operation{
			OperationID: "get_loggedin_user",
			Method:      http.MethodGet,
			Path:        "/current_user",
			Summary:     "Get current user",
			Description: "Get users that is currently logged in",
		},
		userHandler.GetLoggedinUser,
	)
	huma.Register(
		userGroup,
		huma.Operation{
			OperationID: "all_users",
			Method:      http.MethodGet,
			Path:        "/",
			Summary:     "Get users",
			Description: "Get users",
		},
		userHandler.GetUsers,
	)
	huma.Register(
		userGroup,
		huma.Operation{
			OperationID: "user_by_username",
			Method:      http.MethodGet,
			Path:        "/user_by_username/:username",
			Summary:     "Get user by username",
			Description: "Get a user (user profile) by their username",
		},
		userHandler.GetUserByUsername,
	)
	huma.Register(
		userGroup,
		huma.Operation{
			OperationID: "user_by_id",
			Method:      http.MethodGet,
			Path:        "/user_by_id/:id",
			Summary:     "Get user by id",
			Description: "Get a user (user profile) by their id",
		},
		userHandler.GetUserByID,
	)

	huma.Register(
		userGroup,
		huma.Operation{
			OperationID: "user_login",
			Method:      http.MethodPost,
			Path:        "/login",
			Summary:     "User Login",
			Description: "Login a user",
		},
		userHandler.Login,
	)
	huma.Register(
		userGroup,
		huma.Operation{
			OperationID: "user_register",
			Method:      http.MethodPost,
			Path:        "/register",
			Summary:     "User Register",
			Description: "Register a user",
		},
		userHandler.Register,
	)
	huma.Register(
		userGroup,
		huma.Operation{
			OperationID: "user_logout",
			Method:      http.MethodPost,
			Path:        "/logout",
			Summary:     "User Logout",
			Description: "Logouts a logged in user",
		},
		userHandler.Logout,
	)

	handler := http.Handler(router)
	_ = cors.Default().Handler(handler)

	return router, nil
}
