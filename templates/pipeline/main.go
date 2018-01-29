package main

import (
	"log"
	"strings"
	"os"
	"text/template"
	"math"
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


// Parse the tempalte file and check for errors (.Must)
func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

func double(x int) int {
	return x + x
}

func square(x int) float64 {
	return math.Pow(float64(x), 2)
}

func sqRoot(x float64) float64 {
	return math.Sqrt(x)
}

var fm = template.FuncMap{
	"fdbl": double,
	"fsq": square,
	"fsqrt": sqRoot,
}
//
func main() {

	// Execute the template and print to standard out, replacing occcurances of "{{.}}" in "tpl.gohtml" with the values
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", 3)

	// Print any errors if there were any
	if err != nil {
		log.Fatalln(err)
	}
}
