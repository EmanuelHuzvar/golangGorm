package instructor

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"hibernateGolang/sk.kasv.huzvare/structs"
)

func UseInstructorDb(db *gorm.DB) {
	err := db.AutoMigrate(structs.Instructor{}, structs.InstructorDetail{})
	if err != nil {
		return
	}
}

func AddInstructor(db *gorm.DB, instructor structs.Instructor) error {
	fmt.Println(instructor.InstructorDetail.ID)
	if instructor.InstructorDetail.Hobby == "" || instructor.InstructorDetail.YoutubeChannel == "" {
		newDetail := structs.InstructorDetail{}
		if err := db.Create(&newDetail).Error; err != nil {
			return err
		}
		instructor.InstructorDetailID = newDetail.ID
	} else {
		if err := db.Create(&instructor.InstructorDetail).Error; err != nil {
			return err
		}
		instructor.InstructorDetailID = instructor.InstructorDetail.ID
	}

	if err := db.Create(&instructor).Error; err != nil {
		return err
	}

	return nil
}
func GetInstructorByID(db *gorm.DB, id uint) (structs.Instructor, error) {
	var instructor structs.Instructor
	result := db.Preload("InstructorDetail").First(&instructor, id)
	if result.Error != nil {
		return structs.Instructor{}, result.Error
	}
	return instructor, nil
}
func GetInstructorsWithYouTubeChannel(db *gorm.DB) ([]structs.Instructor, error) {
	var instructors []structs.Instructor
	//mozme pouzit IS NOT NULL
	err := db.Joins("InstructorDetail").Where("`InstructorDetail`.`youtube_channel` <> '' ").Find(&instructors).Error
	if err != nil {
		return nil, err
	}
	return instructors, nil
}

func DeleteInstructorById(db *gorm.DB, id uint) {
	var instructor structs.Instructor
	inst, _ := GetInstructorByID(db, id)
	instructor.InstructorDetailID = inst.InstructorDetailID
	db.Delete(&instructor, id)
	deleteInstructorDetailById(db, instructor.InstructorDetailID)
}
func deleteInstructorDetailById(db *gorm.DB, detailId uint) error {
	if err := db.Delete(&structs.InstructorDetail{}, detailId).Error; err != nil {
		return err
	}
	return nil
}

func deleteHobbyById(db *gorm.DB, id uint) {
	var instructorDetail structs.InstructorDetail
	db.Delete(&instructorDetail, id)
}
func UpdateInstructorHobbyOrChannel(db *gorm.DB, instructorID uint, newHobby string, newChannel string) error {

	var instructor structs.Instructor
	result := db.Preload("InstructorDetail").First(&instructor, instructorID)
	if result.Error != nil {
		return result.Error
	}
	if len(newHobby) != 0 {
		instructor.InstructorDetail.Hobby = newHobby
	}
	if len(newChannel) != 0 {
		instructor.InstructorDetail.YoutubeChannel = newChannel
	}

	result = db.Save(&instructor.InstructorDetail)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func GetInstructorDetailByID(db *gorm.DB, instructorID uint) (*structs.InstructorDetail, error) {
	var instructor structs.Instructor
	result := db.Preload("InstructorDetail").First(&instructor, instructorID)
	if result.Error != nil {
		return nil, result.Error
	}

	return &instructor.InstructorDetail, nil
}
func FindInstructorByHobby(db *gorm.DB, hobby string) ([]structs.Instructor, error) {
	var instructors []structs.Instructor
	err := db.Preload("InstructorDetail", "hobby LIKE ?", hobby).Find(&instructors).Error
	if err != nil {
		return nil, err
	}
	var filteredInstructors []structs.Instructor
	for _, instructor := range instructors {
		if instructor.InstructorDetail.Hobby == hobby {
			filteredInstructors = append(filteredInstructors, instructor)
		}
	}

	return filteredInstructors, nil
}

func ListCoursesByInstructorID(db *gorm.DB, instructorID uint) ([]structs.Course, error) {
	var courses []structs.Course
	if err := db.Where("instructor_id = ?", instructorID).Find(&courses).Error; err != nil {
		return nil, err
	}
	return courses, nil
}
func UpdateInstructorDetail(db *gorm.DB, instructorID uint, newDetail structs.InstructorDetail) error {
	var instructor structs.Instructor
	if err := db.First(&instructor, instructorID).Error; err != nil {
		return err
	}

	if instructor.InstructorDetailID == 0 {
		return errors.New("instructor has no associated instructor_detail")
	}

	newDetail.ID = instructor.InstructorDetailID

	if err := db.Model(&structs.InstructorDetail{}).Where("id = ?", newDetail.ID).Updates(newDetail).Error; err != nil {
		return err
	}

	return nil
}
func AddCourseToInstructor(db *gorm.DB, instructorID int, courseID int) error {
	var instructor structs.Instructor
	if err := db.First(&instructor, instructorID).Error; err != nil {
		return err
	}

	err := db.Model(&structs.Course{}).Where("id = ?", courseID).
		Update("instructor_id", instructorID).Error

	return err
}
func RemoveInstructorFromCourse(db *gorm.DB, courseID uint) error {
	err := db.Model(&structs.Course{}).Where("id = ?", courseID).
		Update("instructor_id", gorm.Expr("NULL")).Error

	return err
}
func UpdateInstructor(db *gorm.DB, id uint, updatedData structs.InstructorUpdate) error {
	var instructor structs.Instructor

	if err := db.First(&instructor, id).Error; err != nil {
		return err
	}

	if err := db.Model(&instructor).Updates(structs.Instructor{
		FirstName: updatedData.FirstName,
		LastName:  updatedData.LastName,
		Email:     updatedData.Email,
	}).Error; err != nil {
		return err
	}

	return nil
}
func FindInstructorByDetails(db *gorm.DB, firstName, lastName, email string) ([]structs.Instructor, error) {
	var instructors []structs.Instructor
	result := db.Where("first_name = ? OR last_name = ? OR email = ?", firstName, lastName, email).Find(&instructors)
	if result.Error != nil {
		return nil, result.Error
	}
	return instructors, nil
}
