package main

import (
	"fmt"
	"io"
	"net/http"
)

var webUrl string = "https://example.com"

func main() {
	fmt.Println("Hello this is web request")

	response, err := http.Get(webUrl)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	databyte, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	content := string(databyte)

	fmt.Println(content)
}
