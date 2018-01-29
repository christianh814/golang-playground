package main

import (
	"log"
	"strings"
	"os"
	"text/template"
	"time"
)

//

// create "tpl" that is a pointer to a template
var tpl *template.Template

type sage struct {
	Name string
	Motto string
}

type car struct {
	Manu string
	Model string
	Doors int
}

type items struct {
	Wisdom []sage
	Transport []car
}


//Create a FuncMap to register functions
// uc is what it will be called in the template
// uc is the toUpper function from pkg strings
// ft is the func I declared 
// ft slices a string

// Parse the tempalte file and check for errors (.Must)
func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

func monthDayYear(t time.Time) string {
	return t.Format("01-02-2006")
}

var fm = template.FuncMap{
	"fdateMDY": monthDayYear,
}
//
func main() {

	// Execute the template and print to standard out, replacing occcurances of "{{.}}" in "tpl.gohtml" with the values
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", time.Now())

	// Print any errors if there were any
	if err != nil {
		log.Fatalln(err)
	}
}
