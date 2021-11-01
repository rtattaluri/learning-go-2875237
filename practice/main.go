package main

import (
	"fmt"
)

func main() {
	poodle := Dog{"poodle", 15}
	fmt.Println("Dog : ", poodle)

	fmt.Printf("%+v \n", poodle)

	poodle.Weight = 10

	fmt.Printf("Breed: %v\n Weight:%v\n", poodle.Breed, poodle.Weight)

}

type Dog struct {
	Breed  string
	Weight int
}
