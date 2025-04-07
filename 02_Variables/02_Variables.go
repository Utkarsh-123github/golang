package main

import "fmt"

func main() {
	var x, y int
	fmt.Println("x:", x, "y:", y) // Since we haven't given any predefined value, so it'll take value 0
	var a, b int = 53, 23
	fmt.Println("a:", a, "b:", b)
	const num int = 20
	// num = 30 this will give error since we cannot change constant values
	fmt.Println(num)
}
