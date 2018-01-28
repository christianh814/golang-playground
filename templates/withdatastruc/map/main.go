package main

import (
	"log"
	"os"
	"text/template"
)

//

// create "tpl" that is a pointer to a template
var tpl *template.Template

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
	//create map
	sages := map[string]string {
		"India": "Gandhi",
		"America": "MLK",
		"Meditate": "Buddha",
		"Love": "Jesus",
		"Prophet": "Muhammad",
	}
	// Execute the template and print to standard out, replacing occcurances of "{{.}}" in "tpl.gohtml" with the values
	err := tpl.Execute(os.Stdout, sages)

	// Print any errors if there were any
	if err != nil {
		log.Fatalln(err)
	}
}
