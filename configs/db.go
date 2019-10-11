package configs

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func init() {
	var err error
	DB, err = gorm.Open("mysql", "root:@/crudapicourses?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err)
	}

	DB.LogMode(true)
	fmt.Println("K0n3ksi database sukses..")
}
