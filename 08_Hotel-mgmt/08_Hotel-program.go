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
	// Note that we used Printf. This method will print a string like Println. Interestingly, you can inject variable/constant values inside a string. " Occupancy Rate: %0.2f %%\n" is called the “format specifier”. It’s a specification of the format. In this specification, we can add format verbs. A format verb is used to indicate to the printer that we want to inject in the printed output a value. The verb indicates how we want it to format the value.

	// %0.2f indicate that we want to print a float value (f) and that we want to display only two decimals after the decimal separator (0.2). Note that format verbs always begin with the sign %. Here are some useful format verbs :

	// %s for printing a string.

	// Ex : fmt.Printf("Hotel: %s", hotelName)
	// %d for printing an integer in base 10.

	// Ex : fmt.Printf("Number of rooms: %d", totalRooms)
	// If you want to print a percentage sign, the syntax is %%. Note also that to generate a line break added a \n at the end of our format specifier.
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
