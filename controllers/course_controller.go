package controllers

import (
	"course/data"
	"encoding/json"
	"net/http"
)

func GetCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(data.Courses)
}
