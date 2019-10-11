package main

import (
	"crudapicourses/configs"
	"crudapicourses/models"
	"crudapicourses/routes"
	"fmt"
)

func main() {
	//run migration dulu
	configs.DB.Debug().AutoMigrate(&models.Course{}, &models.CourseCategory{})
	fmt.Println("Helloooo")
	routes.Setup()
}
