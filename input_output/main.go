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

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Number a : ")
	a, err := reader.ReadString('\n')
	fmt.Print("Enter Number b : ")
	b, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	numberA, err := strconv.ParseInt(strings.TrimSpace(a), 10, 64)
	numberB, err := strconv.ParseInt(strings.TrimSpace(b), 10, 64)
	fmt.Println(numberA + numberB)
}
