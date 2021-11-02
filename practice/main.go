package main

import (
	"fmt"
)

func main() {
	poodle := Dog{"Poodle", 10, "Wrf!"}
	fmt.Println(poodle)
	fmt.Printf("%+v\n", poodle)
	fmt.Printf("Breed: %v\nWeight: %v\n", poodle.Breed, poodle.Weight)
	poodle.sounds()
	poodle.soundsThreeTimes()
	fmt.Println(poodle)
}

// Dog is a struct
type Dog struct {
	Breed  string
	Weight int
	bark   string
}

func (d Dog) sounds() {
	fmt.Println("Dog sounds ", d.bark)
}

func (d *Dog) soundsThreeTimes() {
	d.bark = "urf!"
	d.bark = fmt.Sprintf("%v %v %v", d.bark, d.bark, d.bark)
	fmt.Println("Dog Barks : ", d.bark)
}
