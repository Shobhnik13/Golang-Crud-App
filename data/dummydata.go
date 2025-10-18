package data

import "course/models"

var Courses = []models.Course{
	{
		ID:          1,
		Title:       "Introduction to Go",
		Description: "Learn the basics of Go programming language.",
		Price:       49,
		Author: models.Author{
			ID:   1,
			Name: "John Doe",
		},
	},
	{
		ID:          2,
		Title:       "Introduction to Js",
		Description: "Learn the basics of Js programming language.",
		Price:       48,
		Author: models.Author{
			ID:   2,
			Name: "Jane Smith",
		},
	},
}
