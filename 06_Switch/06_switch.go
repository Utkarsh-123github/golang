package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var agePaul, ageJohn int = rand.Intn(110), rand.Intn(110)
	fmt.Println("Paul age is:", agePaul)
	fmt.Println("John age is:", ageJohn)
	switch ageJohn {
	case 10:
		fmt.Println("John is 10 years old")
	case 20:
		fmt.Println("John is 20 years old")
	case 100:
		fmt.Println("John is 100 years old")
	default:
		fmt.Println("John has neither 10,20 nor 100 years old")
	}

	var ageSum int = ageJohn + agePaul
	switch ageSum {
	case 10:
		fmt.Println("ageJohn + agePaul = 10")
	case 20, 30, 40: //*\label{switchMulti}
		fmt.Println("ageJohn + agePaul = 20 or 30 or 40")
	case 2 * agePaul:
		fmt.Println("ageJohn + agePaul = 2 times agePaul")
	}

	switch {
	case agePaul > ageJohn:
		fmt.Println("agePaul > ageJohn")
	case agePaul == ageJohn:
		fmt.Println("agePaul == ageJohn")
	case agePaul < ageJohn:
		fmt.Println("agePaul < ageJohn")
	}

}
