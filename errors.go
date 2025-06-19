package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	sq, e := sqrt(16)
	if e != nil {
		fmt.Println(e)
	} else {
		fmt.Printf("%.4f\n", sq)
	}

	sq, e = sqrt(-2)
	if e != nil {
		fmt.Println(e)
	} else {
		fmt.Printf("%.4f\n", sq)
	}

	data := []byte{}
	if err := process(data); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Successfully processed data")
	}

	customError := errorProcess()
	if customError != nil {
		fmt.Println(customError)
	}

	erro := readData()
	if erro != nil {
		fmt.Println(erro)
	} else {
		fmt.Println("Successfully read data")
	}
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("x must be greater or equal to 0")
	} else {
		return math.Sqrt(x), nil
	}
}

func process(data []byte) error {
	if len(data) == 0 {
		return errors.New("data is empty")
	} else {
		return nil
	}
}

type myError struct {
	message string
}

func (e *myError) Error() string {
	return fmt.Sprintf("Error: %s", e.message)
}

func errorProcess() error {
	return &myError{"Custom error message"}
}

func readData() error {
	err := readConfig()
	if err != nil {
		return fmt.Errorf("Error reading config file: %w", err)
	}
	return nil
}

func readConfig() error {
	return errors.New("Read config error")
}
