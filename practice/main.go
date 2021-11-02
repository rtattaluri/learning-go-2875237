package main

import (
	"fmt"
)

func main() {
	fmt.Println("Functions:")
	doSomething()
	total := addValues(1, 3)
	fmt.Println("values are: ", total)

	total = addMultipleValues(2, 7, 10, 20)
	fmt.Println("values are: ", total)

	size := 0
	total, size = addMultipleValuesAndReturnSize(7, 10, 20)
	fmt.Println("values are: ", total, size)

	a := func() {
		fmt.Println("Inside anonymized function")
	}
	a()
}

func doSomething() {
	fmt.Println("Inside do something function")
}

func addValues(a int, b int) int {
	return a + b
}

func addMultipleValues(a ...int) int {
	total := 0
	for _, k := range a {
		total += k
	}
	return total
}

func addMultipleValuesAndReturnSize(a ...int) (int, int) {
	total := 0
	for i := range a {
		total += a[i]
	}
	return total, len(a)
}
