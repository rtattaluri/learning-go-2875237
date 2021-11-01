package main

import (
	"fmt"
)

func main() {
	colors := [3]string{"Red", "Blue", "Green"}
	fmt.Println("Colors :", colors)

	var num [3]int

	num[1] = 42
	num[2] = 38
	num[0] = 0
	fmt.Println("Numbers: ", num)

	fmt.Println("length of colors: ", len(colors))
	fmt.Println("length of numbers: ", len(num))

}
