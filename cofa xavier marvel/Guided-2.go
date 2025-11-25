package main

import "fmt"

func main() {
	var plant_name string
	fmt.Scan(&plant_name)
	switch plant_name {
	case "nepenthes", "drosera":
		fmt.Println("Including Carnivorous Plants.")
		fmt.Println("Native to Indonesia.")
	case "venus", "sarracenia":
		fmt.Println("Including Carnivorous Plants.")
		fmt.Println("Not Native to Indonesia.")
	default:
		fmt.Println("Excluding Carnivorous Plants.")
	}
}
