package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("------------------------------")

	i := 1
	for i < 10 {
		fmt.Println(i)
		i++
	}
}
