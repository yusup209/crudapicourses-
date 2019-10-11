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

func GetCourses(ctx *fasthttp.RequestCtx) {
	//membuat object map dengan make
	var data = make(map[string]interface{})

	//isi datanya shaaay
	// data["data"] = "halo"
	// data["products"] = []string{"AHhhha", "ahhhahhhhhaa", "hhahaahahhah"}
	// ctx.Response.Header.Set("Content-Type", "application/json")
	//convert to jSon
	// res, err := json.Marshal(data)

	// if err != nil {
	// 	log.Println("PANIC PANIC... (get course)")
	// 	panic(err)
	// }
	//write to output
	courses, err := repositories.FindAllCourses(configs.DB)

	if err != nil {
		panic(err)
	}

	data["courses"] = courses

	// res, err := json.Marshal(data)

	// ctx.Write(res)
	// ctx.SetContentType("application/json")
	// ctx.SetStatusCode(fasthttp.StatusOK)
	helpers.JSONify(ctx, data)
}

func GetCourseByID(ctx *fasthttp.RequestCtx) {
	var data = make(map[string]interface{})
	id := fmt.Sprintf("%v", ctx.UserValue("id"))
	id_finale, err := strconv.Atoi(id)
	courses, err := repositories.FindCourseByID(configs.DB, uint(id_finale))

	if err != nil {
		panic(err)
	}

	data["courses"] = courses
	helpers.JSONify(ctx, data)
}

func CreateCourses(ctx *fasthttp.RequestCtx) {
	postValue := ctx.PostBody()
	course := models.Course{}

	fmt.Println("create course")
	fmt.Printf("course : %v", postValue)
	log.Println(postValue)

	var data = make(map[string]interface{})

	if err := json.Unmarshal(postValue, &course); err != nil {
		log.Println("ada panic di create course")
		data["error"] = true
		data["message"] = err
		helpers.JSONify(ctx, data)
		panic(err)
	}

	if err := repositories.CreateCourses(configs.DB, course); err != nil {
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

func UpdateCourses(ctx *fasthttp.RequestCtx) {
	postValue := ctx.PostBody()
	course := models.Course{}

	var data = make(map[string]interface{})

	if err := json.Unmarshal(postValue, &course); err != nil {
		log.Println("ada panic di update course")
		data["error"] = true
		data["message"] = err
		helpers.JSONify(ctx, data)
		panic(err)
	}

	if id, err := repositories.UpdateCourse(configs.DB, course); err != nil {
		log.Println("gagal menyimpan data ke dalam database")
		panic(err)
		data["error"] = true
		data["message"] = err
		helpers.JSONify(ctx, data)
	} else {
		data["success"] = true
		data["message"] = fmt.Sprintf("%s %s", "sukses mengupdate data course", strconv.Itoa(int(id)))
		helpers.JSONify(ctx, data)
	}
}

func DeleteCourses(ctx *fasthttp.RequestCtx) {
	var err error
	var data = make(map[string]interface{})
	id := fmt.Sprintf("%v", ctx.UserValue("id"))

	courseID, err := strconv.Atoi(id)
	if err != nil {
		fmt.Printf("gagal konversi courseid(interface) ke int")
	}

	_, err = repositories.DeleteCourse(configs.DB, courseID)
	if err != nil {
		log.Println("gagal delete data course")
		data["error"] = true
		data["message"] = err
	} else {
		data["success"] = true
		data["message"] = "sukses menghapus data course"
	}

	helpers.JSONify(ctx, data)

}
