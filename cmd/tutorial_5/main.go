package main

import (
	"fmt"
	"strings"
)

func main() {
	var myString = "résumé" // string with Unicode characters
	fmt.Println(myString)
	var indexed = myString[0]                // indexing a string returns a byte (uint8)
	fmt.Println(indexed)                     // prints the byte value of the first character
	fmt.Printf("%v, %T\n", indexed, indexed) // prints the value and type of indexed

	//rune = int32
	var myString2 = []rune("résumé") // convert string to rune slice
	fmt.Println(myString2[1])

	var strSlice = []string{"a", "b", "c"} // create a slice of strings
	var strBuilder strings.Builder
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i]) // append each string to the builder
	}
	var catStr = strBuilder.String() // convert builder to string
	fmt.Println(catStr)              // prints "abc"
}
