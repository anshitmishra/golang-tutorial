<div align="center">
  <img src="/images/array/one.webp" alt="Go Arrays and Slices Icon" height="200">
</div>

# Go Arrays and Slices: An Introduction

Welcome to the world of Go (Golang) arrays and slices! This README provides a comprehensive introduction to arrays and slices in Go, explaining their purpose, usage, and providing examples to illustrate their behavior.

## 🎯 What are Arrays?

In Go, an array is a fixed-size sequence of elements of the same type. The size of an array is determined at compile time and cannot be changed during runtime. Arrays provide a way to store and access multiple values in a contiguous block of memory.

### Array Declaration and Initialization

To declare an array in Go, specify the type of its elements and the number of elements enclosed in square brackets. Here's an example:

```go
var numbers [5]int
```

In this example, we declare an array named numbers that can store 5 integers. The array is initialized with the zero value of the element type, which is int in this case.

You can also initialize an array with values using an array literal. For example:

```go
var fruits = [3]string{"apple", "banana", "cherry"}
```

Here, we declare an array named fruits that can store 3 strings. The array is initialized with the specified values.
Accessing Array Elements

To access an element of an array, use square brackets [] with the index of the element. Array indices start from 0. Here's an example:

```go
var numbers = [5]int{1, 2, 3, 4, 5}
fmt.Println(numbers[0]) // Output: 1
```

In this example, we access the first element of the numbers array using numbers[0] and print its value, which is 1.
## 🌟 What are Slices?

Slices are a more flexible and dynamic alternative to arrays in Go. Unlike arrays, slices have a variable length and can be resized during runtime. Slices provide a view into an underlying array and allow for more convenient manipulation of sequences of elements.
Slice Declaration and Initialization

To declare a slice in Go, specify the type of its elements without specifying the length. Here's an example:


```go
var numbers []int
```

In this example, we declare a slice named numbers that can store integers. The slice is initialized as an empty slice.

You can also create a slice from an existing array or another slice using a slicing expression. For example:

```go
var fruits = [5]string{"apple", "banana", "cherry", "durian", "elderberry"}
var slicedFruits = fruits[1:4] // Create a slice from index 1 to 3 (inclusive)
```

Here, we create a slice named slicedFruits from the fruits array, containing elements from index 1 to 3 (inclusive). The resulting slice will be ["banana", "cherry", "durian"].
Modifying Slices

Slices allow for convenient modification operations, such as appending elements, slicing, and reslicing. Here are a few examples:

```go
var numbers = []int{1, 2, 3, 4, 5}
numbers = append(numbers, 6)      // Append an element
var slicedNumbers = numbers[1:4]  // Create a slice from index 1 to 3 (inclusive)
slicedNumbers = slicedNumbers[1:] // Reslice from index 1 to end
```

In this example, we append an element (6) to the numbers slice, create a new slice (slicedNumbers) from index 1 to 3 (inclusive), and then reslice slicedNumbers from index 1 to the end. Note that all these operations modify the underlying slice.


## 🏁 Conclusion

Arrays and slices are fundamental concepts in Go that allow you to work with sequences of elements efficiently. While arrays provide a fixed-size storage, slices offer more flexibility with variable lengths. By understanding and utilizing arrays and slices effectively, you can build powerful and dynamic programs in Go.

Remember to use arrays when you know the exact size of the sequence in advance, and use slices when you need a more dynamic and resizable collection.

Happy coding with Go arrays and slices!


## 🚀 Want Start

To get started with the code examples and projects in this repository, follow these steps:

1. Clone the repository to your local machine.
   ```shell
   git clone https://github.com/anshitmishra/golang-tutorial.git


## 🤝 Contributing

We welcome contributions to this repository! If you have ideas for new projects, improvements to existing projects, or additional code examples, feel free to submit a pull request. Let's collaborate and make this repository even better!

Please ensure that your contributions align with the purpose and standards of this repository. For more details, refer to the [CONTRIBUTING](CONTRIBUTING.md) file.

## 📄 License

This repository is licensed under the [MIT License](LICENSE). You are free to use, modify, and distribute the code in this repository for personal or commercial purposes. However, we assume no liability for any misuse or damages caused by the use of this code. Please review the license file for more details.
