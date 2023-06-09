package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {
	fmt.Println(person{"Ann", 21})
	fmt.Println(newPerson("Joseph"))

	fmt.Println("--------------")

	other := person{"Mathew", 15}
	fmt.Println(other)
	fmt.Println(other.age)
	fmt.Println(other.name)
}
