package main

import "fmt"

func main() {
	fmt.Println("Hello This is Lec 05 Maps")

	language := make(map[int]string)

	language[0] = "reactjs"
	language[1] = "javascript"
	language[2] = "java"
	language[3] = "php"
	language[4] = "python"

	fmt.Println("language is ", language)

	// access single value
	fmt.Println("language is js ", language[1])

	// delete single value
	delete(language, 2)

	for key, value := range language {
		fmt.Printf("language key is %v and value is %v \n", key, value)
	}
}
