package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var hotelName string = "The Ramadas"
	fmt.Println("Hotel:", hotelName)
	var totalRooms int = 100
	var roomsAvailable int = rand.Intn(100)
	fmt.Println("Total rooms:", totalRooms)
	if roomsAvailable > 0 {
		fmt.Println("Available rooms:", roomsAvailable)
	} else {
		fmt.Println("No rooms available!!")
	}
}
