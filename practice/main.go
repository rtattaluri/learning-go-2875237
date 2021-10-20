package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println("Math")
	i1, i2, i3 := 20, 30, 40
	sumInt := i1 + i2 + i3

	f1, f2, f3 := 45.6, 60.3, 72.8
	sumFloat := f1 + f2 + f3

	fmt.Println("Int Sum : ", sumInt)
	fmt.Println("Float Sum : ", sumFloat)
	fmt.Println("Float Sum formatted:", math.Round(sumFloat*100)/100)

	radius := 4.0
	circumference := radius * 2 * math.Pi

	fmt.Println("Circumference : ", circumference)
	fmt.Printf("Circumference rounded : %.2f", circumference)

}
