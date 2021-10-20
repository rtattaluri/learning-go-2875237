package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("A simple calculator")
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter a number1:")
	stg1, _ := reader.ReadString('\n')
	fmt.Printf("Enter a number1:")
	stg2, _ := reader.ReadString('\n')

	number1, err := strconv.ParseFloat(strings.TrimSpace(stg1), 64)
	number2, err2 := strconv.ParseFloat(strings.TrimSpace(stg2), 64)

	if err != nil || err2 != nil {
		panic("unable to parse number, please pass a valid format")
	}

	number3 := number1 + number2
	fmt.Println("Total : ", number3)

}
