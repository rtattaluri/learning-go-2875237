package main

import (
	"fmt"
)

func main() {
	colors := []string{"Red", "Green", "Blue"}
	fmt.Println(colors)

	for i := 0; i < len(colors); i++ {
		fmt.Println(colors[i])
	}

	for i := range colors {
		fmt.Println(colors[i])
	}

	for _, color := range colors {
		fmt.Println(color)
	}

	x := 0
	for x < 25 {
		x += 1
		fmt.Println("Sum is ", x)
		if x > 20 {
			break
		}
		fmt.Println("final Sum is ", x)
	}

	fmt.Println("Program finished")

}
