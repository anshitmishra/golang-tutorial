package main

import "fmt"

func main() {
	fmt.Println("Hello This is First Lec 01 variables")
	// string type variable
	var username string = "Anshit mishra"
	fmt.Println("username => ", username)
	fmt.Printf("username type is %T \n", username)

	// bool
	var isUserName bool = true
	fmt.Println("username is available => ", isUserName)
	fmt.Printf("username is available type is %T \n", isUserName)

	// int different types
	// ================================================================================

	// uint8       the set of all unsigned  8-bit integers (0 to 255)
	// uint16      the set of all unsigned 16-bit integers (0 to 65535)
	// uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
	// uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)
	var age uint8 = 20
	fmt.Println("age => ", age)
	var age2 uint = 20000
	fmt.Println("age2 => ", age2)

	// int8        the set of all signed  8-bit integers (-128 to 127)
	// int16       the set of all signed 16-bit integers (-32768 to 32767)
	// int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
	// int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)
	var number int8 = -20
	fmt.Println("number => ", number)
	var number2 int = 20000
	fmt.Println("number2 => ", number2)

	// float32     the set of all IEEE-754 32-bit floating-point numbers
	// float64     the set of all IEEE-754 64-bit floating-point numbers
	var float float32 = 20.2425212321312323123 // only takes 6 values after decimal
	fmt.Println("float => ", float)
	var float2 float64 = 20000.123312321312
	fmt.Println("float2 => ", float2)

	// default values and aliaese

	var initNumber int
	// no garbage value if you did't add any value it always take 0 as default value
	fmt.Println("initial number => ", initNumber) 

	// implicit type
	var website = "https://github.com/anshitmishra"
	fmt.Println(website) // here lexer came in picture it will decide for you the type of variable and you can't change it later on...

	// no var variable

	numberVariable := 123
	fmt.Println(numberVariable) // no var keyword is required when you use (:=) it will also auto define type of variable

	// const and public keyword
	const LoginToken string = "s12kj3n2kn312n"
	fmt.Println(LoginToken) // when we make first word of variable capital it is consider as public  it can be accessed by any function under the scope

}
