package main

import "fmt"

func main() {
	var name = "Example exa"
	const pi float64 = 3.14123211
	learning := true
	x := 15

	fmt.Printf("%.3f \n", pi)     // 3 dig before .
	fmt.Printf("%T \n", name)     // type
	fmt.Printf("%t \n", learning) // boolean
	fmt.Printf("%d \n", x)        // ints
	fmt.Printf("%b \n", 10)       // binary
	fmt.Printf("%c \n", 97)       // ascii code
	fmt.Printf("%x \n", 15)       // hex
	fmt.Printf("%e \n", pi)       // science notation

}
