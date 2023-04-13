package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nico-mayer/go-elements/config"
	"github.com/nico-mayer/go-elements/controllers"
	"github.com/nico-mayer/go-elements/db"
)

func main() {
	err := db.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	handleRequests()
}

func handleRequests() {
	http.HandleFunc("/elements", controllers.GetAllElements)
	http.HandleFunc("/element/atomicNumber/", controllers.ElementByAtomicNumber)

	fmt.Println("Server listening on PORT: " + config.PORT)
	log.Fatal(http.ListenAndServe(":"+config.PORT, nil))
}
