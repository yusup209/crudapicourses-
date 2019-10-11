package controllers

import (
	"crudapicourses/configs"
	"crudapicourses/helpers"
	"crudapicourses/models"
	repositories "crudapicourses/repositories"
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/valyala/fasthttp"
)

func GetAllCourseCategory(ctx *fasthttp.RequestCtx) {
	var data = make(map[string]interface{})

	// cat, err := repositories.FindAllCoursesCategories(configs.DB)
	cat, err := repositories.FindAllCourseCategories(configs.DB)

	if err != nil {
		fmt.Printf("ada error di get all course category\n%v", err)
		data["error"] = true
		data["message"] = err
		helpers.JSONify(ctx, data)
	}

	data["course_categories"] = cat

	helpers.JSONify(ctx, data)

}

func CreateCourseCategory(ctx *fasthttp.RequestCtx) {
	postValue := ctx.PostBody()
	course := models.CourseCategory{}

	var data = make(map[string]interface{})

	if err := json.Unmarshal(postValue, &course); err != nil {
		log.Println("ada panic di create course category")
		data["error"] = true
		data["message"] = err
		helpers.JSONify(ctx, data)
		panic(err)
	}

	if err := repositories.CreateCourseCategory(configs.DB, course); err != nil {
		log.Println("gagal menyimpan data ke dalam database")
		panic(err)
		data["error"] = true
		data["message"] = err
		helpers.JSONify(ctx, data)
	} else {
		data["success"] = true
		data["message"] = "sukses menambahkan data course"
		helpers.JSONify(ctx, data)
	}

}

func UpdateCourseCategory(ctx *fasthttp.RequestCtx) {
	postValue := ctx.PostBody()
	course := models.CourseCategory{}

	var data = make(map[string]interface{})

	if err := json.Unmarshal(postValue, &course); err != nil {
		log.Println("ada panic di update course category")
		data["error"] = true
		data["message"] = err
		helpers.JSONify(ctx, data)
		panic(err)
	}

	if id, err := repositories.UpdateCourseCategory(configs.DB, course); err != nil {
		log.Println("gagal menyimpan data ke dalam database")
		panic(err)
		data["error"] = true
		data["message"] = err
		helpers.JSONify(ctx, data)
	} else {
		data["success"] = true
		data["message"] = fmt.Sprintf("%s %s", "sukses mengupdate data course category", strconv.Itoa(int(id)))
		helpers.JSONify(ctx, data)
	}
}

func DeleteCourseCategory(ctx *fasthttp.RequestCtx) {
	var err error
	var data = make(map[string]interface{})
	id := fmt.Sprintf("%v", ctx.UserValue("id"))

	courseID, err := strconv.Atoi(id)
	if err != nil {
		fmt.Printf("gagal konversi courseid(interface) ke int")
	}

	err = repositories.DeleteCourseCategory(configs.DB, uint(courseID))
	if err != nil {
		log.Println("gagal delete data category")
		data["error"] = true
		data["message"] = err
	} else {
		data["success"] = true
		data["message"] = "sukses menghapus data category"
	}

	helpers.JSONify(ctx, data)

}
