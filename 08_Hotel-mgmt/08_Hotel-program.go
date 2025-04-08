// We are making a program that looks like the attached screenshot in the folder
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	const hotelName = "The Ramadas"
	const totalRooms = 110
	const startingRoomNumber = 101
	roomsOccupied := rand.Intn(totalRooms)
	roomsAvailable := totalRooms - roomsOccupied
	occupancyRate := (float64(roomsOccupied) / float64(totalRooms)) * 100
	var occupancyLevel string
	if occupancyRate > 70 {
		occupancyLevel = "High"
	} else if occupancyRate > 20 {
		occupancyLevel = "Medium"
	} else {
		occupancyLevel = "Low"
	}

	fmt.Println("Hotel:", hotelName)
	fmt.Println("Total Number of Rooms:", totalRooms)
	fmt.Println("Rooms Available:", roomsAvailable, "      Occupancy level:", occupancyLevel)
	fmt.Printf("                          Occupancy Rate: %0.2f %%\n", occupancyRate)
	if roomsAvailable > 0 {
		fmt.Println("Rooms:")
		for i := 0; i < roomsAvailable; i++ {
			roomNumber := startingRoomNumber + i
			size := rand.Intn(6) + 1
			nights := rand.Intn(10) + 1
			fmt.Println(roomNumber, ":", size, "people /", nights, " nights ")
		}
	} else {
		fmt.Println("No rooms available for tonight")
	}

}
