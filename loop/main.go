package main

import "fmt"

func main() {
	fmt.Println("Hello This is Lec on Loops")

	data := []string{"sunday", "monday", "tuesday", "wednesday", "thursday", "friday", "saturday"}

	for d := 0; d < len(data); d++ {
		fmt.Println(data[d])
	}

	for d := range data {
		fmt.Println(data[d])
	}

}
