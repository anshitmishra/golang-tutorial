package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Hello This is Lec on switch")
	rand.Seed(time.Now().Unix())
	value := rand.Intn(6) + 1

	switch value {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	case 4:
		fmt.Println("four")
	case 5:
		fmt.Println("five")
	case 6:
		fmt.Println("six")
	default:
		fmt.Println("unknown")
	}
}
