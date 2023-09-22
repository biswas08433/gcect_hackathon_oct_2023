package main

import (
	"net/http"

	"github.com/biswas08433/chatter/data"
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
	// fmt.Println("Index Requested") // Debug
	threads, err := data.Threads()
	if err != nil {
		ShowError(res, req, err.Error())
	} else {
		if !IsLoggedIn(res, req) {
			GenerateHTML(res, threads, "layout", "public_navbar", "index")

		} else {
			GenerateHTML(res, threads, "layout", "private_navbar", "index")
		}
	}

}
