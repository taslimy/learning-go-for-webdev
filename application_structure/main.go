package main

import "fmt"

const secondsInHour = 3600

func main() {
	fmt.Println("Hello Go World!!")
	distance := 60.8
	// %f is called a verb and tells Printf how to format the output.
	fmt.Printf("the distance in miles is %f \n", distance*0.62137)
}
