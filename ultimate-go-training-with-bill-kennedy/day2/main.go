package main

func main() {
	println("\nConcurrency 1:") //go scheduler
	concurrency1()

	println("\nConcurrency 2:") //go scheduler
	concurrency2()

	println("\nConcurrency 3:") //in-parallel
	concurrency3()
}
