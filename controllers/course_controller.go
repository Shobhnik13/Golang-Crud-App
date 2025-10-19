package controllers

import (
	"course/data"
	"course/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for _, course := range data.Courses {
		if course.ID == id {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	http.Error(w, "Course not found", http.StatusNotFound)
}

func UpdateCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	params := mux.Vars(r)
	idString := params["id"]

	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "Invalid course ID", http.StatusBadRequest)
		return
	}

	var updatedCourse models.Course
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&updatedCourse)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	for i, course := range data.Courses {
		if course.ID == id {
			updatedCourse.ID = id
			data.Courses[i] = updatedCourse
			json.NewEncoder(w).Encode(updatedCourse)
			return
		}
	}
	http.Error(w, "Course not found", http.StatusNotFound)
}

func DeleteCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	params := mux.Vars(r)
	idString := params["id"]

	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "Invalid course ID", http.StatusBadRequest)
		return
	}

	for i, course := range data.Courses {
		if course.ID == id {
			data.Courses = append(data.Courses[:i], data.Courses[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "Course not found", http.StatusNotFound)
}
