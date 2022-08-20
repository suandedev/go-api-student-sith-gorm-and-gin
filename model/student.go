package model

type Student struct {
	ID   int    `gorm:"primaryKey autoIncrement"`
	Name string `form:"name"`
	Age  int    `form:"age"`
}
