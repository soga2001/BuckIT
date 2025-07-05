package main

import (
	"fmt"
	"net/http"
	"server/api/router"

	_ "github.com/danielgtaylor/huma/v2/formats/cbor"
)

func main() {
	// Create a new router & API.
	r, err := router.CreateRouter()
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting server on http://localhost:8080")
	http.ListenAndServe("localhost:8080", r)
}
