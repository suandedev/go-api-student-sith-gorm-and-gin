package router

import (
	"new2/middleware"

	"github.com/gin-gonic/gin"
)

func Router() {
	// gin
	router := gin.Default()
	// get all student
	router.GET("/student", middleware.GetAllStudent)
	// get student by id
	router.GET("/student/:id", middleware.GetStudentById)
	// create student
	router.POST("/student", middleware.CreateStudent)
	// update student
	router.PUT("/student/:id", middleware.UpdateStudent)
	// delete student
	router.DELETE("/student/:id", middleware.DeleteStudent)
	router.Run()
}
