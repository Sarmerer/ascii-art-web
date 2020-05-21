package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"

	student "./static"
)

var indexTpl *template.Template
var tpl404 *template.Template

type Data struct {
	Output    string
	ErrorCode int
	Error     string
}

func init() {
	indexTpl = template.Must(template.ParseGlob("templates/index/*.html"))
	tpl404 = template.Must(template.ParseGlob("templates/404/*.html"))

	t := time.Now()
	fmt.Println(t.Format("3:4:5pm"), "Init complete.")
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", index)

	t := time.Now()
	fmt.Println(t.Format("3:4:5pm"), "Starting server, go to localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		data404 := Data{
			ErrorCode: 404,
			Error:     "404 Page not found",
		}
		tpl404.ExecuteTemplate(w, "404.html", data404)
		return
	}

	switch r.Method {
	case "GET":
		info, err := student.Draw("1. Type\n2. Select font\n3. Submit\n4. Enjoy :)", "standard")
		dataI := Data{
			Output:    info,
			Error:     info,
			ErrorCode: err,
		}
		if err != 1 {
			if errorHandle(w, r, dataI) {
				return
			}
		}
		indexTpl.ExecuteTemplate(w, "index.html", dataI)
		break
	case "POST":
		output, err := student.Draw(r.FormValue("text"), r.FormValue("font"))
		data := Data{
			ErrorCode: err,
			Output:    output,
			Error:     output,
		}
		if errorHandle(w, r, data) {
			return
		}
		indexTpl.Execute(w, data)
		break
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
		break
	}
}

func errorHandle(w http.ResponseWriter, r *http.Request, data Data) bool {
	if data.ErrorCode != 1 {
		t := time.Now()
		fmt.Println(t.Format("3:4:5pm"), "Error while printing. Stopped the operation.")
		tpl404.ExecuteTemplate(w, "404.html", data)
		return true
	}
	return false
}
