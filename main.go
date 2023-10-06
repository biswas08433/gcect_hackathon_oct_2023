package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.Static("/static", "./public")

	// mux.Handle("/static/", http.StripPrefix("/static/", static_file_server))

	router.GET("/", Index)
	router.GET("/err", Err)
	router.GET("/login", Login)
	router.GET("/logout", Logout)
	router.GET("/signup", Signup)
	router.POST("/signup-account", SignupAccount) // POST
	router.POST("/authenticate", Authenticate)    // POST

	fmt.Println("Server", Version(), "started at", config.Address)
	router.Run(config.Address)
}
