package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/models", modelsHandler)
	http.HandleFunc("/manufacturers", manufacturerHandler)
	http.HandleFunc("/categories", categoryHandler)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer((http.Dir("static")))))

	fmt.Println("Starting Server.....")
	fmt.Println("http://localhost:5000")
	http.ListenAndServe(":5000", nil)

}
