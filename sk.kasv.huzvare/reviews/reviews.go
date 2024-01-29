package reviews

import (
	"gorm.io/gorm"
	"hibernateGolang/sk.kasv.huzvare/structs"
)

func AddReview(db *gorm.DB, review *structs.Review) error {
	//
	if err := db.Create(review).Error; err != nil {
		return err
	}
	return nil
}

func UpdateReviewByID(db *gorm.DB, reviewID uint, updatedData structs.Review) error {
	var review structs.Review
	if err := db.First(&review, reviewID).Error; err != nil {
		return err // Review not found
	}

	if err := db.Model(&review).Updates(updatedData).Error; err != nil {
		return err // Error while updating
	}
	return nil
}
func GetCourseWithReviews(db *gorm.DB, courseID uint) (*structs.CourseWithReviews, error) {
	var course structs.Course
	var reviews []structs.Review

	// Find the course by ID
	if err := db.First(&course, courseID).Error; err != nil {
		return nil, err
	}

	// Get reviews for the course
	err := db.Where("course_id = ?", courseID).Find(&reviews).Error
	if err != nil {
		return nil, err
	}

	// Check if there are no reviews
	if len(reviews) == 0 {
		reviews = append(reviews, structs.Review{Comment: "No rating yet"})
	}

	return &structs.CourseWithReviews{Course: course, Reviews: reviews}, nil
}
