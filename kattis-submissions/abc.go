package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

//ABC Problem
//Link: https://open.kattis.com/problems/abc
//Difficulty Level: 1.7
//Not submitted
func abc(ints string, lets string) string {

	var istr = strings.Split(ints, " ")
	var i = []int{}

	for _, x := range istr {
		y, err := strconv.Atoi(x)
		if err != nil {
			panic(err)
		}
		i = append(i, y)
	}
	sort.Ints(i)

	var l = strings.Split(lets, "")
	var r [3]int

	for x := 0; x < 3; x++ {
		if l[x] == "A" {
			r[x] = i[0]
		}
		if l[x] == "B" {
			r[x] = i[1]
		}
		if l[x] == "C" {
			r[x] = i[2]
		}
	}

	return fmt.Sprintf("%d %d %d", r[0], r[1], r[2])

}

func main() {
	var ints, lets string
	var res string

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		ints = scanner.Text()
	}
	fmt.Scanf("%s", &lets)

	res = abc(ints, lets)

	// fmt.Println(abc("1 2 3", "BCA")) //3 1 2
	// fmt.Println(abc("1 2 3", "ABC")) //1 2 3

	fmt.Printf("%s\n", res)
}
