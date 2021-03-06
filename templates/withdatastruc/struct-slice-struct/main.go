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
	gandi := sage {
		Name: "Gandi",
		Motto: "Be the change",
	}
	mlk := sage {
		Name: "MLK",
		Motto: "I have a dream",
	}
	ford := car {
		Manu: "Ford",
		Model: "Tarus",
		Doors: 4,
	}
	chevy := car {
		Manu: "Chevy",
		Model: "Camero",
		Doors: 2,
	}

	// Now create a slice from that struct
	sages := []sage{buddah, gandi, mlk}
	//create another slice while you're at it
	cars := []car{ford, chevy}

	//Create data with the sclices
	data := items{
		Wisdom: sages,
		Transport: cars,
	}
	// Execute the template and print to standard out, replacing occcurances of "{{.}}" in "tpl.gohtml" with the values
	err := tpl.Execute(os.Stdout, data)

	// Print any errors if there were any
	if err != nil {
		log.Fatalln(err)
	}
}
