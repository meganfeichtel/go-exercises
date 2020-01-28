// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how polymorphic behavior with interfaces.
package main

import "fmt"

// reader is an interface that defines the act of reading data.
type reader interface {
	read(b []byte) (int, error)
	// read() ([]byte, error) {
	// 	b:= make([]byte, 32*1024)
	// 	return b
	//}
}

// file defines a system file.
type file struct {
	name string
}

// read implements the reader interface for a file.
func (file) read(b []byte) (int, error) {
	s := "<rss><channel><title>Going Go Programming</title></channel></rss>"
	copy(b, s)
	return len(s), nil
}

// pipe defines a named pipe network connection.
type pipe struct {
	name string
}

// read implements the reader interface for a network connection.
func (pipe) read(b []byte) (int, error) {
	s := `{name: "bill", title: "developer"}`
	copy(b, s)
	return len(s), nil
}

func interfaces1() {

	// Create two values one of type file and one of type pipe.
	f := file{"data.json"}
	p := pipe{"cfg_service"}

	// Call the retrieve function for each concrete type.
	retrieve(f)
	retrieve(p)
}

// retrieve can read any device and process the data.
func retrieve(r reader) error {
	data := make([]byte, 100)

	len, err := r.read(data)
	if err != nil {
		return err
	}

	fmt.Println(string(data[:len]))
	return nil
}

// notifier is an interface that defines notification
// type behavior.
type notifier interface {
	notify()
}

// user3 defines a user in the program.
type user3 struct {
	name  string
	email string
}

// notify implements the notifier interface with a pointer receiver.
func (u *user3) notify() {
	fmt.Printf("Sending user3 Email To %s<%s>\n",
		u.name,
		u.email)
}

func interfaces2() {

	// Create a value of type user3 and send a notification.
	u := user3{"Bill", "bill@email.com"}

	// Values of type user3 do not implement the interface because pointer
	// receivers don't belong to the method set of a value.

	sendNotification(&u) //this is the fix

	//sendNotification(u)
	// ERROR ./example1.go:36: cannot use u (type user3) as type notifier in argument to sendNotification:
	//   user3 does not implement notifier (notify method has pointer receiver)
}

// sendNotification accepts values that implement the notifier
// interface and sends notifications.
func sendNotification(n notifier) {
	n.notify()
}

type duration int

// notify implements the notifier interface.
func (d *duration) notify() {
	fmt.Println("Sending Notification in", *d)
}

func interfaces3() {
	println(duration(42))

	// duration(42).notify()
	// ERROR ./example3.go:18: cannot call pointer method on duration(42)
	// ./example3.go:18: cannot take the address of duration(42)
}
