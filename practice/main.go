package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	inputString := bufio.NewReader(os.Stdin)
	fmt.Print("Please input a string: ")
	input, _ := inputString.ReadString('\n')
	fmt.Println("Entered String is : " + input)
	fmt.Printf("Format of Entered String is : %T\n", input)
	//fmt.Println("End")
}
