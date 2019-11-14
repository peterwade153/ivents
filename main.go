package main

import (
	"log"
	"net/http"
	"os"

	"github.com/peterwade153/ivents/api/routes"
)

func main() {

	port := os.Getenv("PORT")

	http.Handle("/", routes.Handlers())
	log.Printf("\nServer starting on port '%s'", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
