package main

import "fmt"

// Person is a struct
type person struct {
	first string
	last  string
	age   int
}

// Speak is a method or receiver function
func (p person) speak() {
	fmt.Println("I am", p.first, p.last, "and I am", p.age, "years old.")
}

type human interface {
	speak()
}

// main is the entry point for the program
func main() {

	p1 := person{
		first: "James",
		last:  "Bond",
		age:   42,
	}
	p2 := person{
		first: "Miss",
		last:  "Moneypenny",
		age:   61,
	}
	// fmt.Println(p1.first, p1.last, p1.age)	// James Bond true
	// fmt.Println(p2.first, p2.last, p2.age)	// Miss Moneypenny false
	fmt.Println(p1)
	fmt.Println(p2)

	p1.speak()
	p2.speak()

	var x human
	x = p1
	fmt.Printf("%T\n", x)

}
