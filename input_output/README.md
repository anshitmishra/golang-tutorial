<p align="center" >
  <img src="https://github.com/anshitmishra/golang-tutorial/blob/main/images/variable/one.webp" />
</p>

# Go (Golang) Variables

Welcome to the Go (Golang) Variables repository! 🚀 This repository is dedicated to explaining the concept of variables in the Go programming language.

## 📚 Introduction

In Go, variables are used to store and manipulate data. They provide a way to hold values of different types such as numbers, strings, and more. Understanding how variables work in Go is essential for writing effective and flexible programs.

This repository aims to provide an in-depth understanding of variables in Go, covering topics such as variable declaration, initialization, data types, scope, and best practices.

## 📂 Code Structure

variable init
```go
   // string type variable
	var username string = "Anshit mishra"
	fmt.Println("username => ", username)
	fmt.Printf("username type is %T \n", username)
```

boolean variable only have two values true/false
```go
	// bool
	var isUserName bool = true
	fmt.Println("username is available => ", isUserName)
	fmt.Printf("username is available type is %T \n", isUserName)
```

integer variable with different types
```go
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

```

golang default values

** if you did't init any value to variable then for string is empty and for integer is 0 for bool it false **
```go
// default values and aliaese

	var initNumber int
	fmt.Println("initial number => ", initNumber) // no garbage value if you did't add any value it always take 0 as default value
```

golang lexer will auto define value for that variable if you don't define
	```go
 // implicit type
	var website = "https://github.com/anshitmishra"
	fmt.Println(website) // here lexer came in picture it will decide for you the type of variable and you can't change it later on...

	// no var variable
	numberVariable := 123
	fmt.Println(numberVariable) // no var keyword is required when you use (:=) it will also auto define type of variable
```

it make public variable that is accessable from any fun under the scope
```go
	// const and public keyword
	const LoginToken string = "s12kj3n2kn312n"
	fmt.Println(LoginToken) // when we make first word of variable capital it is consider as public  it can be accessed by any function under the scope
```



## 🚀 Getting Started

To get started with the code examples and projects in this repository, follow these steps:

1. Clone the repository to your local machine.
   ```shell
   git clone https://github.com/anshitmishra/golang-tutorial.git


## 🤝 Contributing

We welcome contributions to this repository! If you have ideas for new projects, improvements to existing projects, or additional code examples, feel free to submit a pull request. Let's collaborate and make this repository even better!

Please ensure that your contributions align with the purpose and standards of this repository. For more details, refer to the [CONTRIBUTING](CONTRIBUTING.md) file.

## 📄 License

This repository is licensed under the [MIT License](LICENSE). You are free to use, modify, and distribute the code in this repository for personal or commercial purposes. However, we assume no liability for any misuse or damages caused by the use of this code. Please review the license file for more details.
