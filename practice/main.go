package main

import (
	"fmt"
)

func main() {
	a := 43
	var result string
	if a < 0 {
		result = "Less than zero"
	} else if a == 0 {
		result = "Equal to zero"
	} else {
		result = "Greater than zero"
	}
	fmt.Println(result)

	if a := -1; a < 0 {
		result = "Less than zero"
	} else if a == 0 {
		result = "Equal to zero"
	} else {
		result = "Greater than zero"
	}
	fmt.Println(result)
}
