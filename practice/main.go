package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	dow := rand.Intn(7) + 1
	fmt.Println("Day", dow)

	var result string
	switch dow {
	case 0:
		result = "sunday"
	case 1:
		result = "monday"
	case 2:
		result = "tuesday"
	case 3:
		result = "wednesday"
	case 4:
		result = "thursday"
	case 5:
		result = "friday"
	case 6:
		result = "saturday"
	case 7:
		result = "sunday"
	default:
		result = "no other day left!!!"
	}

	fmt.Println(" Day of week: ", result)
}
