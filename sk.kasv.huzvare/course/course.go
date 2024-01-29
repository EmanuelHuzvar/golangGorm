package course

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"hibernateGolang/sk.kasv.huzvare/structs"
)

func UseCourseDb(db *gorm.DB) {
	err := db.AutoMigrate(structs.Course{})
	if err != nil {
		return
	}
}
func AddCourseToInstructor2(db *gorm.DB, instructorID uint, newCourse structs.Course) error {
	newCourse.InstructorID = instructorID
	if err := db.Create(&newCourse).Error; err != nil {
		return err
	}
	return nil
}
func AddCourseToInstructor(db *gorm.DB, instructorID uint, newCourse structs.Course) error {

	var inst structs.Instructor
	if err := db.First(&inst, instructorID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {

			return fmt.Errorf("instructor with ID %d does not exist", instructorID)
		}

		return err
	}

	newCourse.InstructorID = instructorID
	if err := db.Create(&newCourse).Error; err != nil {
		return err
	}
	return nil
}
func GetCoursesByInstructorID(db *gorm.DB, instructorID uint) ([]structs.Course, error) {
	var courses []structs.Course
	if err := db.Where("instructor_id = ?", instructorID).Find(&courses).Error; err != nil {
		return nil, err
	}
	return courses, nil
}

func PrintCourses(courses []structs.Course) {
	for i := 0; i < len(courses); i++ {

		fmt.Print(courses[i].Title, "  ")
	}
}
func UpdateCourseByID(db *gorm.DB, courseID uint, newTitle string) error {
	var course structs.Course
	if err := db.First(&course, courseID).Error; err != nil {
		return err
	}

	course.Title = newTitle
	if err := db.Save(&course).Error; err != nil {
		return err
	}
	return nil
}
func DeleteCourseByID(db *gorm.DB, courseID uint) error {
	if err := db.Delete(&structs.Course{}, courseID).Error; err != nil {
		return err
	}
	return nil
}
func GetCourseByID(db *gorm.DB, courseID uint) (*structs.Course, error) {
	var course structs.Course
	err := db.First(&course, courseID).Error
	if err != nil {
		return nil, err
	}
	return &course, nil
}
func GetAllCourses(db *gorm.DB) ([]structs.Course, error) {
	var courses []structs.Course
	err := db.Find(&courses).Error
	if err != nil {
		return nil, err
	}
	return courses, nil
}
func GetStudentsInCourse(db *gorm.DB, courseID uint) ([]structs.Student, error) {
	var students []structs.Student
	err := db.Joins("JOIN course_student ON course_student.student_id = students.id").
		Where("course_student.course_id = ?", courseID).Find(&students).Error
	if err != nil {
		return nil, err
	}
	return students, nil
}
func FindCoursesByTitle(db *gorm.DB, title string) ([]structs.Course, error) {
	var courses []structs.Course
	err := db.Where("title LIKE ?", "%"+title+"%").Find(&courses).Error
	if err != nil {
		return nil, err
	}
	return courses, nil
}
func AssignInstructorToCourse(db *gorm.DB, courseID, instructorID int) error {
	err := db.Model(&structs.Course{}).Where("id = ?", courseID).
		Update("instructor_id", instructorID).Error
	if err != nil {
		return err
	}
	return nil
}

func RemoveInstructorFromCourse(db *gorm.DB, courseID uint) error {
	err := db.Model(&structs.Course{}).Where("id = ?", courseID).
		Update("instructor_id", nil).Error
	if err != nil {
		return err
	}
	return nil
}
