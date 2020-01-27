package main

func main() {
	println("Value-semantics (own cope of data) version:")
	valueof()

	println("Data-semantics (sharing data) version:")
	valueof2()

	println("Escape analysis:")
	escapeAnalysis()
}
