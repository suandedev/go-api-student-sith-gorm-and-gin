package middleware

import (
	"new2/fun"
	"new2/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

type student struct {
	ID   int    `gorm:"primaryKey autoIncrement"`
	Name string `form:"name"`
	Age  string `form:"age"`
}

func GetAllStudent(c *gin.Context) {

	res := fun.GetAllStudent()

	// read all
	var students []model.Student
	res.Scan(&students)

	c.JSON(200, gin.H{
		"message": students,
	})
}

func GetStudentById(c *gin.Context) {
	// TODO
	// get data
	res := fun.GetStudentById(c.Param("id"))

	// read data
	var student model.Student
	res.Scan(&student)
	c.JSON(200, gin.H{
		"message": student,
	})
}

func CreateStudent(c *gin.Context) {
	// TODO
	// get record value
	var student student
	c.ShouldBind(&student)

	// get
	name := student.Name
	age := student.Age

	ageInt, err := strconv.Atoi(age)

	if err != nil {
		panic(err)
	}

	// get data
	res := fun.CreateStudent(name, ageInt)

	// read data
	res.Scan(&student)

	c.JSON(200, gin.H{
		"message": "success store data",
		"data":    student,
	})
}

func UpdateStudent(c *gin.Context) {
	// TODO
	// get record value
	var student student
	c.ShouldBind(&student)

	// get
	name := student.Name
	age := student.Age
	id := c.Param("id")

	ageInt, err := strconv.Atoi(age)
	idInt, err := strconv.Atoi(id)

	if err != nil {
		panic(err)
	}

	// get data
	res := fun.UpdateStudent(idInt, name, ageInt)

	// read data
	res.Scan(&student)

	c.JSON(200, gin.H{
		"message": "success store data",
		"data":    student,
	})
}

func DeleteStudent(c *gin.Context) {
	id := c.Param("id")

	// get data
	fun.DeleteStudent(id)

	c.JSON(200, gin.H{
		"message": "success delete data",
		"id":      id,
	})
}
