package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/guestbook", viewHandler)
	http.HandleFunc("/salut", frenchHandler)
	http.HandleFunc("/namaste", hindHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}

func viewHandler(writer http.ResponseWriter, request *http.Request) {
	html, err := template.ParseFiles("view.html")
	check(err)
	err = html.Execute(writer, nil)
	check(err)
}

func frenchHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Salut, web!")
}

func hindHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Nameste, web!")
}

func write(writer http.ResponseWriter, message string) {
	_, err := writer.Write([]byte(message))
	check(err)
}
func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func executeTemplate(text string, data interface{}) {
	tmpl, err := template.New("test").Parse(text)
	check(err)
	err = tmpl.Execute(os.Stdout, data)
	check(err)
}
