package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func readInputNumbers() float64 {
	reader := bufio.NewReader(os.Stdin)
	input1, _ := reader.ReadString('\n')
	float1, err := strconv.ParseFloat(strings.TrimSpace(input1), 64)
	if err != nil {
		fmt.Println(err)
		panic("Value 1 must be a number")
	}
	return float1
}

func readInputOperation() string {
	reader := bufio.NewReader(os.Stdin)
	input1, _ := reader.ReadString('\n')
	return strings.TrimSpace(input1)
}

func main() {

	fmt.Print("Value 1: ")
	firstInput := readInputNumbers()
	fmt.Print("Value 2: ")
	secondInput := readInputNumbers()
	fmt.Print("Enter the operation you would like to perform : + - * /")
	operation := readInputOperation()
	var output float64
	switch operation {
	case "+":
		output = sum(firstInput, secondInput)
	case "-":
		output = substraction(firstInput, secondInput)
	case "*":
		output = multiplication(firstInput, secondInput)
	case "/":
		output = Division(firstInput, secondInput)
	}

	output = math.Round(output*100) / 100
	fmt.Printf("Output : %v \n", output)

}

func sum(i, j float64) float64 {
	return i + j
}

func substraction(i, j float64) float64 {
	return i - j
}

func multiplication(i, j float64) float64 {

	return i * j
}

func Division(i, j float64) float64 {
	return i / j
}
