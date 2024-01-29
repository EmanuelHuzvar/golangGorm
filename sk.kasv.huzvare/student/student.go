package student

import (
	"gorm.io/gorm"
	"hibernateGolang/sk.kasv.huzvare/structs"
)

func GetStudentByID(db *gorm.DB, id int) structs.Student {
	var student structs.Student
	db.First(&student, id)
	return student
}

func CreateStudent(db *gorm.DB, student *structs.Student) {
	db.Create(student)
}

func UpdateStudentByID(db *gorm.DB, id int, updatedData structs.Student) {
	var student structs.Student
	db.First(&student, id)
	db.Model(&student).Updates(updatedData)
}

func DeleteStudentByID(db *gorm.DB, id int) {
	var student structs.Student
	db.Delete(&student, id)
}
func GetCoursesForStudent(db *gorm.DB, studentID int) ([]structs.Course, error) {
	var courses []structs.Course
	err := db.Table("course_student").Select("course.*").
		Joins("join course on course.id = course_student.course_id").
		Where("course_student.student_id = ?", studentID).
		Scan(&courses).Error

	if err != nil {
		return nil, err
	}
	return courses, nil
}
func AddStudentToCourse(db *gorm.DB, studentID int, courseID int) error {
	courseStudent := structs.CourseStudent{
		CourseID:  courseID,
		StudentID: studentID,
	}

	if err := db.Create(&courseStudent).Error; err != nil {
		return err
	}
	return nil
}
func RemoveStudentFromCourse(db *gorm.DB, studentID int, courseID uint) error {
	if err := db.Where("student_id = ? AND course_id = ?", studentID, courseID).
		Delete(&structs.CourseStudent{}).Error; err != nil {
		return err
	}
	return nil
}
func SearchStudents(db *gorm.DB, firstName string, lastName string, email string) ([]structs.Student, error) {
	var students []structs.Student

	query := db.Model(&structs.Student{})

	if firstName != "" {
		query = query.Where("first_name LIKE ?", "%"+firstName+"%")
	}
	if lastName != "" {
		query = query.Where("last_name LIKE ?", "%"+lastName+"%")
	}
	if email != "" {
		query = query.Where("email LIKE ?", "%"+email+"%")
	}

	err := query.Find(&students).Error
	if err != nil {
		return nil, err
	}

	return students, nil
}
