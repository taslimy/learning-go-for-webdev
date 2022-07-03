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

	// Declaring multiple variables.
	car, cost := "BMW", 60000
	fmt.Println(car, cost)
	car, year := "Toyota", 2018
	_ = year

	opened := false
	opened, file := true, "a.txt"

	_, _ = opened, file

	// Clearer way to write multiple variables - better readability
	var (
		salary    float64
		firstName string
		gender    bool
	)

	fmt.Println(salary, firstName, gender)

	// multiple same type
	var a, b, c int
	fmt.Println(a, b, c)

	// multiple assignment
	var d, e int
	d, e = 5, 8

	d, e = e, d // Swapping variables
	println(d, e)

	sum := 6 + 2.3
	fmt.Println(sum)
}
