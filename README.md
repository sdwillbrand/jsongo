# JSONGO: Golang JSON Parser

## Table of Contents

1. **Introduction**
   - 1.1 Purpose
   - 1.2 Features
   - 1.3 Project Structure
2. **Getting Started**
   - 2.1 Installation
   - 2.2 Usage Example
3. **API Reference**
   - 3.1 `ParseFromString(data string) *JSValue`
   - 3.2 `ParseFromByte(data []byte) *JSValue`
   - 3.3 `ParseFromRune(data []rune) *JSValue`
4. **JSValue Struct**
   - 4.1 `Type() Type`
   - 4.2 `String() string`
   - 4.3 `Int() (int64, error)`
   - 4.4 `Float() (float64, error)`
   - 4.5 `Bool() (bool, error)`
   - 4.6 `Array() ([]*JSValue, error)`
   - 4.7 `Object() (map[string]*JSValue, error)`

5. **Contributing**
   - 5.1 Bug Reports and Feature Requests
   - 5.2 Pull Requests

6. **License**

---

## 1. Introduction

Welcome to the JSONGO documentation! JSONGO is a JSON parsing library for the Go programming language that provides easy-to-use functions to parse JSON data into a structured format.

### 1.1 Purpose

JSONGO aims to simplify JSON parsing in Go applications by offering a user-friendly API that allows you to parse JSON data from strings, byte slices, or rune slices into structured `JSValue` objects. These objects can then be easily manipulated and queried to extract the necessary data.

### 1.2 Features

- **Multiple Input Formats:** JSONGO supports parsing JSON data from strings, byte slices, and rune slices, giving you flexibility in handling different data sources.
- **Structured Parsing:** The library parses JSON data into a structured format represented by the `JSValue` struct, making it easy to navigate and extract data.
- **Data Extraction:** `JSValue` provides methods to extract data of various types (string, int, float, bool, array, object) from the parsed JSON structure.
- **Error Handling:** The library handles parsing errors gracefully, returning meaningful error messages to aid in debugging.
- **Ease of Use:** JSONGO's functions and methods are designed to be intuitive and straightforward.

### 1.3 Project Structure

```
jsongo/
|-- jsongo.go
|-- jsvalue.go
|-- LICENSE
|-- README.md
|-- .gitignore
|-- example/
|   |-- main.go
|
|-- tests/
|   |-- jsongo_test.go
|-- .gitignore
```

- `jsongo.go`: Contains the main parsing functions.
- `jsvalue.go`: Defines the `JSValue` struct and its methods.
- `example`: Contains an example usage of JSONGO.
- `tests`: Unit tests for the JSON parser.
- `LICENSE`: The project's license information.
- `README.md`: The main documentation file.

---

## 2. Getting Started

### 2.1 Installation

To use JSONGO in your project, use the `go get` command:

```sh
go get github.com/countersoda/jsongo
```

### 2.2 Usage Example

Here's a basic example of how to use JSONGO in your application:

```go
package main

import (
	"fmt"
	"github.com/yourusername/jsongo"
)

func main() {
	jsonString := `{"name": "John", "age": 30, "city": "New York"}`

	jsValue := jsongo.ParseFromString(jsonString)

	if jsValue != nil {
		fmt.Println(jsValue.String())
	}
}
```

---

## 3. API Reference

### 3.1 `ParseFromString(data string) *JSValue`

This function parses the provided JSON string and returns a pointer to a `JSValue` struct representing the parsed JSON structure.

**Parameters:**

- `data string`: The JSON data as a string.

**Returns:**

- `*JSValue`: A pointer to the parsed JSON structure (`JSValue`).

### 3.2 `ParseFromByte(data []byte) *JSValue`

This function parses the provided JSON data as a byte slice and returns a pointer to a `JSValue` struct representing the parsed JSON structure.

**Parameters:**

- `data []byte`: The JSON data as a byte slice.

**Returns:**

- `*JSValue`: A pointer to the parsed JSON structure (`JSValue`).

### 3.3 `ParseFromRune(data []rune) *JSValue`

This function parses the provided JSON data as a rune slice and returns a pointer to a `JSValue` struct representing the parsed JSON structure.

**Parameters:**

- `data []rune`: The JSON data as a rune slice.

**Returns:**

- `*JSValue`: A pointer to the parsed JSON structure (`JSValue`).

---

## 4. JSValue Struct

The `JSValue` struct represents the parsed JSON structure. It provides various methods to extract data of different types from the JSON data.

### 4.1 `Type() Type`

This method returns the type of the JSON value represented by the `JSValue`.

### 4.2 `String() string`

This method returns the JSON value as a string.

### 4.3 `Int() (int64, error)`

This method returns the JSON value as an integer.

### 4.4 `Float() (float64, error)`

This method returns the JSON value as a floating-point number.

### 4.5 `Bool() (bool, error)`

This method returns the JSON value as a boolean.

### 4.6 `Array() ([]*JSValue, error)`

This method returns the JSON value as an array of `JSValue` pointers.

### 4.7 `Object() (map[string]*JSValue, error)`

This method returns the JSON value as an object (map) of `JSValue` pointers.

---

## 5. Contributing

Contributions to JSONGO are welcome! If you encounter any bugs or have ideas for improvements, please follow the guidelines below.

### 5.1 Bug Reports and Feature Requests

If you find a bug or have a feature request, please open an issue on the project's GitHub repository. Provide as much detail as possible to help us understand and address the problem.

### 5.2 Pull Requests

If you want to contribute code to JSONGO, follow these steps:

1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Make your changes and ensure that the code passes all tests.
4. Open a pull request, explaining the changes you've made.

We will review your pull request and provide feedback as needed.

---

## 6. License

JSONGO is licensed under the [MIT License](LICENSE).

---

Feel free to explore, use, and contribute to the JSONGO project. Happy JSON parsing! 🚀