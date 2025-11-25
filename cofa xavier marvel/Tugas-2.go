package main

import "fmt"

func main() {
	var duration, price int
	var vehicle string

	fmt.Scan(&vehicle, &duration)

	if duration == 0 {
		duration = 1
	}

	switch vehicle {
	case "Motorcycle", "motorcycle":
		price = 2000 * duration
	case "Car", "car":
		price = 5000 * duration
	case "Truck", "truck":
		price = 8000 * duration
	default:
		fmt.Println("Invalid vehicle type or parking duration")
	}
	fmt.Printf("Parking rates: Rp %d\n", price)
}
