package main

import "fmt"

func a() {
	number := 3

	if number > 0 {
		switch number {
		case 1:
			fmt.Println("One")
		case 2:
			fmt.Println("Two")
		case 3:
			fmt.Println("Three")
		default:
			fmt.Println("Its more than Three.")
		}
	} else {
		fmt.Println("Its zero or negative")
	}

}
