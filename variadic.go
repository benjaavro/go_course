package main

import "fmt"

func main() {
	// Infinite number of params
	operation, sum := sumOfIntegers("Sum", 1, 3, 5, 1)
	fmt.Println(operation, sum)

	// Infinite number of params passed as slice/list
	numbers := []int{1, 2, 3, 4, 5}
	operation, sum = sumOfIntegers("Second Sum", numbers...)
	fmt.Println(operation, sum)
}

// ... Ellipsis (0+ params)
func sumOfIntegers(returnString string, numbers ...int) (string, int) {
	sum := 0
	// Received params are a list(slice) of given type
	for _, number := range numbers {
		sum += number
	}
	return returnString, sum
}
