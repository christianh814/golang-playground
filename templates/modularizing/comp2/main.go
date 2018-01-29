package main

import (
	"log"
	"os"
	"text/template"
)

//
type course struct {
	Number, Name, Untis string
}

type semester struct {
	Term string
	Courses []course
}

type year struct {
	Fall, Spring, Summer semester
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main () {
	y := year {
		Fall: semester {
				Term: "Fall",
				Courses: []course{
					course{"WWW-01", "Intro"},
					course{"WWW-02", "Advanced"},
					course{"WWW-01", "Master"},
				},
				Term: "Sprint",
				Courses: []course{
					course{"YYY-01", "Intro yeah"},
					course{"YYY-02", "Advanced uh oh"},
					course{"YYY-01", "Master one"},
				},
			},
	}
	err := tpl.Execute(os.Stdout, y)
	if err != nil {
		log.Fatalln(err)
	}
}
