package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Hello This is Lec 04 array")

	var fruitList [4]string

	fruitList[0] = "apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "orange"

	fmt.Println("fruit list ", fruitList)
	fmt.Println("fruit list length", len(fruitList))

	var veglist = [3]string{"potato", "beans", "tomato"}

	fmt.Println("veg list ", veglist)

	// slices
	// =======================================================

	fmt.Println("Hello This is First Lec 04 part b slices")

	var slices = []string{}

	fmt.Printf("type of this slice is %T \n", slices)

	slices = append(slices, "anshit", "animesh", "agam")
	fmt.Println("values of this slice is ", slices)

	slices = append(slices[1:])
	fmt.Println("values of this slice is ", slices)
	slices = append(slices[1:2])
	fmt.Println("values of this slice is ", slices)

	// another way to make array
	marks := make([]int, 4)
	marks[0] = 200
	marks[1] = 220
	marks[2] = 190
	marks[3] = 120

	// if i add another it will throw an error
	// marks[4] = 1201

	// but by using append i can allocate new element

	marks = append(marks, 1231)
	fmt.Println("values of this marks is ", marks)

	sort.Ints(marks) // sort marks in increasing order
	fmt.Println("values after sort marks is ", marks)

}
