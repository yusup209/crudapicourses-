package models

import "time"

type Course struct {
	ID               uint           `gorm:"size:11;primary_key;AUTO_INCREMENT" json:"id"`
	CourseCategoryID uint           `gorm:"ForeignKey:ID" json:"course_category_id"`
	Name             string         `gorm:"size:100;not null" json:"name"`
	Description      string         `gorm:"size:100;type:text" json:"description"`
	PricePerHour     int            `gorm:"size:100" json:"price_per_hour"`
	Avatar           string         `gorm:"type:text" json:"avatar"`
	CourseCategory   CourseCategory `json:"course_category"`
	CreatedAt        time.Time      `json:"-"`
	UpdatedAt        time.Time      `json:"-"`
	DeletedAt        *time.Time     `json:"-"`
}

type CourseCategory struct {
	ID        uint       `gorm:"size:11;primary_key;AUTO_INCREMENT" json:"id"`
	Name      string     `gorm:"size:100;not null" json:"name"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
}
