package main

import (
	"errors"
	"fmt"
)

/*
func main() {
	printMe()
}

func printMe() {
	fmt.Println("Hello, World!") // Function to print a message
}
*/

func main() {
	var printValue string = "Hello, World!" // Declare a variable to hold the string
	printMe(printValue)

	var a int = 10                                                          // Declare an integer variable
	var b int = 5                                                           // Declare another integer variable
	result := intDivision(a, b)                                             // Call the intDivision function and store the result
	fmt.Printf("The result of intdividing %d by %d is: %d\n", a, b, result) // Print the result

	quotient, remainder := intDivisionWithRemainder(a, b)                                                       // Call the function to get both quotient and remainder
	fmt.Printf("The result of intdividing %v by %v is: %v with a remainder of %v\n", a, b, quotient, remainder) // Print both values

	var result2, remainder2, err = intDiv(a, b)
	/*
		if err != nil { //if there is an error
			fmt.Println(err.Error()) //print the error
		} else if remainder2 == 0 {
			fmt.Printf(("The result of intdividing %v by %v is: %v with no remainder\n"), a, b, result2) // Print the result if no error and no remainder
		} else {
			fmt.Printf("The result of intdividing %v by %v is: %v with a remainder of %v\n", a, b, result2, remainder2) // Print both values
		}
	*/
	switch {
	case err != nil: //if there is an error
		fmt.Println(err.Error()) //print the error
	case remainder2 == 0:
		fmt.Printf("The result of intdividing %v by %v is: %v with no remainder\n", a, b, result2) // Print the result if no error and no remainder
	default:
		fmt.Printf("The result of intdividing %v by %v is: %v with a remainder of %v\n", a, b, result2, remainder2) // Print
	}

	switch remainder {
	case 0:
		fmt.Println("The division has no remainder")
	case 1, 2:
		fmt.Println("The division has a small remainder")
	default:
		fmt.Println("The division has a large remainder")
	}
}

func printMe(printValue string) {
	fmt.Println(printValue) // Function to print a message
}

func intDivision(numerator int, denominator int) int {
	if denominator == 0 {
		fmt.Println("Division by zero is not allowed")
		return 0
	}
	return numerator / denominator // Function to perform integer division
}

func intDivisionWithRemainder(numerator int, denominator int) (int, int) {
	if denominator == 0 {
		fmt.Println("Division by zero is not allowed")
		return 0, 0
	}
	quotient := numerator / denominator  // Calculate the quotient
	remainder := numerator % denominator // Calculate the remainder
	return quotient, remainder           // Return both values
}

func intDiv(numerator int, denomiator int) (int, int, error) {
	var err error
	if denomiator == 0 {
		err = errors.New("division by zero is not allowed")
		return 0, 0, err
	}
	var result int = numerator / denomiator
	var remainder int = numerator % denomiator
	return result, remainder, err
}
