package main

import "fmt"

// Function to add two nums
func addTwoNums(x int, y int) int {
	return x + y
}

// Note : In Go a function can return any number of results unlike c++
// for example : here below I'll return two strings

func swapStrings(x, y string) (string, string) {
	return y, x
}

func main() {
	fmt.Println("The sum is:", addTwoNums(10, 20))
	var x string = "Hello"
	var y string = "World"
	fmt.Println("Initially x:", x, "y:", y)
	x, y = swapStrings(x, y)
	fmt.Println("Finally x:", x, "y:", y)

}
