package main

import "fmt"

func main() {
	var name = "Robert Junior"
	var age = 26

	fmt.Printf("%s is a: '%T'\n", name, name)
	fmt.Printf("%d is a: '%T'\n", age, age)
}
