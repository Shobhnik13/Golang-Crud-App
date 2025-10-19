package routes

import (
	"course/controllers"

	"github.com/gorilla/mux"
)

func RegisterCourseRoutes(router *mux.Router) {
	router.HandleFunc("/courses", controllers.GetCourses).Methods("GET")
	router.HandleFunc("/courses", controllers.CreateCourse).Methods("POST")
	router.HandleFunc("/courses/{id}", controllers.GetCourseByID).Methods("GET")
	router.HandleFunc("/courses/{id}", controllers.UpdateCourse).Methods("PUT")
	router.HandleFunc("/delete/courses/{id}", controllers.DeleteCourse).Methods("DELETE")
}
