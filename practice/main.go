package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	input, _ := reader.ReadString('\n')
	fmt.Println("You entered:", input)

	reader2 := bufio.NewReader(os.Stdin)
	fmt.Print("Enter number: ")
	input2, _ := reader2.ReadString('\n')
	aFloat, err := strconv.ParseFloat(strings.TrimSpace(input2), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Print("Entered Number is : ")
		fmt.Println(aFloat)
		fmt.Printf("Entered Number format is : %T\n", aFloat)
	}

	anInt, err := strconv.ParseFloat(strings.TrimSpace(input2), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Print("Entered Number is : ")
		fmt.Println(anInt)
		fmt.Printf("Entered Number format is : %T\n", anInt)
	}
	var anInt2 string = strconv.FormatFloat(anInt, 'f', 5, 64)
	fmt.Print("Integer is : ")
	fmt.Println(anInt2)

}
