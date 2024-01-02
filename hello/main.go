package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Write and read local text files")
	content := "Hello from Go"
	file, err := os.Create("./fromString.txt")
	checkError(err)
	length, err := io.WriteString(file, content)
	checkError(err)
	fmt.Printf("Wrote a file with %v characters\n", length)
	defer file.Close() //wait until everything else is done and then execute this command. And I'll call file.Close.
	defer readFile("./fromString.txt")
}

func readFile(fileName string) {
	data, err := ioutil.ReadFile(fileName)
	checkError(err)
	fmt.Println("Text read from the file", string(data))
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
