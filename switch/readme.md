<p align="center" >
  <img src="https://github.com/anshitmishra/golang-tutorial/blob/main/images/switch/one.png" />
</p>

# Golang Switch Case README

## Introduction

This README provides a comprehensive guide on using the `switch` statement in the Go programming language (Golang). The `switch` statement is a versatile control flow construct used to make decisions based on the value of an expression. This document will help you understand its syntax, usage, and best practices.

## Table of Contents

1. [Basic Syntax](#basic-syntax)
2. [Expression Types](#expression-types)
3. [Comparison Rules](#comparison-rules)
4. [Multiple Expressions](#multiple-expressions)
5. [Fallthrough](#fallthrough)
6. [Type Switch](#type-switch)
7. [Empty `switch`](#empty-switch)
8. [Conclusion](#conclusion)

## 1. Basic Syntax <a name="basic-syntax"></a>

The basic syntax of a `switch` statement in Go looks like this:

```go
switch expression {
case value1:
    // Code to execute when expression equals value1
case value2:
    // Code to execute when expression equals value2
default:
    // Code to execute when expression does not match any case
}
```

expression: The expression whose value you want to compare with the cases.
value1, value2, etc.: The possible values to compare expression against.
default: An optional case that is executed when none of the other cases match.
## 2. Expression Types
The expression in a switch statement can be of any type, including numeric types, strings, or even custom types.


```go
var dayOfWeek = "Monday"

switch dayOfWeek {
case "Monday":
    fmt.Println("It's Monday!")
case "Tuesday":
    fmt.Println("It's Tuesday!")
default:
    fmt.Println("It's another day of the week.")
}
```


## 3. Comparison Rules
Go's switch statements are not limited to constants like some other languages. You can use any type as the expression, and the cases are compared against the expression's value.
The cases are evaluated from top to bottom until a match is found.
If a case matches, its associated code block is executed, and the switch statement exits.
If no cases match, the code block under default (if provided) is executed.

## 4. Multiple Expressions
You can use multiple expressions in a single switch statement, and Go will evaluate each of them. If any of the expressions match, the corresponding case will execute.

```go
var a, b int = 1, 2

switch a + b {
case 0:
    fmt.Println("Sum is zero")
case 3:
    fmt.Println("Sum is three")
default:
    fmt.Println("Sum is something else")
}
```

## 5. Fallthrough
In Go, the fallthrough keyword can be used within a case to transfer control to the next case, even if it doesn't match the condition. Be cautious when using fallthrough, as it can lead to unexpected behavior.

```go
var num = 2

switch num {
case 1:
    fmt.Println("One")
    fallthrough
case 2:
    fmt.Println("Two")
case 3:
    fmt.Println("Three")
}

```


## 6. Type Switch

Go also supports type switches, which are used to compare the type of an interface value.


```go
var x interface{} = 10

switch x.(type) {
case int:
    fmt.Println("x is an integer")
case string:
    fmt.Println("x is a string")
default:
    fmt.Println("x is of an unknown type")
}

```

## 7. Empty switch

An empty switch statement can be used as an alternative to a long if-else chain when you have multiple conditions to evaluate.

```go
var age = 25

switch {
case age < 18:
    fmt.Println("You are a minor")
case age >= 18 && age < 65:
    fmt.Println("You are an adult")
default:
    fmt.Println("You are a senior citizen")
}

```

## 8. Conclusion

The switch statement in Go is a powerful tool for making decisions in your code. It allows you to compare an expression against multiple values and execute different code blocks accordingly. Understanding the rules and usage of switch can make your Go code more readable and maintainable.

For more information on Go and its features, refer to the official Go documentation.

Feel free to contribute to this README or use it as a reference in your Go projects!