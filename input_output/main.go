package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Hello This is Lec 02 input output")
	// making input reader this allow us to take input from input devices
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Number a : ")
	// taking input with the help of reader (input taker)
	a, err := reader.ReadString('\n')
	fmt.Print("Enter Number b : ")
	b, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}

	// using strings package & strconv in this to convert the input in number (INT)
	// strconv.ParseInt = take out number from string
	// strings.TrimSpace(a) = it remove or trim space in out case it's \n
	numberA, err := strconv.ParseInt(strings.TrimSpace(a), 10, 64)
	numberB, err := strconv.ParseInt(strings.TrimSpace(b), 10, 64)
	fmt.Println(numberA + numberB)
}
