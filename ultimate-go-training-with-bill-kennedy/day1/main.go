package main

func main() {
	println("\nValue-semantics (own cope of data) version:")
	valueof()

	println("\nData-semantics (sharing data) version:")
	valueof2()

	println("\nEscape analysis:")
	escapeAnalysis()

	println("\nArrays exercise 1:")
	arrays1()

	println("\nArrays exercise 2:")
	arrays2()
}
