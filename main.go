package main

import (
	"course/routes"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// creating a new router
	r := mux.NewRouter()

	// register course routes
	routes.RegisterCourseRoutes(r)

	// start the server
	fmt.Println("Server running on port 5000")
	http.ListenAndServe(":5000", r)
}
