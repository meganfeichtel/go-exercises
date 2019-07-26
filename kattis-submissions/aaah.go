package main

import (
	"fmt"
)

//Aaah! Problem
//Link: https://open.kattis.com/problems/aaah
//Difficulty Level: 1.6
//ACCEPTED
func aaah() {
	var a, b string
	var res string
	fmt.Scanf("%s", &a)
	fmt.Scanf("%s", &b)

	if len(a) >= len(b) {
		res = "go"
	} else {
		res = "no"
	}

	fmt.Printf("%s\n", res)
}
