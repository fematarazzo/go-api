package main

import (
	"fmt"
	"log"
	"net/http"

	"api/src/config"
	"api/src/router"
)

func main() {
	config.Load()

	fmt.Println(config.DBStringConnection)

	r := router.Generate()

	log.Printf("Listening to port %d", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
