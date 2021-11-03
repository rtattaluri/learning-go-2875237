package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const url = "http://services.explorecalifornia.org/json/tours.php"

func main() {

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	//fmt.Printf("Response type: %T\n", resp)

	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	content := string(bytes)
	fmt.Print(content)
	//fmt.Println("Network requests")
	allTours := toursFromJson(content)
	fmt.Println("Tour Name", ":", "Tour Price")
	for _, eachTour := range allTours {
		fmt.Println(eachTour.Name, ":", eachTour.Price, eachTour.Difficulty)
	}
}

type Tour struct {
	TourId, PackageId, PackageTitle, Name, Blurb, Description, Price, Difficulty, Length, Graphic, Region string
}

func toursFromJson(jsonString string) []Tour {
	tours := make([]Tour, 0)
	jsonDecoder := json.NewDecoder(strings.NewReader(jsonString))
	_, err := jsonDecoder.Token()
	if err != nil {
		panic(err)
	}
	var eachTour Tour
	for jsonDecoder.More() {
		err = jsonDecoder.Decode(&eachTour)
		if err != nil {
			panic(err)
		}
		tours = append(tours, eachTour)
	}
	return tours
}
