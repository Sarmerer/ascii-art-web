package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"text/template"
	"time"

	student "./static"
)

var indexTpl *template.Template
var tpl404 *template.Template
var dataG Data

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
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", index)
	http.HandleFunc("/export", export)

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
		dataG = Data{
			ErrorCode: err,
			Output:    output,
			Error:     output,
		}
		if errorHandle(w, r, dataG) {
			return
		}
		indexTpl.Execute(w, dataG)
		//export(w, r)
		break
	default:
		data404 := Data{
			ErrorCode: 405,
			Error:     "You are not supposed to be here",
		}
		tpl404.ExecuteTemplate(w, "404.html", data404)
		return
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

func export(w http.ResponseWriter, r *http.Request) {
	format := r.FormValue("format")
	fileName := randomName() + format
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	l, err := f.WriteString(dataG.Output)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(l, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	w.Header().Set("Content-Disposition", "attachment; filename="+fileName)
	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
	http.ServeFile(w, r, fileName)
	os.Remove(fileName)
}

func randomName() string {
	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	length := 4
	b := make([]byte, length)
	var res string
	charset := "1234567890abcdefghijklmnopqrstyvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for i := 0; i < 4; i++ {
		for i := range b {
			b[i] = charset[seededRand.Intn(len(charset))]
		}
		if i > 0 && i < 4 {
			res += "-"
			res += string(b)
		} else {
			res += string(b)
		}
	}
	return res
}
