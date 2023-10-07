package main

import (
	"log"

	"github.com/biswas08433/teachwise/data"
	"github.com/gin-gonic/gin"
)

// GET /err?msg=
// shows the error message page
func Err(ctx *gin.Context) {
	response_data := gin.H{
		"Title":        "Error",
		"ErrorMessage": ctx.Query("msg"),
	}
	if !IsLoggedIn(ctx) {
		GenerateHTML(ctx, response_data, "layout", "public_navbar", "error")
	} else {
		GenerateHTML(ctx, response_data, "layout", "private_navbar", "error")
	}
}

func Index(ctx *gin.Context) {
	response_data := gin.H{
		"Title": "TeachWise",
	}
	if IsLoggedIn(ctx) {
		_, user := GetUserIfLoggedIn(ctx)
		log.Println(user)
		response_data["User"] = user
		trending_teacher, err := data.TrendingTeachers()
		if err != nil {
			log.Fatalln("TrendingTeachers", err)
		}
		response_data["TrendingTeachers"] = trending_teacher
		log.Println(trending_teacher)
		GenerateHTML(ctx, response_data, "layout", "private_navbar", "dashboard")

	} else {
		GenerateHTML(ctx, response_data, "layout", "public_navbar", "index")
	}
	// }

}

func About(ctx *gin.Context) {
	if IsLoggedIn(ctx) {
		GenerateHTML(ctx, nil, "layout", "private_navbar", "about")
	} else {
		GenerateHTML(ctx, nil, "layout", "public_navbar", "about")
	}
}
