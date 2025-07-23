package main

import (
	"fmt"
)

func main() {
	var p *int32 = new(int32) // pointer to int32 value
	var i int32
	fmt.Printf("The value p points to is: %v\n", *p) // dereference pointer to get value
	fmt.Printf("The value of i is: %v\n", i)         // prints 0, as int32 is initialized to 0
	*p = 10                                          // set the value at the memory location pointed to by p to 10
	fmt.Printf("The value p points to is: %v\n", *p) // prints

	//on line 8 we are using * to say that we want to intialize a pointer to an int32 value
	//on lines 10 and 12 we are using * to say we want the value at the memory location pointed to by p

	p = &i // set p to point to the memory location of i
	*p = 1 // does the same as i = 1, but using pointer dereferencing

	//Slices

	var slice = []int32{1, 2, 3} // create a slice of int32 values
	var sliceCopy = slice        // create a copy of the slice
	sliceCopy[2] = 4             // modify the second element of the copy
	fmt.Println(slice)
	fmt.Println(sliceCopy) // both print [1 2 4], as slices are reference types, their pointers point to the same underlying array

	//pointers in functions
	var thing1 = [5]float64{1, 2, 3, 4, 5}                          // create an array of float64 values
	fmt.Printf("The memory address of thing1 is: %p\n", &thing1)    // print the memory address of the array
	result := square(thing1)                                        // pass the array to the square function
	fmt.Printf("The result of squaring the array is: %v\n", result) // print the result
	fmt.Printf("The value of thing1 is: %v\n", thing1)              // print the value of the array
	//Using twice the memory, as we are passing a copy of the array to the function

	var thing3 = [5]float64{1, 2, 3, 4, 5}                           // create an array of float64 values
	fmt.Printf("The memory address of thing3 is: %p\n", &thing3)     // print the memory address of the array
	result3 := squareWithPointer(&thing3)                            // pass the array to the square function
	fmt.Printf("The result of squaring the array is: %v\n", result3) // print the result
	fmt.Printf("The value of thing3 is: %v\n", thing3)               // print the value of the array
	//Using less memory, as we are passing a pointer to the array to the function
	//However, we are modifying the original array, so we need to be careful with this
}

func square(thing2 [5]float64) [5]float64 {
	fmt.Printf("The memory address of thing2 is: %p\n", &thing2) // print the memory address of the array passed to the function
	for i := range thing2 {
		thing2[i] *= thing2[i] // square each element in the array
	}
	return thing2 // return the modified array
}

func squareWithPointer(thing2 *[5]float64) [5]float64 {
	fmt.Printf("The memory address of thing2 is: %p\n", thing2) // print the memory address of the array passed to the function
	for i := range thing2 {
		thing2[i] *= thing2[i] // square each element in the array
	}
	return *thing2 // return the modified array
}
