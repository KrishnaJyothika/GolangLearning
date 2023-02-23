package main

import "fmt"

var a = 9 //Package Varaible

func demo() {
	var a = 8 //Function Variable
	fmt.Println("In Demo()", a)
}

func main() {

	demo()
	fmt.Println("In Main()", a)
}
