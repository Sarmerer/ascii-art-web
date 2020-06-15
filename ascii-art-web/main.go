package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
	"time"

	student "./pkg/student"
)

var indexTpl *template.Template
var tpl404 *template.Template

type Data struct {
	Output    string
	ErrorCode int
	Error     string
	ID        int
}

func init() {
	indexTpl = template.Must(template.ParseGlob("templates/index/*.html"))
	tpl404 = template.Must(template.ParseGlob("templates/404/*.html"))
	t := time.Now()
	fmt.Println(t.Format("3:4:5pm"), "Init complete.")
}

func main() {

	port := ":4241"

	router := http.NewServeMux()
	router.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	router.HandleFunc("/favicon.ico", faviconHandler)
	router.HandleFunc("/process", process)
	router.HandleFunc("/export", export)
	router.HandleFunc("/", index)

	t := time.Now()
	fmt.Println(t.Format("3:4:5pm"), "Starting server, go to localhost"+port)
	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatal(err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		callErrorPage(w, r, 404)
		return
	}
	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		indexTpl.ExecuteTemplate(w, "index.html", nil)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		callErrorPage(w, r, 405)
		return
	}
}

func process(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/process" {
		callErrorPage(w, r, 404)
		return
	}
	switch r.Method {
	case "POST":
		output, err := student.Draw(r.FormValue("text"), r.FormValue("font"))
		if err != 1 {
			callErrorPage(w, r, err)
		}
		b, err1 := json.Marshal(output)
		if err1 != nil {
			callErrorPage(w, r, 500)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		callErrorPage(w, r, 405)
		return
	}
}

func export(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/export" {
		callErrorPage(w, r, 404)
		return
	}
	switch r.Method {
	case "GET":
		output := r.FormValue("output")
		format := r.FormValue("format")
		fileName := r.FormValue("input")

		w.Header().Set("Content-Length", strconv.Itoa(len(output)))
		switch format {
		case ".txt":
			w.Header().Set("Content-Type", "text/plain")
			w.Header().Set("Content-Disposition", `attachment; filename="`+fileName+format+`"`)
		default:
			w.WriteHeader(http.StatusBadRequest)
			callErrorPage(w, r, 400)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(output))
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		callErrorPage(w, r, 405)
		return
	}
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/assets/favicon.ico")
}

func callErrorPage(w http.ResponseWriter, r *http.Request, errorCode int) {
	var errorMsg string

	switch errorCode {
	case 404:
		w.WriteHeader(http.StatusNotFound)
		errorMsg = "404 Page not found"
	case 405:
		w.WriteHeader(http.StatusMethodNotAllowed)
		errorMsg = "405 Wrong method"
	case 400:
		w.WriteHeader(http.StatusBadRequest)
		errorMsg = "400 Bad request"
	default:
		w.WriteHeader(http.StatusInternalServerError)
		errorMsg = "500 Internal error"
		errorCode = 500
	}

	data404 := Data{
		ErrorCode: errorCode,
		Error:     errorMsg,
	}
	tpl404.ExecuteTemplate(w, "404.html", data404)
	return
}
