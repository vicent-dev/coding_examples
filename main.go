package main

import "fmt"

func sum(n1 float32, n2 float32) float32 {
	return n1 + n2
}

func sub(n1 float32, n2 float32) float32 {
	return n1 - n2
}

func mult(n1 float32, n2 float32) float32 {
	return n1 * n2
}

func main() {
	var number1, number2 float32

	number1 = 2.5
	number2 = 3.89

	fmt.Printf("%f + %f = %f \n", number1, number2, sum(number1, number2))
	fmt.Printf("%f - %f = %f \n", number1, number2, sub(number1, number2))
	fmt.Printf("%f * %f = %f \n", number1, number2, mult(number1, number2))
}
