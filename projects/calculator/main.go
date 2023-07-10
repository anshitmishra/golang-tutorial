package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func addition() {
	var numberA int64
	var numberB int64
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("- Enter value a : ")
	inputA, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	} else {
		numberA, _ = strconv.ParseInt(strings.TrimSpace(inputA), 10, 32)
	}

	fmt.Print("- Enter value b : ")
	inputB, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	} else {
		numberB, _ = strconv.ParseInt(strings.TrimSpace(inputB), 10, 32)
	}

	fmt.Println("Total of a & b : ", numberA+numberB)
}

func substraction() {
	var numberA int64
	var numberB int64
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("- Enter value a : ")
	inputA, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	} else {
		numberA, _ = strconv.ParseInt(strings.TrimSpace(inputA), 10, 32)
	}

	fmt.Print("- Enter value b : ")
	inputB, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	} else {
		numberB, _ = strconv.ParseInt(strings.TrimSpace(inputB), 10, 32)
	}

	fmt.Println("Total of a & b : ", numberA-numberB)
}

func multiplication() {
	var numberA float64
	var numberB float64
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("- Enter value a : ")
	inputA, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	} else {
		numberA, _ = strconv.ParseFloat(strings.TrimSpace(inputA), 32)
	}

	fmt.Print("- Enter value b : ")
	inputB, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	} else {
		numberB, _ = strconv.ParseFloat(strings.TrimSpace(inputB), 32)
	}

	fmt.Println("Total of a & b : ", numberA*numberB)

}

func division() {
	var numberA float64
	var numberB float64
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("- Enter value a : ")
	inputA, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	} else {
		numberA, _ = strconv.ParseFloat(strings.TrimSpace(inputA), 32)
	}

	fmt.Print("- Enter value b : ")
	inputB, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	} else {
		numberB, _ = strconv.ParseFloat(strings.TrimSpace(inputB), 32)
	}

	fmt.Println("Total of a & b : ", numberA/numberB)

}

func main() {
	fmt.Println("Hello This is project 01 calculator")
	fmt.Println("- Type 1 for addition")
	fmt.Println("- Type 2 for substraction")
	fmt.Println("- Type 3 for multiplication")
	fmt.Println("- Type 4 for division")

	reader := bufio.NewReader(os.Stdin)

	inputOption, err := reader.ReadString('\n')

	if err != nil {
		panic(err)
	} else {
		options, _ := strconv.ParseInt(strings.TrimSpace(inputOption), 10, 64)
		if options == 1 {
			addition()
		}
		if options == 2 {
			substraction()
		}
		if options == 3 {
			multiplication()
		}
		if options == 4 {
			division()
		}
	}
}
