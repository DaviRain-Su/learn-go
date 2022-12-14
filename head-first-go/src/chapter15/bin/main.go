package main

import (
	"fmt"
	"log"
	"net/http"
)

func sayHi() {
	fmt.Println("Hi")
}

func main() {
	var myFunction func()
	myFunction = sayHi
	myFunction()
	http.HandleFunc("/hello", viewHandler)
	http.HandleFunc("/salut", frenchHandler)
	http.HandleFunc("/namaste", hindHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}

func viewHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Hello, web!")
}

func frenchHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Salut, web!")
}

func hindHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Nameste, web!")
}

func write(writer http.ResponseWriter, message string) {
	_, err := writer.Write([]byte(message))
	if err != nil {
		log.Fatal(err)
	}
}
