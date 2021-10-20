package main

import (
	"fmt"
)

func main() {
	anInt := 42
	p := &anInt
	fmt.Println("pointer ", *p)
	fmt.Println("value ", anInt)

	anInt = 21
	fmt.Println("reassignment pointer ", *p)
	fmt.Println("reassignment value ", anInt)

	aFloat := 4.2
	p2 := &aFloat
	fmt.Println("Float pointer ", *p2)
	fmt.Println("Float value ", aFloat)

	*p2 = *p2 / 2
	fmt.Println("Float pointer ", *p2)
	fmt.Println("Float value ", aFloat)

}
