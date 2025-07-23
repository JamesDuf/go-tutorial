package main

import (
	"fmt"
	"unicode/utf8"
) //importing fmt package to use its functions
//importing unicode/utf8 package to use its functions

func main() {
	var intNum int //int8 int16 int32 int64 	//uint8 uint16 uint32 uint64
	fmt.Println(intNum)

	var floatNum float64 //float32 float64  	//float32 is less precise than float64
	fmt.Println(floatNum)

	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2
	var result float32 = floatNum32 + float32(intNum32) //type conversion
	fmt.Println(result)

	var intNumm1 int = 3
	var intNumm2 int = 2
	fmt.Println(intNumm1 / intNumm2) //division   //int division truncates the decimal part =1
	fmt.Println(intNumm1 % intNumm2) //modulus //remainder of division = 1

	var myString string = "Hello \nWorld!" //string type double quotes
	fmt.Println(myString)

	var myString2 string = `Hello 
	World!` //raw string type backticks
	fmt.Println(myString2)

	fmt.Println(len("A"))                    //length of string is fine for ASCII characters
	fmt.Println(len("𐍈"))                    //length of string is more than 1 for irregular cahracters
	fmt.Println(utf8.RuneCountInString("𐍈")) //this works with fancy characters

	var myRune rune = '𐍈' //rune type is an alias for int32, used for Unicode characters
	fmt.Println(myRune)   //prints the Unicode code point of the character

	var myByte byte = 'A' //byte type is an alias for uint8, used for ASCII characters
	fmt.Println(myByte)   //prints the ASCII value of the character

	var myBool bool = true //boolean type
	fmt.Println(myBool)

	//Defualt values
	var defaultInt int
	var defaultFloat float64
	var defaultString string
	var defaultBool bool
	fmt.Println(defaultInt, defaultFloat, defaultString, defaultBool) //prints 0 0 "" false

	//short variable declaration
	var myVar = "Hello" //type inference
	fmt.Println(myVar)
	myVar2 := "Hello"   //same as above, but shorter syntax
	fmt.Println(myVar2) //prints "Hello"

	myVar3, myVar4 := "Hello", 10 //multiple variable declaration
	fmt.Println(myVar3, myVar4)   //prints "Hello 10"

	//Where this can get confusing and you might want to use the var and type keywords
	//myVar5 := foo() is unclear
	//var myVar5 string = foo() //using var keyword to declare a variable with a function return value
	//fmt.Println(myVar5) //prints "Hello from foo"

	const myConst string = "constant value" //constant type, cannot be changed, needs to be initialized at declaration
	//const myConst2 = 10 //constant type, can be inferred type, cannot be declared without initialization
	fmt.Println(myConst) //prints "constant value"
}
