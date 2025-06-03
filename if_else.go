package main

import "fmt"

func main() {
	age := 25

	// if/else statement
	if age >= 18 {
		fmt.Println("You are an adult")
	} else {
		fmt.Println("You can not have cerveza")
	}

	// Switch statement
	fruit := "apple"
	switch fruit {
	case "apple":
		fmt.Println("apple!")
	case "banana":
		fmt.Println("banana!")
	default:
		fmt.Println("unknown fruit")
	}

	switch fruit {
	case "apple", "banana", "mango":
		fmt.Println("cool fruits!")
	case "tomato":
		fmt.Println("boring")
	default:
		fmt.Println("unknown fruit")
	}

	// Switch with expressions
	number := 15

	switch {
	case number < 10:
		fmt.Println("number is less than 10")
	case number >= 10 && number < 20:
		fmt.Println("number is greater than 10 but less than 20")
	default:
		fmt.Println("number is bigger than 20")
	}

	checkType(2)
	checkType("ola k ase")
	checkType(2.231)
	checkType(nil)
}

func checkType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("int")
	case float64:
		fmt.Println("float64")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("unknown")
	}
}
