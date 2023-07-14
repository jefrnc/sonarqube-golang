package main

import (
	"fmt"
)

// Calculate realiza la operación especificada en los dos números proporcionados.
// Soporta las operaciones de suma, resta, multiplicación y división.
func calculate(num1, num2 float64, operator string) float64 {
	result := 0.0

	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		result = num1 / num2
	default:
		fmt.Println("Invalid operator")
	}

	return result
}

func main() {
	var num1, num2 float64
	var operator string

	fmt.Print("Enter the first number: ")
	fmt.Scanln(&num1)

	fmt.Print("Enter the second number: ")
	fmt.Scanln(&num2)

	fmt.Print("Enter the operator (+, -, *, /): ")
	fmt.Scanln(&operator)

	result := calculate(num1, num2, operator)

	fmt.Printf("The result of the operation is: %.2f\n", result)
}
