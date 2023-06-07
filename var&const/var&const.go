package main

import "fmt"

func main() {

	// const
	const pe float64 = 3.1415926

	// multiple variables like this ...
	var (
		varA = 1
		varB = 2
	)

	fmt.Println(varA, varB)

	var Name string = "Joseph"
	fmt.Println(Name)

	// Get string lenght
	fmt.Println(len(Name))
}
