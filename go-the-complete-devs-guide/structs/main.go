package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	name string
	age  int
	contactInfo
	//contact contactInfo
}

func main() {
	// alex := person{"Alex", 23}
	// alex := person{name: "Alex", age: 23}
	// fmt.Println(alex)

	// var alex person
	// alex.name = "Alexander"
	// alex.age = 24
	// fmt.Printf("%+v\n", alex)

	jim := person{
		name: "Jim",
		age:  52,
		contactInfo: contactInfo{ //contact: contactInfo{
			email:   "jim@gmail.net",
			zipCode: 12345,
		},
	}

	// jim.print()
	// jimPtr := &jim
	// jimPtr.updateName("Jimmy")
	// jim.print()

	jim.print()
	jim.updateName("Jimmy")
	jim.print()

	mySlice := []string{"hi", "there", "how", "are", "you"}
	fmt.Println(mySlice)
	updateSlice(mySlice)
	fmt.Println(mySlice)

	name := "bill"
	updateValue(name)
	fmt.Println(name)

}

func updateValue(n string) {
	n = "Alex"
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (pointerToPerson *person) updateName(newName string) {
	(*pointerToPerson).name = newName
}

func updateSlice(s []string) {
	s[0] = "bye"
}
