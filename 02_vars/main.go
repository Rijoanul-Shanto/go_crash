package main

import "fmt"

func main() {
	var name = "Robert Junior"
	var age = 26

	const PI = 3.1415

	name += " #"
	age = 25

	fmt.Printf("%s is a: '%T'\n", name, name)
	fmt.Printf("%d is a: '%T'\n", age, age)
	fmt.Printf("%f is a: '%T'\n", PI, PI)
}
