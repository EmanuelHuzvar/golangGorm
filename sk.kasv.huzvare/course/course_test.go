package course

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"hibernateGolang/sk.kasv.huzvare/structs"
	"log"
	"testing"
)

func setupTestDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open(structs.Dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}

	db.AutoMigrate(&structs.Course{})
	return db
}
func TestUseCourseDb(t *testing.T) {
	db := setupTestDB()
	defer teardownTestDB(db)

	UseCourseDb(db)
	// Since `UseCourseDb` doesn't return an error, you might want to check if the table exists
	if !db.Migrator().HasTable(&structs.Course{}) {
		t.Error("TestUseCourseDb: table for Course was not created")
	}
}

func TestAddCourseToInstructor(t *testing.T) {
	db := setupTestDB()
	defer teardownTestDB(db)

	// Assuming an instructor is already added to the database
	instructor := structs.Instructor{
		BaseModel:        structs.BaseModel{},
		FirstName:        "Marek",
		LastName:         "Les",
		Email:            "jan.jan@ja.sk",
		InstructorDetail: structs.InstructorDetail{},
		Courses:          nil,
	}
	db.Create(&instructor)

	newCourse := structs.Course{Title: "New Course"}
	err := AddCourseToInstructor(db, instructor.ID, newCourse)
	if err != nil {
		t.Errorf("TestAddCourseToInstructor: failed to add course: %v", err)
	}
}

func TestGetCoursesByInstructorID(t *testing.T) {
	db := setupTestDB()
	defer teardownTestDB(db)

	var courses []structs.Course
	instructor := structs.Instructor{
		BaseModel: structs.BaseModel{},
		FirstName: "S",
		LastName:  "Z",
		Email:     "P",
	}
	var err error
	courses, err = GetCoursesByInstructorID(db, instructor.ID)
	if err != nil {
		fmt.Printf("error is: ", err)
	} else {
		PrintCourses(courses)
	}
}

func teardownTestDB(db *gorm.DB) {
	db.Exec("DELETE FROM courses")
}
