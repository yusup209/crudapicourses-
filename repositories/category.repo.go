package repositories

import (
	models "crudapicourses/models"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
)

func FindAllCourseCategories(db *gorm.DB) ([]models.CourseCategory, error) {
	categories := []models.CourseCategory{}

	if err := db.Find(&categories).Error; err != nil {
		fmt.Printf("ada error pas get all data category\n%v", err)
		return nil, err
	}

	return categories, nil
}

func CreateCourseCategory(db *gorm.DB, course models.CourseCategory) error {
	if err := db.Create(&course).Error; err != nil {
		return err
	} else {
		return nil
	}
}

func UpdateCourseCategory(db *gorm.DB, course models.CourseCategory) (uint, error) {
	categoryDB := models.CourseCategory{}

	if db.Where("id = ?", course.ID).First(&categoryDB).RecordNotFound() {
		log.Println("tidak ada data yang ingin anda update")
		return 0, nil
	}

	categoryDB.Name = course.Name

	if err := db.Save(&categoryDB).Error; err != nil {
		fmt.Printf("gagal mengupdate data category\n%v", err)
		return 0, err
	}

	return course.ID, nil

}

func DeleteCourseCategory(db *gorm.DB, category_id uint) error {
	courseCategory := models.CourseCategory{}

	if err := db.Find(&courseCategory, category_id).Error; err != nil {
		fmt.Printf("tidak ada data yang ingin anda hapus\n%v", err)
		return err
	}

	if err := db.Where("id = ? AND deleted_at IS NULL", int(category_id)).Delete(&courseCategory).Error; err != nil {
		fmt.Printf("gagal menghapus data category\n%v", err)
		return err
	}

	return nil
}
