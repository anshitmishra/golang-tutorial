package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Hello This is FileReadWrite")
	WriteFile("./file.txt", "Hello this is anshitmishra repo https://github.com/anshitmishra follow me")
	ReadFile("./file.txt")
}

func WriteFile(fileName string, message string) {
	file, err := os.Create(fileName)
	CheckErrorNil(err)

	length, err := io.WriteString(file, message)
	CheckErrorNil(err)

	fmt.Println("length of content ", length)
}

func ReadFile(fileName string) {
	data, err := os.ReadFile(fileName)
	CheckErrorNil(err)
	fmt.Println(string(data))
}

func CheckErrorNil(err error) {
	if err != nil {
		panic(err)
	}
}
