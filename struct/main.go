package main

import "fmt"

func main() {
	fmt.Println("Hello This is Lec 06 stucts")

	// calling struct
	Anshit := User{"Anshit", "anshitmishra03@gmail.com", 20, "gwalior"}

	// getting details of struct
	fmt.Printf("user details %v \n", Anshit)

	// getting single calue
	fmt.Printf("user email %v \n", Anshit.Email)
}

// its like class and name of struct in capital
// there is no inheritance
type User struct {
	Name   string
	Email  string
	Age    int
	Addess string
}
