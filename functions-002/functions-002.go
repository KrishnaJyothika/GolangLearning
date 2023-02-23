package main

import "fmt"

func add(x, y int) int {

	out := x + y
	return out
}

// func calc(x, y int) (int, int) {
// 	out1 := x + y
// 	out2 := x - y
// 	return out1, out2
// }

func calc(x, y int) (out1, out2 int) {
	out1 = x + y
	out2 = x - y
	return
}

func main() {

	num1 := 4
	num2 := 3

	result := add(num1, num2)
	fmt.Println(result)

	result1, result2 := calc(num1, num2)
	fmt.Println(result1, result2)

}
