package main

import (
	"net/http"

	"github.com/biswas08433/chatter/data"
)

// POST /authenticate
func Authenticate(res http.ResponseWriter, req *http.Request) {
	req.ParseForm()

	user, err := data.GetUserByEmail(req.PostFormValue("email"))
	if err != nil {
		Danger(err, "Cannot find user")
	}

	if user.Password == data.Encrypt(req.PostFormValue("password")) {
		session, err := user.CreateSession()
		if err != nil {
			Danger(err, "Cannot create session")
		}
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.Uuid,
			HttpOnly: true,
		}
		http.SetCookie(res, &cookie)
		http.Redirect(res, req, "/", http.StatusFound)
	} else {
		http.Redirect(res, req, "/login", http.StatusFound)
	}
}

func Login(res http.ResponseWriter, req *http.Request) {
	GenerateHTML(res, nil, "layout_login", "login")
}

// GET /logout
func Logout(res http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("_cookie")
	if err == http.ErrNoCookie {
		Warning(err, "Failed to get cookie")
	}
	data.DeleteSessionByUuid(cookie.Value)
	http.Redirect(res, req, "/", http.StatusFound)

}

// GET /signup
func Signup(res http.ResponseWriter, req *http.Request) {
	GenerateHTML(res, nil, "layout_login", "signup")
}

// POST /signup-account
func SignupAccount(res http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		Danger(err, "Cannot parse form")
	}
	user := data.User{
		Name:     req.PostFormValue("name"),
		Email:    req.PostFormValue("email"),
		Password: req.PostFormValue("password"),
	}
	if err := user.Create(); err != nil {
		Danger(err, "Cannot create user")
	}
	http.Redirect(res, req, "/login", http.StatusFound)
}
