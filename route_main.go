package main

import (
	"github.com/gin-gonic/gin"
)

// GET /err?msg=
// shows the error message page
func Err(ctx *gin.Context) {
	data := gin.H{
		"Title":        "Error",
		"ErrorMessage": ctx.Query("msg"),
	}
	if !IsLoggedIn(ctx) {
		GenerateHTML(ctx, data, "layout", "public_navbar", "error")
	} else {
		GenerateHTML(ctx, data, "layout", "private_navbar", "error")
	}
}

func Index(ctx *gin.Context) {
	data := gin.H{
		"Title": "TeachWise",
	}
	if !IsLoggedIn(ctx) {
		GenerateHTML(ctx, data, "layout", "public_navbar", "index")

	} else {
		GenerateHTML(ctx, data, "layout", "private_navbar", "index")
	}
	// }

}
