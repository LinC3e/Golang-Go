package main

import "fmt"

func main() {
	result1 := plus(1, 2)
	fmt.Println(result1)

	result2 := plusPlus(1, 2, 3)
	fmt.Println(result2)
}

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a int, b int, c int) int {
	return a + b + c
}
