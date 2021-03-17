package main

import "fmt"

func main() {
	var name = "Robert Junior"
	var age = 26
	var isHuman = true

	const PI = 3.1415

	name += " #"
	age = 25
	isHuman = false

	fmt.Printf("%s is a: '%T'\n", name, name)
	fmt.Printf("%d is a: '%T'\n", age, age)
	fmt.Printf("%f is a: '%T'\n", PI, PI)
	fmt.Printf("%t is a: '%T'\n", isHuman, isHuman)
}
