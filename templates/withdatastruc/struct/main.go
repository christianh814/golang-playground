package main

import (
	"log"
	"os"
	"text/template"
)

//

// create "tpl" that is a pointer to a template
var tpl *template.Template
type sage struct {
	Name string
	Motto string
}

// Parse the tempalte file and check for errors (.Must)
func init() {
	tf := os.Args[1]
	if _, err := os.Stat(tf); os.IsNotExist(err) {
		log.Fatalln(err)
	}
	tpl = template.Must(template.ParseFiles(tf))
}

//
func main() {
	//create struct
	buddah := sage {
		Name: "Buddah",
		Motto: "The b of no b",
	}
	// Execute the template and print to standard out, replacing occcurances of "{{.}}" in "tpl.gohtml" with the values
	err := tpl.Execute(os.Stdout, buddah)

	// Print any errors if there were any
	if err != nil {
		log.Fatalln(err)
	}
}
