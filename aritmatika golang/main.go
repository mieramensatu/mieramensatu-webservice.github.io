package main

import "fmt"

// math input of addition or subtraction or multiplication or division values

func main() {
	var input1, input2 int
	var operator string
	fmt.Println("Enter the first number")
	fmt.Scanln(&input1)
	fmt.Println("Enter the second number")
	fmt.Scanln(&input2)
	fmt.Println("Enter the operator")
	fmt.Scanln(&operator)
	switch operator {
	case "+":
		fmt.Println(input1 + input2)
	case "-":
		fmt.Println(input1 - input2)
	case "*":
		fmt.Println(input1 * input2)
	case "/":
		fmt.Println(input1 / input2)
	default:
		fmt.Println("Invalid operator")
	}
}