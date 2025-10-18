package controllers

import (
	"course/data"
	"course/models"
	"encoding/json"
	"net/http"
)

func GetCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(data.Courses)
}

func CreateCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var course models.Course

	// body ko decode karne ke lie
	decoder := json.NewDecoder(r.Body)

	// decoding the body into course struct
	err := decoder.Decode(&course)

	// error handling
	if err != nil {
		http.Error(w, "course is required", http.StatusBadRequest)
		return
	}

	// validation
	if course.Title == "" || course.Description == "" {
		http.Error(w, "course title and desc is required", http.StatusBadRequest)
		return
	}

	// adding unique IDs
	courseID := len(data.Courses) + 1
	authorID := len(data.Courses) + 1

	// assigning IDs
	course.Author.ID = authorID
	course.ID = courseID
	data.Courses = append(data.Courses, course)

	// sending response
	err = json.NewEncoder(w).Encode(course)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetCourseByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	
}
