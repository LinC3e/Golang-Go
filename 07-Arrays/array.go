package main

import "fmt"

func main() {

	var Numbers [5]int

	Numbers[0] = 10
	Numbers[1] = 11
	Numbers[2] = 12
	Numbers[3] = 13
	Numbers[4] = 14

	fmt.Println("----------------------")
	fmt.Println(Numbers[2]) // print 12
	fmt.Println("----------------------")
	// Short form

	Nums := [5]int{10, 11, 12, 13, 14}

	fmt.Println(Nums[3]) // print 13

	// iterate array
	fmt.Println("----------------------")
	for _, v := range Nums {
		fmt.Println(v)
	}
}
