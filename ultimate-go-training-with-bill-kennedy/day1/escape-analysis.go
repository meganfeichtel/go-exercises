// Sample program to teach the mechanics of escape analysis.
package main

// user represents a user in the system.
type user1 struct {
	name  string
	email string
}

// main is the entry point for the application.
func escapeAnalysis() {
	u1 := createUserV1()

	//once we're done with this function, the stackframe gets cleaned up
	//this is an integrity issue because we're now calling a function that creates a pointer and shares it back to main
	//now we have something (u2) pointing to nothing at the end of this function
	u2 := createUserV2()

	println("u1", &u1, "u2", u2)
	println("u1 name", u1.name, "u2 name", u2.name)
}

// createUserV1 creates a user value and passed
// a copy back to the caller.
//go:noinline
func createUserV1() user1 {

	//zero value construction: var u user

	u := user1{
		name:  "Bill",
		email: "bill@ardanlabs.com",
	}

	//share value down the call-stack
	println("V1", &u)

	//making a copy of the user in the go-routine
	return u
}

// createUserV2 creates a user value and shares
// the value with the caller.
//go:noinline
func createUserV2() *user1 {
	//using pointer semantics
	u := user1{
		name:  "Megan",
		email: "megan@megan.com",
	}

	//this time we're taking the address and share it up the call-stack
	println("V2", &u)

	return &u
}

/* go build -gcflags -m=2
./escape-analysis.go:36:16: createUserV1 &u does not escape
./escape-analysis.go:55:9: &u escapes to heap
./escape-analysis.go:55:9:      from ~r0 (return) at ./escape-analysis.go:55:2
./escape-analysis.go:47:2: moved to heap: u
./escape-analysis.go:53:16: createUserV2 &u does not escape
./escape-analysis.go:19:16: escapeAnalysis &u1 does not escape
*/
