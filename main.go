package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	static_file_server := http.FileServer(http.Dir(config.Static))

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", static_file_server))

	mux.HandleFunc("/", Index)
	mux.HandleFunc("/err", Err)
	mux.HandleFunc("/login", Login)
	mux.HandleFunc("/logout", Logout)
	mux.HandleFunc("/signup", Signup)
	mux.HandleFunc("/signup-account", SignupAccount) // POST
	mux.HandleFunc("/authenticate", Authenticate)    // POST
	mux.HandleFunc("/thread/read-thread", ReadThread)
	mux.HandleFunc("/thread/new-thread", NewThread)
	mux.HandleFunc("/thread/create-thread", CreateThread) // POST
	mux.HandleFunc("/thread/create-post", CreatePost)     // POST

	server := &http.Server{
		Addr:           config.Address,
		Handler:        mux,
		ReadTimeout:    time.Duration(config.ReadTimeout * int64(time.Second)),
		WriteTimeout:   time.Duration(config.WriteTimeout * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Println("Server", Version(), "started at", config.Address)
	server.ListenAndServe()
}
