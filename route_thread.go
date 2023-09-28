package main

import (
	"fmt"
	"net/http"

	"github.com/biswas08433/teachwise/data"
)

// GET /threads/new-thread
// Show the new thread form page
func NewThread(res http.ResponseWriter, req *http.Request) {
	if !IsLoggedIn(res, req) {
		http.Redirect(res, req, "/login", http.StatusFound)
	} else {
		GenerateHTML(res, nil, "layout", "private_navbar", "new_thread")
	}
}

// POST /thread/create-thread
// Creates the thread
func CreateThread(res http.ResponseWriter, req *http.Request) {
	logged_in, user := GetUserIfLoggedIn(res, req)
	if !logged_in {
		http.Redirect(res, req, "/login", http.StatusFound)
	} else {
		err := req.ParseForm()
		if err != nil {
			Danger(err, "Cannot parse form")
		}
		topic := req.PostFormValue("topic")
		if _, err := user.CreateThread(topic); err != nil {
			Danger(err, "Cannot create thread")
		}
		http.Redirect(res, req, "/", http.StatusFound)
	}
}

// GET /thread/read-thread?uuid=
// Show the details of the thread, including the posts and the form to write a post
func ReadThread(res http.ResponseWriter, req *http.Request) {
	vals := req.URL.Query()
	uuid := vals.Get("uuid")
	thread, _ := data.GetThreadByUuid(uuid)
	if false {
		ShowError(res, req, "")
	} else {
		fmt.Println(thread)
		if !IsLoggedIn(res, req) {
			GenerateHTML(res, &thread, "layout", "public_navbar", "public_thread")
		} else {
			GenerateHTML(res, &thread, "layout", "private_navbar", "private_thread")
		}
	}
}

// POST /thread/post
// Create the post
func CreatePost(res http.ResponseWriter, req *http.Request) {
	logged_in, user := GetUserIfLoggedIn(res, req)
	if !logged_in {
		http.Redirect(res, req, "/login", http.StatusFound)
	} else {
		err := req.ParseForm()
		if err != nil {
			Danger(err, "Cannot parse form")
		}
		body := req.PostFormValue("body")
		uuid := req.PostFormValue("uuid")
		thread, err := data.GetThreadByUuid(uuid)
		if err != nil {
			ShowError(res, req, "Cannot read thread")
		}
		if _, err := user.CreatePost(thread, body); err != nil {
			Danger(err, "Cannot create post")
		}
		url := fmt.Sprint("/thread/read?id=", uuid)
		http.Redirect(res, req, url, http.StatusFound)
	}
}
