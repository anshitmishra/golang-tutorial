package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello This is Lec 07 array")

	var age int

	fmt.Scanf("%d", &age)

	if age >= 18 {
		fmt.Println("you can vote")
	} else if age > 0 && age < 18 {
		fmt.Println("you can't vote")
	} else {
		fmt.Println("Enter correct age")
	}

	// different type of confition

	if num := 3; num < 5 {
		fmt.Println("this is init under if condition")
	}
}
