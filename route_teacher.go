package main

import (
	"log"
	"strings"

	"github.com/biswas08433/teachwise/data"
	"github.com/gin-gonic/gin"
)

func Teachers(ctx *gin.Context) {
	subject_query := strings.ToLower(ctx.Query("subject"))
	// city_query = strings.ToLower(ctx.Query("city"))

	if subject_query != "" {
		teachers, err := data.GetTeachersBySubject(subject_query)
		if err != nil {
			log.Println(err.Error())
		}
		response_data := gin.H{
			"TItle":    subject_query,
			"Teachers": teachers,
		}
		GenerateHTML(ctx, response_data, "layout", "private_navbar", "teachers_subject")
	}
}
