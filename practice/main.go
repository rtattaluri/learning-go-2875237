package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Dates and times")
	t := time.Now()
	fmt.Println("Current time is : ", t)

	t2 := time.Date(1987, time.July, 6, 00, 15, 0, 0, time.Now().Location())
	utc_t2 := t2.UTC()
	fmt.Println("I was born on date ", utc_t2)
	fmt.Println("Ansi date is :", utc_t2.Format(time.ANSIC))

	parsedDate, err := time.Parse(time.ANSIC, "Sun Jul  5 18:45:00 1987")
	if err != nil {
		fmt.Println("failed while parsing date", err)
	} else {
		fmt.Println("Parsed Date : ", parsedDate)
	}

	age := t.Sub(t2).Hours() / 24 / 365.25
	fmt.Println("Age is ", age) // 366

}
