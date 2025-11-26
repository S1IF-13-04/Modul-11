package main

import "fmt"

func main() {
	var tanaman string
	fmt.Scan(&tanaman)

	switch tanaman {
	case "nepenthes":
		fmt.Println("Termasuk tanaman karnivora")
		fmt.Println("Asli Indonesia")
	case "venus":
		fmt.Println("Termasuk tanaman karnivora")
		fmt.Println("Bukan asli Indonesia")
	default:
		fmt.Println("Bukan termasuk tanaman karnivora")
	}
}
