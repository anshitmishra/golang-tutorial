package main

import "fmt"

// example why we need to use pointers
func add(d *int) int {
	*d = *d + 1
	return *d
}

func main() {
	fmt.Println("Hello This is Lec 03 Pointer")
	var ptr *int //To declare a pointer
	mynumber := 12
	ptr = &mynumber // pass by reference
	fmt.Println("value of pointer ", *ptr)
	fmt.Println(add(&mynumber))
	fmt.Println(mynumber)
}
