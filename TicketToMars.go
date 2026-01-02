// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"math/rand"
)

func main() {

	//fmt.Println("Spaceline Days Trip type Price")
	fmt.Printf("%-22v %5v %5v %5v\n", "Spaceline", "Days", "Trip type", "Price")
	fmt.Println("==============================")
	for count := 10; count > 0; count-- {
		var name = getSpaceFlightName()
		var speed = getSpeed()
		var days = getDurationInDays(speed)
		var priceMap = getPriceMap()
		var tripType = "Round trip"
		var price = priceMap[speed]

		fmt.Printf("%-20v %5v %5v $%5v\n", name, days, tripType, price)
	}
}

// For each ticket, randomly select one of the following spacelines: Space Adventures, SpaceX, or Virgin Galactic.
func getSpaceFlightName() string {
	flightName := [3]string{"Space Adventures", "SpaceX", "Virgin Galactic"}
	idx := rand.Intn(3)
	return flightName[idx]
}

// Use October 13, 2020 as the departure date for all tickets. Mars will be 62,100,000 km away from Earth at the time.
// Randomly choose the speed the ship will travel, from 16 to 30 km/s.
func getSpeed() int {
	min := 16
	max := 30
	return rand.Intn(max-min+1) + min
}

// This will determine the duration for the trip to Mars and also the ticket price. Make faster ships more expen-
// sive, ranging in price from $36 million to $50 million. Double the price for round trips.
func getDurationInDays(speed int) int {
	var distance = 62100000
	var seconds = distance / speed
	var minutes = seconds / 60
	var hours = minutes / 60
	var days = hours / 24
	return days
}

func getPriceMap() map[int]int {
	myMap := make(map[int]int, 15)

	// Use a loop to populate the map.
	// The loop runs 15 times (i from 0 to 14).
	for i := 0; i < 15; i++ {
		// Key: 16 + i (ranges from 16 to 30)
		key := 16 + i
		// Value: 36 + i (ranges from 36 to 50)
		value := 36 + i

		myMap[key] = value
	}
	return myMap
}
