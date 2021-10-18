package main

import (
	"fmt"
)

const aConstant string = "128"

func main() {

	var aString string = "This is a string"
	fmt.Println("aString " + aString)
	fmt.Printf("aString type is %T\n", aString)

	var anInt int = 32
	fmt.Println(anInt)
	fmt.Printf("anInt type is %T\n", anInt)

	var anInt2 int
	fmt.Println(anInt2)
	fmt.Printf("anInt2 type is %T\n", anInt2)

	var anotherString = "This is another string"
	fmt.Println("anotherString " + anotherString)
	fmt.Printf("anotherString type is %T\n", anotherString)

	anotherString2 := "This is not an implicit string"
	fmt.Println("anotherString2 " + anotherString2)
	fmt.Printf("anotherString2 type is %T\n", anotherString2)

	fmt.Println(aConstant)
	fmt.Printf("aConstant type is %T\n", aConstant)

	//fmt.Println("Hello from Go!")
}
