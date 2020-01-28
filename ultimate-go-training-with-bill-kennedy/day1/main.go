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

	println("\nSlices exercise 1:")
	slices1()

	println("\nSlices exercise 2:")
	slices2()

	println("\nSlices exercise 3:")
	slices3()

	println("\nMethods exercise 1:")
	methods1()

	println("\nMethods exercise 2:")
	methods2()

	println("\nInterfaces exercise 1:")
	interfaces1()

	println("\nInterfaces exercise 4:")
	interfaces4()

}
