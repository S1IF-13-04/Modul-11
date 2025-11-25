package main

import "fmt"

func main() {
	var vehicle string
	var duration int
	var price int
	fmt.Print("Enter the type of vehicle (Motorcycle/Car/Truck): ")
	fmt.Scan(&vehicle)
	fmt.Print("Enter the parking duration (in hours): ")
	fmt.Scan(&duration)
	switch {
	case vehicle == "Motorcycle" && duration >= 1 && duration <= 2:
		price = 7000
	case vehicle == "Motorcycle" && duration > 2:
		price = 9000
	case vehicle == "Car" && duration >= 1 && duration <= 2:
		price = 15000
	case vehicle == "Car" && duration > 2:
		price = 20000
	case vehicle == "Truck" && duration >= 1 && duration <= 2:
		price = 25000
	case vehicle == "Truck" && duration > 2:
		price = 35000
	default:
		fmt.Println("Invalid vehicle type or parking duration")
	}
	fmt.Printf("Parking rates: Rp %d\n", price)
}
