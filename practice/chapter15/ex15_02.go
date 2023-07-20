package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var templateStr = `
<html>
	<h1>Customer {{.ID}}</h1>
	{{if .ID}}
		<p>Details</p>
		<ul>
			{{if .Name}}<li>Name: {{.Name}}</li>{{end}}
			{{if .Surname}}<li>Surname: {{.Surname}}</li>{{end}}
			{{if .Age}}<li>Age: {{.Age}}</li>{{end}}
		</ul>
	{{else}}
		<p>Data not availabe</p>
	{{end}}
</html>
`

type Customer struct {
	ID      int
	Name    string
	Surname string
	Age     int
}

func Hello(w http.ResponseWriter, r *http.Request) {
	v1 := r.URL.Query()
	customer := Customer{}

	id, ok := v1["id"]
	if ok {
		customer.ID, _ = strconv.Atoi(id[0])
	}

	name, ok := v1["name"]
	if ok {
		customer.Name = name[0]
	}

	surname, ok := v1["surname"]
	if ok {
		customer.Surname = surname[0]
	}

	age, ok := v1["age"]
	if ok {
		customer.Age, _ = strconv.Atoi(age[0])
	}

	tmpl, _ := template.New("Exercise").Parse(templateStr)
	tmpl.Execute(w, customer)
}

func main() {
	http.HandleFunc("/", Hello)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
