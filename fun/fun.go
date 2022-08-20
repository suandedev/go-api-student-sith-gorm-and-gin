package fun

import (
	"new2/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConDB() *gorm.DB {
	// TODO
	db, err := gorm.Open(sqlite.Open("test.sqlite"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// auto migrate
	db.AutoMigrate(&model.Student{})

	// create
	// db.Create(&model.Student{ID: 1, Name: "A", Age: 20})

	return db
}

func GetAllStudent() *gorm.DB {
	// connect db
	db := ConDB()
	res := db.Raw("SELECT * FROM students")
	return res
}

func GetStudentById(id string) *gorm.DB {
	// connect db
	db := ConDB()
	res := db.First(&model.Student{}, id)
	return res
}

func CreateStudent(name string, age int) *gorm.DB {
	db := ConDB()
	res := db.Create(&model.Student{Name: name, Age: age})
	return res
}

func UpdateStudent(id int, name string, age int) *gorm.DB {
	// connect db
	db := ConDB()
	res := db.Model(&model.Student{}).Where("id = ?", id).Updates(model.Student{Name: name, Age: age})

	return res
}

func DeleteStudent(id string) *gorm.DB {
	// connect db
	db := ConDB()
	res := db.Delete(&model.Student{}, id)
	return res
}
