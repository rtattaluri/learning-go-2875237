package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Files")
	writeToFile("./myownfile", "This is a message\n")
	fileread := readFromFile("./myownfile")
	fmt.Println("Output from file read: ", fileread)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func writeToFile(filename string, contentToWrite string) {
	/*newFile, err := os.Create(filename)
	checkError(err)
	len, err := newFile.Write([]byte(contentToWrite))
	checkError(err)
	fmt.Printf("Wrote to file %v characters\n", len)*/

	err := ioutil.WriteFile(filename, []byte(contentToWrite), fs.FileMode(os.O_APPEND|os.O_WRONLY|os.O_CREATE|os.O_RDWR))
	checkError(err)
	newFile, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0655)
	checkError(err)
	secondMessage := "This is second message\n"
	len, err := newFile.Write([]byte(secondMessage))
	checkError(err)
	fmt.Println("Number of chars in second message: ", len)
	defer newFile.Close()
}

func readFromFile(filename string) string {
	filecontent, err := ioutil.ReadFile(filename)
	checkError(err)
	return string(filecontent)
}
