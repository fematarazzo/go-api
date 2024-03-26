package main

import (
	"log"
	"net/http"

	"api/src/router"
)

func main() {
	r := router.Generate()

	log.Println("Listening to port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
