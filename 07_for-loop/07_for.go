package main

import (
	"fmt"
)

func main() {
	var sum int = 0
	fmt.Println("Initial sum:", sum)
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("Final sum:", sum)

}
