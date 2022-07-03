package main

import "fmt"

func main() {

	var a = 4
	var b = 5.2

	a = int(b)
	fmt.Println(a, b)

	/* var x int
	x = "5"
	println(x)
	*/

	var (
		value int
		price float64
		name  string
		done  bool
	)

	println(value, price, name, done)

}
