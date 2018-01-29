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
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}


//
func main() {
	xs := []string{"zero", "one", "two", "three", "four", "five",}

	// Execute the template and print to standard out, replacing occcurances of "{{.}}" in "tpl.gohtml" with the values
	err := tpl.Execute(os.Stdout, xs)

	// Print any errors if there were any
	if err != nil {
		log.Fatalln(err)
	}
}
