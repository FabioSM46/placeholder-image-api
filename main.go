package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/FabioSM46/placeholder-image-api/api"
	"github.com/joho/godotenv"
)



func main() {
    err := godotenv.Load()
    if err != nil {
        log.Printf("No .env file found or unable to load: %v", err)
    }

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
        log.Println("PORT not set in .env file, using default port 8080")
    }

    http.HandleFunc("/", api.ImageHandler)
    http.HandleFunc("/favicon.ico", api.FaviconHandler)
    fmt.Printf("Server starting at :%s\n", port)
    log.Fatal(http.ListenAndServe(":"+port, nil))
}