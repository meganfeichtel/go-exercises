// Sample program to show how what we are doing is NOT embedding
// a type but just using a type as a field.
package main

import "fmt"

// user defines a user in the program.
type user4 struct {
	name  string
	email string
}

// notify implements a method notifies users
// of different events.
func (u *user4) notify() {
	fmt.Printf("Sending user email To %s<%s>\n",
		u.name,
		u.email)
}

// admin represents an admin user with privileges.
type admin struct {
	user4 // Embedded Type
	level string
}

func embedding1() {

	// Create an admin user.
	ad := admin{
		user4: user4{
			name:  "john smith",
			email: "john@yahoo.com",
		},
		level: "super",
	}

	// We can access the inner type's method directly.
	ad.user4.notify()

	// The inner type's method is promoted.
	ad.notify()
}

//hallmark example:
// type user5 struct {
// 	name  string
// 	email string
// 	alias string
// }

// //this is what it should be
// func sendEmail(name string, email string, alias string) {
// 	email.send(name, alias)
// }

// //bad
// func sendEmail(u *user5) {
// 	email.send(u.name, u.email, u.alias)
// }

// //bad bad bad
// func (u *user4) sendEmail() {
// 	email.send(u.name, u.email, u.alias)
// }
