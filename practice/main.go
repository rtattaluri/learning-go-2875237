package main

import (
	"fmt"
	"sort"
)

func main() {
	var colors = []string{"Red", "Green", "Blue"}
	fmt.Println(colors)
	fmt.Println(len(colors))

	colors = append(colors, "Yellow")
	fmt.Println(colors)
	fmt.Println(len(colors))

	colors = append(colors[1:len(colors)])
	fmt.Println(colors)
	fmt.Println(len(colors))
	colors = append(colors[:len(colors)-1])
	fmt.Println(colors)
	fmt.Println(len(colors))

	numbers := make([]int, 5)
	fmt.Println(numbers)
	fmt.Println(len(numbers))
	numbers[0] = 101
	numbers[1] = 72
	numbers[2] = 32
	numbers[3] = 96
	numbers[4] = 54
	fmt.Println(numbers)
	fmt.Println(len(numbers))

	numbers = append(numbers, 110)
	fmt.Println(numbers)
	fmt.Println(len(numbers))

	sort.Ints(numbers)
	fmt.Println(numbers)

}
