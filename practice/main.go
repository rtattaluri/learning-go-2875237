package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "http://services.explorecalifornia.org/json/tours.php"

func main() {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	responseBodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	//fmt.Println("Response body", string(responseBodyBytes))

	var jsonStringIntend bytes.Buffer
	err = json.Indent(&jsonStringIntend, responseBodyBytes, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Println("Response body", jsonStringIntend.String())
	//fmt.Println("Network requests")
}
