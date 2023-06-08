package main

import "fmt"

func main() {

	Student := make(map[string]int)

	Student["Jhon"] = 15
	Student["Ann"] = 14
	Student["Sarah"] = 17
	Student["Max"] = 12
	Student["Richard"] = 15

	fmt.Println(Student)        // entire map
	fmt.Println(Student["Ann"]) // single
	fmt.Println(len(Student))   // map lenght
}
