// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to declare methods and how the Go
// compiler supports them.
package main

import (
	"fmt"
)

// user2 defines a user2 in the program.
type user2 struct {
	name  string
	email string
}

// notify implements a method with a value receiver.
func (u user2) notify() {
	fmt.Printf("Sending user2 Email To %s<%s>\n",
		u.name,
		u.email)
}

// changeEmail implements a method with a pointer receiver.
func (u *user2) changeEmail(email string) {
	u.email = email
}

func methods1() {

	// Values of type user2 can be used to call methods
	// declared with both value and pointer receivers.
	bill := user2{"Bill", "bill@email.com"}
	bill.changeEmail("bill@hotmail.com")
	bill.notify()

	// Pointers of type user2 can also be used to call methods
	// declared with both value and pointer receiver.
	joan := &user2{"Joan", "joan@email.com"}
	joan.changeEmail("joan@hotmail.com")
	joan.notify()

	// Create a slice of user2 values with two user2s.
	users := []user2{
		{"ed", "ed@email.com"},
		{"erick", "erick@email.com"},
	}

	// Iterate over the slice of user2s switching
	// semantics. Not Good!
	//for i := range users { //fix: pointer semantics
	for _, u := range users {
		u.changeEmail("it@wontmatter.com")
	}

	// Exception example: Using pointer semantics
	// for a collectoin of strings.
	keys := make([]string, 10)
	for i := range keys {
		keys[i] = func() string { return "key-gen" }()
	}
}

// data is a struct to bind methods to.
type data struct {
	name string
	age  int
}

// displayName provides a pretty print view of the name.
func (d data) displayName() {
	fmt.Println("My Name Is", d.name)
}

// setAge sets the age and displays the value.
func (d *data) setAge(age int) {
	d.age = age
	fmt.Println(d.name, "Is Age", d.age)
}

func methods2() {

	// Declare a variable of type data.
	d := data{
		name: "Bill",
	}

	fmt.Println("Proper Calls to Methods:")

	// How we actually call methods in Go.
	d.displayName()
	d.setAge(45)

	fmt.Println("\nWhat the Compiler is Doing:")

	// This is what Go is doing underneath.
	data.displayName(d)
	(*data).setAge(&d, 45)

	// =========================================================================

	fmt.Println("\nCall Value Receiver Methods with Variable:")

	// Declare a function variable for the method bound to the d variable.
	// The function variable will get its own copy of d because the method
	// is using a value receiver.
	f1 := d.displayName

	// Call the method via the variable.
	f1()

	// Change the value of d.
	d.name = "Joan"

	// Call the method via the variable. We don't see the change.
	f1()

	// =========================================================================

	fmt.Println("\nCall Pointer Receiver Method with Variable:")

	// Declare a function variable for the method bound to the d variable.
	// The function variable will get the address of d because the method
	// is using a pointer receiver.
	f2 := d.setAge

	// Call the method via the variable.
	f2(45)

	// Change the value of d.
	d.name = "Sammy"

	// Call the method via the variable. We see the change.
	f2(45)
}
