package main

import (
	"os"
	"text/template"
	"log"
	)
//

// create "tpl" that is a pointer to a template
var tpl *template.Template

// Parse the tempalte file and check for errors (.Must)
func init() {
	tpl = template.Must(template.ParseFiles("templates/tpl.gohtml"))
}

//
func main() {
	// Execute the template and print to standard out, replacing occcurances of "{{.}}" in "tpl.gohtml" with a value (in this case a string)
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", `Release self-focus; embrace other-focus.`)

	// Print any errors if there were any
	if err != nil {
		log.Fatalln(err)
	}
}
