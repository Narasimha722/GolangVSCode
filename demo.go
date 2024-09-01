package main

import (
	"fmt"
)

func main() {
	var num1, num2 float64
	var operator string

	// Taking input from the user
	fmt.Print("Enter the first number: ")
	fmt.Scanln(&num1)

	fmt.Print("Enter the operator (+, -, *, /): ")
	fmt.Scanln(&operator)

	fmt.Print("Enter the second number: ")
	fmt.Scanln(&num2)

	// Performing the calculation based on the operator
	var result float64
	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 != 0 {
			result = num1 / num2
		} else {
			fmt.Println("Error: Division by zero is not allowed.")
			return
		}
	default:
		fmt.Println("Error: Invalid operator.")
		return
	}

	// Printing the result
	fmt.Printf("\nThe result of %.2f %s %.2f is %.2f.\n", num1, operator, num2, result)
}
