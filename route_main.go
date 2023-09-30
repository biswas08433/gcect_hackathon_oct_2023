package main

import (
	"net/http"
)

// GET /err?msg=
// shows the error message page
func Err(res http.ResponseWriter, req *http.Request) {
	vals := req.URL.Query()

	if !IsLoggedIn(res, req) {
		GenerateHTML(res, vals.Get("msg"), "layout", "public_navbar", "error")
	} else {
		GenerateHTML(res, vals.Get("msg"), "layout", "private_navbar", "error")
	}
}

func Index(res http.ResponseWriter, req *http.Request) {
	// if err != nil {
	// 	ShowError(res, req, err.Error())
	// } else {
	if !IsLoggedIn(res, req) {
		GenerateHTML(res, nil, "layout", "public_navbar", "index")

	} else {
		GenerateHTML(res, nil, "layout", "private_navbar", "index")
	}
	// }

}
