package main

import "fmt"

// struct: like a structure of things
type person struct {
	fname string
	lname string
}


// I can include other stuct types
type secretAgent struct {
	person
	lToKill bool
}

//Functions, you can attach "types" to them (what is it expecting to come in as)
/// Here you're declaring "p" as a person type inside this fucntion
func (p person) speak() {
	fmt.Println(p.fname, "says hello")
}


// since this has the same "fucntion name" it's part of the same struct
func (sa secretAgent) speak() {
	fmt.Println(sa.fname, "says Shaken not stirred")
}


// Interfaces can take values of different types, basically you can be more of one type
// Above; secretAgent and person BOTH implement a "speak" function; therefore they are "human"
type human interface {
	speak()
}

// this function expects and argument of "human" ...which, by extention, can be either a person or a secretAgent
func saySomething(h human) {
	h.speak()
}

func main() {
	// []THING means "slice"...like an array
	//		eg: []string, []int, []float

	x := []int{1, 2, 3}
	//fmt.Println(x)
	// For loop is saying "for index" (i.e. 0..n) "assign num from the range of x"
	for i, num := range x {
		fmt.Println(i, num)
	}

	// If you don't care about the index number, just use _ and golang will ignore it
	for _, num := range x {
		fmt.Println(num)
	}

	/// map = key value pairs

	//m := map[string]int{
	//	"Christian": 26,
	//	"Mark": 56,
	//	"Joe": 77,
	//}


	/// STruct example from above
	//fmt.Println(m["Christian"])
	//var you person

	//you.fname = "Christian"
	//you.lname = "Hernandez"

	//you := person{
	//	"Christian",
	//	"Hernandez",
	//}

	//fmt.Println(you)

	//use of a function with the "person" type
	//you.speak()

	// this is expecting a "person" struct and a bool if they have an lToKill
	sa1 := secretAgent {
		person {
			"James",
			"Bond",
		},
		true,
	}

	//sa1.speak()
	// Since sa1 is a "persoN" it can access the other functions
	//sa1.person.speak()


	// interface above I can saySomething and just pass it the secretAgent I created
	saySomething(sa1)
}
