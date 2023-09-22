package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/biswas08433/chatter/data"
)

type Configuration struct {
	Address      string
	ReadTimeout  int64
	WriteTimeout int64
	Static       string
	Version      string
}

var logger *log.Logger
var config Configuration

func loadConfig() {
	file, err := os.Open("server_config.json")
	if err != nil {
		log.Fatalln("Cannot open server-config file. Maybe create a server_config file?", err)
	}
	config = Configuration{}
	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		log.Fatalln("Cannot get configuration from file", err)
	}
}
func init() {
	loadConfig()
	file, err := os.OpenFile("chatter.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file", err)
	}
	logger = log.New(file, "INFO ", log.Ldate|log.Ltime|log.Lshortfile)
}

// Checks if the user is logged in and has a session, if not err is not nil
func IsLoggedIn(res http.ResponseWriter, req *http.Request) bool {
	cookie, err := req.Cookie("_cookie")
	if err != nil {
		return false
	}
	if ok, _ := data.CheckSessionValidity(cookie.Value); !ok {
		return false
	}
	return true
}

func GetUserIfLoggedIn(res http.ResponseWriter, req *http.Request) (logged_in bool, user data.User) {
	cookie, err := req.Cookie("_cookie")
	if err != nil {
		logged_in = false
		return
	}
	logged_in, user = data.GetUserBySessionUuid(cookie.Value)
	return
}

// for logging
func Info(args ...interface{}) {
	logger.SetPrefix("INFO ")
	logger.Println(args...)
}

func Danger(args ...interface{}) {
	logger.SetPrefix("ERROR ")
	logger.Println(args...)
}

func Warning(args ...interface{}) {
	logger.SetPrefix("WARNING ")
	logger.Println(args...)
}

// version
func Version() string {
	return config.Version
}

func ShowError(res http.ResponseWriter, req *http.Request, msg string) {
	http.Redirect(res, req, fmt.Sprintf("/err?msg=%v", msg), http.StatusFound)
}

func parseTemplateFiles(filenames ...string) (t *template.Template) {
	t = template.New("layout")
	t = template.Must(t.ParseFiles(Files("templates/%s.html", filenames...)...))
	return
}

func Files(format string, filenames ...string) (files []string) {
	for _, file := range filenames {
		files = append(files, fmt.Sprintf(format, file))
	}
	return
}

func GenerateHTML(res http.ResponseWriter, data interface{}, file_names ...string) {
	file_names = append(file_names, "common_head", "common_scripts")
	templates := parseTemplateFiles(file_names...)
	templates.Execute(res, data)
}

// func ListFilesInDirectory(dirPath string) ([]string, error) {
// 	// Open the directory
// 	dir, err := os.Open(dirPath)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer dir.Close()

// 	// Read the directory entries
// 	entries, err := dir.Readdir(0)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Initialize a slice to store filenames
// 	var filenames []string

// 	// Iterate through the directory entries
// 	for _, entry := range entries {
// 		if entry.Mode().IsRegular() {
// 			filenames = append(filenames, fmt.Sprintf("%s/%s", dirPath, entry.Name()))
// 		}
// 	}

// 	return filenames, nil
// }
