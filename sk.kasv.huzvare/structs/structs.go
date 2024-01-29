package structs

var Dsn = "hbstudent:hbstudent@tcp(127.0.0.1:3306)/hb_student_tracker1?charset=utf8mb4&parseTime=True&loc=Local"

type BaseModel struct {
	ID uint `gorm:"primary_key"`
}
type Course struct {
	BaseModel
	Title        string `gorm:"size:128"`
	InstructorID uint   `gorm:"index"`
}
type InstructorUpdate struct {
	BaseModel
	FirstName string
	LastName  string
	Email     string
}
type Instructor struct {
	BaseModel
	FirstName          string
	LastName           string
	Email              string
	InstructorDetailID uint
	InstructorDetail   InstructorDetail
	Courses            []Course `gorm:"foreignkey:InstructorID"`
}
type InstructorDetail struct {
	BaseModel
	YoutubeChannel string
	Hobby          string
}

type Student struct {
	ID        int    `gorm:"primaryKey"`
	Email     string `gorm:"column:email"`
	FirstName string `gorm:"column:first_name"`
	LastName  string `gorm:"column:last_name"`
}
type CourseStudent struct {
	CourseID  int `gorm:"primaryKey"`
	StudentID int `gorm:"foreignkey:student_id"`
}
type Review struct {
	ID       uint `gorm:"primaryKey"`
	CourseID uint `gorm:"not null"`
	Comment  string
	Rating   uint `gorm:"check:rating >= 0 AND rating <= 5"`
}
type CourseWithReviews struct {
	Course  Course
	Reviews []Review
}

func (CourseStudent) TableName() string {
	return "course_student"
}
func (Review) TableName() string {
	return "reviews"
}

func (Course) TableName() string {
	return "course"
}
func (Instructor) TableName() string {
	return "instructor"
}
func (InstructorDetail) TableName() string {
	return "instructor_detail"
}
func (Student) TableName() string {
	return "student"
}
