package main

import "fmt"

func main() {
	var ph float64
	fmt.Scan(&ph)

	switch {
	case ph >= 6.5 && ph <= 8.6:
		fmt.Println("Drinkable Water")
	case ph < 6.5 && ph > 8.6:
		fmt.Println("Undrinkable Water")
	case ph < 0 || ph > 14:
		fmt.Println("Invalid input, pH range 0 - 14")
	default:
		fmt.Println("Water not fit for drinking")
	}
}
