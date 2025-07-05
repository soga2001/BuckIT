package main

import (
	"fmt"
	"log"
	"net/http"
	"server/api/router"

	_ "github.com/danielgtaylor/huma/v2/formats/cbor"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Create a new router & API.
	r, err := router.CreateRouter()
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting server on http://localhost:8080")
	http.ListenAndServe("localhost:8080", r)
}
