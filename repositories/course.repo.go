package repositories

import (
	models "crudapicourses/models"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
)

//jika err != nil, variabel error ada isinya
//selain itu data yg ada di struct ada isinya
func FindAllCourses(db *gorm.DB) ([]models.Course, error) {
	//membuat variabel untuk menampung
	courses := []models.Course{}

	//select all from db with GOrm
	if err := db.Preload("CourseCategory").Find(&courses).Error; err != nil {
		// if err := db.Find(&courses).Error; err != nil {
		return nil, err
	}

	return courses, nil
}

func FindCourseByID(db *gorm.DB, id uint) ([]models.Course, error) {
	//membuat variabel untuk menampung
	courses := []models.Course{}

	//select all from db with GOrm
	if err := db.Find(&courses, id).Error; err != nil {
		// if err := db.Find(&courses).Error; err != nil {
		return nil, err
	}

	return courses, nil
}

func CreateCourses(db *gorm.DB, course models.Course) error {
	if err := db.Create(&course).Error; err != nil {
		return err
	} else {
		return nil
	}
}

func UpdateCourse(db *gorm.DB, course models.Course) (uint, error) {
	courseDB := models.Course{}

	if db.Where("id = ?", course.ID).First(&courseDB).RecordNotFound() {
		fmt.Println("Maaf, user yang ingin anda update tidak ada :v")
		return 0, nil
	}

	// courseDB.CourseCategoryID = course.CourseCategoryID
	courseDB.Name = course.Name
	courseDB.Description = course.Description
	courseDB.PricePerHour = course.PricePerHour
	courseDB.Avatar = course.Avatar

	if err := db.Save(&courseDB).Error; err != nil {
		fmt.Printf("Maaf, gagal dalam update data user\n%v", err)
		return 0, err
	}

	return course.ID, nil

}

func DeleteCourse(db *gorm.DB, course_id int) (uint, error) {
	course := models.Course{}

	if db.Find(&course, course_id).RecordNotFound() {
		log.Println("Data yang ingin anda hapus tidak ada")
		return 0, nil
	}

	if err := db.Where("id = ? AND deleted_at IS NULL", uint(course_id)).Delete(&course).Error; err != nil {
		log.Println("Gagal menghapus data course di repo")
	}

	return course.ID, nil

}
