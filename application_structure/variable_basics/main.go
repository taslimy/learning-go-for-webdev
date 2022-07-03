package main

import "fmt"

func main() {
	// To write variable you need a keyword,  type, indentifer
	// I do not need to include the type as go will automatically know.
	var age int = 27
	fmt.Println("Age:", age)

	var name string = "Tas"
	// fmt.Println("My name is:", name)
	_ = name // If i get compile error for unused variables.

	// short declartion operator for writing a variable.
	name1 := "Tas"
	fmt.Println(name1)

	myAge := "old xd"
	fmt.Println(myAge)
}
