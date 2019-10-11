package routes

import (
	ctrl "crudapicourses/controllers"
	"fmt"
	"log"

	cors "github.com/AdhityaRamadhanus/fasthttpcors"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

func handleIndex(ctx *fasthttp.RequestCtx) {
	fmt.Fprintln(ctx, "<marquee>Helllooooo DangD03t</marquee>")
}

func handleHello(ctx *fasthttp.RequestCtx) {
	param := ctx.UserValue("nama")

	fmt.Fprintln(ctx, "param :", param)
}

func Setup() {
	//buat setup
	fmt.Println("routes setup")

	//untuk routing (hmm.. mirip express ya yg ada di nodejs)
	r := fasthttprouter.New()
	r.GET("/", handleIndex)
	r.GET("/hello/:nama", handleHello)

	// untuk course
	r.GET("/api/course", ctrl.GetCourses)
	r.GET("/api/course/:id", ctrl.GetCourseByID)
	r.POST("/api/course/create", ctrl.CreateCourses)
	r.PUT("/api/course/update", ctrl.UpdateCourses)
	r.DELETE("/api/course/delete/:id", ctrl.DeleteCourses)

	//untuk course category
	r.GET("/api/course_category", ctrl.GetAllCourseCategory)
	r.POST("/api/course_category/create", ctrl.CreateCourseCategory)
	r.PUT("/api/course_category/update", ctrl.UpdateCourseCategory)
	r.DELETE("/api/course_category/delete/:id", ctrl.DeleteCourseCategory)

	//portnya
	listenAddr := ":8000"
	fmt.Println("Run fasthttttttp di port", listenAddr)

	withCORS := cors.NewCorsHandler(cors.Options{
		// if you leave allowedOrigins empty then fasthttpcors will treat it as "*"
		AllowedOrigins: []string{"*"}, // Only allow example.com to access the resource
		// if you leave allowedHeaders empty then fasthttpcors will accept any non-simple headers
		AllowedHeaders: []string{"x-something-client", "Content-Type", "content-type"}, // only allow x-something-client and Content-Type in actual request

		// if you leave this empty, only simple method will be accepted
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"}, // only allow get or post to resource
		AllowCredentials: false,                                    // resource doesn't support credentials
		AllowMaxAge:      5600,                                     // cache the preflight result
		Debug:            true,
	})
	if err := fasthttp.ListenAndServe(listenAddr, withCORS.CorsMiddleware(r.Handler)); err != nil {
		log.Fatalf("Error in ListenAndServe: %s", err)
	}
}
