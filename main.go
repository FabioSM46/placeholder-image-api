package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/FabioSM46/placeholder-image-api/api"
)



func main() {
	http.HandleFunc("/", api.ImageHandler)
	fmt.Println("Server starting at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
