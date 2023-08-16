package main

import "fmt"

func main() {
	fmt.Println("Hello Defer lec")

	defer fmt.Println("world")
	fmt.Println("hello")
	ReversePrint(1, 2, 3, 4, 5, 6, 7)
}

func ReversePrint(val ...int) {
	for i := 1; i <= len(val); i++ {
		defer fmt.Println(i)
	}
}
