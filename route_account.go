package main

import (
	"net/http"

	"github.com/biswas08433/teachwise/data"
	"github.com/gin-gonic/gin"
)

// POST /authenticate
func Authenticate(ctx *gin.Context) {
	user, err := data.GetUserByEmail(ctx.PostForm("email"))
	if err != nil {
		Danger(err, "Cannot find user")
	}

	if user.Password == data.Encrypt(ctx.PostForm("password")) {
		session, err := user.CreateSession()
		if err != nil {
			Danger(err, "Cannot create session")
		}
		ctx.SetCookie("session_cookie", session.Uuid, 2*24*3600, "/", "localhost", true, true)
		ctx.Redirect(http.StatusFound, "/")
	} else {
		ctx.Redirect(http.StatusFound, "/login")
	}
}

func Login(ctx *gin.Context) {
	data := gin.H{
		"Title": "Login",
	}
	GenerateHTML(ctx, data, "layout_login", "login")
}

// GET /logout
func Logout(ctx *gin.Context) {
	cookie_value, err := ctx.Cookie("session_cookie")
	if err == http.ErrNoCookie {
		Warning(err, "Failed to get cookie")
	}
	data.DeleteSessionByUuid(cookie_value)
	ctx.Redirect(http.StatusFound, "/")
}

// GET /signup
func Signup(ctx *gin.Context) {
	data := gin.H{
		"Title": "Signup",
	}
	GenerateHTML(ctx, data, "layout_login", "signup")
}

// POST /signup-account
func SignupAccount(ctx *gin.Context) {
	user := data.User{
		FirstName: ctx.PostForm("first-name"),
		LastName:  ctx.PostForm("last-name"),
		Email:     ctx.PostForm("email"),
		Password:  ctx.PostForm("password"),
	}
	rp := ctx.PostForm("retype-password")
	if rp != user.Password {
		ctx.Redirect(http.StatusNotAcceptable, "/signup")
	}
	if err := user.Create(); err != nil {
		Danger(err, "Cannot create user")
	}
	ctx.Redirect(http.StatusFound, "/login")
}

// GET /user/profile
func Profile(ctx *gin.Context) {
	if IsLoggedIn(ctx) {
		response_data := gin.H{
			"Title": "Profile",
		}
		_, user := GetUserIfLoggedIn(ctx)
		response_data["User"] = user
		GenerateHTML(ctx, response_data, "layout", "private_navbar", "profile")
	} else {
		ctx.Redirect(http.StatusFound, "/login")
	}
}

// GET /user/edit-profile
func EditProfile(ctx *gin.Context) {
	if IsLoggedIn(ctx) {
		response_data := gin.H{
			"Title": "Edit Profile",
		}
		_, user := GetUserIfLoggedIn(ctx)
		response_data["User"] = user
		GenerateHTML(ctx, response_data, "layout", "private_navbar", "edit_profile")
	} else {
		ctx.Redirect(http.StatusFound, "/login")
	}
}

// POST /user/edit-profile
func UpdateProfile(ctx *gin.Context) {

}
