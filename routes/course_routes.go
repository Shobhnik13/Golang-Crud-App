package routes

import (
	"course/controllers"

	"github.com/gorilla/mux"
)

func RegisterCourseRoutes(router *mux.Router) {
	router.HandleFunc("/courses", controllers.GetCourses).Methods("GET")
}
