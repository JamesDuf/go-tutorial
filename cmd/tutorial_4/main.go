package main

import (
	"fmt"
)

func main() {
	var intArr [3]int32
	intArr[0] = 123
	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3])

	//print memory location of each element in the array
	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])

	intArr2 := [3]int32{1, 2, 3} //array with initial values
	fmt.Println(intArr2)

	intArr3 := [...]int32{1, 2, 3} //array with initial values
	fmt.Println(intArr3)

	var intSlice []int32 = []int32{4, 5, 6}                                                //slice is a dynamic array
	fmt.Printf("intSlice: %v, len: %d, cap: %d\n", intSlice, len(intSlice), cap(intSlice)) //len is length, cap is capacity
	intSlice = append(intSlice, 7)                                                         //append value to slice
	fmt.Printf("intSlice: %v, len: %d, cap: %d\n", intSlice, len(intSlice), cap(intSlice)) //len is length, cap is capacity

	var intSlice2 []int32 = []int32{8, 9}      //slice with initial length and capacity
	intSlice2 = append(intSlice, intSlice2...) //append values to slice
	fmt.Println(intSlice2)

	var intSlice3 []int32 = make([]int32, 5, 10)                                               //make a slice with length 5 and capacity 10
	fmt.Printf("intSlice3: %v, len: %d, cap: %d\n", intSlice3, len(intSlice3), cap(intSlice3)) //len is length

	var myMap map[string]uint8 = make(map[string]uint8) //create a map with string keys and uint8 values
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"Adam": 23, "Sarah": 45, "Charlie": 3} //create a map with initial values
	fmt.Println(myMap2["Adam"])                                          //access value by key
	fmt.Println(myMap2["Jason"])                                         //returns 0 if key does not exist, no error

	var age, ok = myMap2["Jason"] //check if key exists
	if ok {
		fmt.Printf("Jason's age is %v", age)
	} else {
		fmt.Println("Jason does not exist in the map")
	}

	delete(myMap2, "Charlie") //delete key from map
	fmt.Println(myMap2)

	for name := range myMap2 { //iterate over map keys
		fmt.Println(name)
	}
	for name, age := range myMap2 { //iterate over map keys and values
		fmt.Printf("%s is %d years old\n", name, age)
	}

	for i, v := range intArr {
		fmt.Printf("Index: %v, Value: %v \n", i, v)
	}
}
