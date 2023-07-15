<div align="center">
  <img src="/images/structs/one.webp" alt="Go Struct Icon" width="200" height="200">
</div>

# Go Structs: A Comprehensive Guide

Welcome to the world of Go programming! In this README, we will explore the concept of structs in Go (Golang), providing a comprehensive understanding of how they work, their usage, and practical examples. Let's dive in and unlock the power of structs in Go!

## Table of Contents

- [Introduction to Structs](#introduction-to-structs)
  - [Struct Declaration](#struct-declaration)
  - [Struct Fields](#struct-fields)
  - [Struct Initialization](#struct-initialization)
- [Working with Structs](#working-with-structs)
  - [Accessing Struct Fields](#accessing-struct-fields)
  - [Modifying Struct Fields](#modifying-struct-fields)
- [Struct Methods](#struct-methods)
- [Embedding Structs](#embedding-structs)
- [Conclusion](#conclusion)

## Introduction to Structs

In Go, a struct is a composite data type that allows you to group related fields together. It is a way to define your own custom data structure. Structs provide a powerful mechanism for modeling real-world entities in your programs.

### Struct Declaration

To declare a struct in Go, use the `type` keyword followed by the struct name and the fields enclosed in curly braces. Here's an example:

```go
type Person struct {
    Name string
    Age  int
}
```

In this example, we declare a struct named Person with two fields: Name of type string and Age of type int.

## Struct Fields

Struct fields are the individual components or attributes of a struct. Each field has a name and a corresponding type. Here's an example:

```go

type Rectangle struct {
    Width  float64
    Height float64
}
```

In this example, we define a struct named Rectangle with two fields: Width and Height, both of type float64.

## Struct Initialization

To initialize a struct, you can use the struct literal syntax. Here's an example:

``` go
person := Person{Name: "Alice", Age: 25}
```

In this example, we create a new Person struct named person and initialize its Name field with "Alice" and its Age field with 25.

## Working with Structs
## Accessing Struct Fields

To access the fields of a struct, use the dot notation (.) followed by the field name. Here's an example:

```go

fmt.Println(person.Name) // Output: Alice
fmt.Println(person.Age)  // Output: 25
```

In this example, we access the Name and Age fields of the person struct and print their values.

## Modifying Struct Fields

Struct fields can be modified by assigning new values to them. Here's an example:

```go

person.Age = 26
fmt.Println(person.Age) // Output: 26
```

In this example, we modify the Age field of the person struct by assigning a new value (26) to it. The modified value is then printed.

## Struct Methods

Go supports methods, which are functions associated with a particular type. You can define methods on structs to provide behavior or functionality specific to that struct. Here's an example:

```go

type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}
```

In this example, we define a method named Area() on the Circle struct. This method calculates the area of a circle based on its Radius field.

## Embedding Structs

Go allows you to embed one struct within another to create more complex data structures. This feature is called struct embedding. Here's an example:

```go

type Person struct {
    Name string
    Age  int
}

type Employee struct {
    Person
    Salary float64
}
```

In this example, we define a Person struct and an Employee struct that embeds the Person struct. The Employee struct inherits the fields and methods of the Person struct.

## Conclusion

Congratulations! You have now gained a solid understanding of structs in Go. Structs allow you to define custom data structures with named fields, providing a powerful way to model real-world entities in your programs. By utilizing structs effectively, you can organize and manipulate complex data.

Remember to use structs when you need to group related fields together and define custom data types. Structs enable you to create well-organized and maintainable code.

Now go forth and leverage the power of structs in your Go programs!