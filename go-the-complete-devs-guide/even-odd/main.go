package main

import (
	"fmt"
)

func main() {
	list := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}

	for _, e := range list {
		eo := evenOrOdd(e)
		fmt.Println(e, "is "+eo) //(fmt.Sprintf("%v is "+eo, e))

	}
}

func evenOrOdd(i int) string {
	if i%2 == 0 {
		return "even"
	} else {
		return "odd"
	}
}
