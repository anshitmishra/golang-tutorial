package main

import "fmt"

func main() {
	fmt.Println("Hello This is Lec on function")
	Greeting()
	data := Greeting2()
	fmt.Println(data)

	user := User{"anshit", "anshitmishra03@gmail.com"}
	fmt.Println(user)
	user.EmailGetter()
}

func Greeting() {
	fmt.Println("hello Sir")
}

func Greeting2() string {
	return "hello Sir func two"
}

type User struct {
	Name  string
	Email string
}

func (u User) EmailGetter() {
	u.Email = "mishrajii1714@gmail.com"
	fmt.Println(u.Email)
}
