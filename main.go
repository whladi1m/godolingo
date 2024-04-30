package main

import (
	//"fmt"
	"net/http"
	"html/template"
)

type User struct {
	Name string
	Id uint16
	Tasks_completed int16
}


func homepage(w http.ResponseWriter, r *http.Request) {
	bob := User{"User", 16, 5}
	bob.Name = "Srhhr"
	//fmt.Fprintf(w, user.name)

	tmpl, _ := template.ParseFiles("home.html")
	tmpl.Execute(w, bob)
}

func pages_handler() {
	http.HandleFunc("/", homepage)
	http.ListenAndServe(":8080", nil)
}

func main() {
	pages_handler()
}
