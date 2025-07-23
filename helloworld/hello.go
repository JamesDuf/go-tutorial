package main //standalone executable

import "fmt"

func main() {
	fmt.Println("Hello, World!")

	name, age, likeGo := "Alice", 30, true
	fmt.Printf("Name: %s\nAge: %d\nLikes Go: %t\n", name, age, likeGo)

	//var year int = 2025 // variable declaration
	//var p *int = &year  // pointer to year

}

/*
func main() {
	println("Hello, World!")
}
*/

//to compile this code, use the command:
//go build hello.go
//./hello

//or you can run it directly without building
//go run hello.go

//link to remote packages on github
