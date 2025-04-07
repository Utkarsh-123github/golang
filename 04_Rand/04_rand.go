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
	fmt.Println("Available rooms:", roomsAvailable)
}
