package main

import "fmt"

func main() {
	var tanaman string
	fmt.Scan(&tanaman)

	switch tanaman {
	case "nepenthes":
		fmt.Println("Termasuk Tanaman Karnivora \nAsli Indonesia")
	case "venus":
		fmt.Println("Termasuk Tanaman Karnivora \nBukan Asli Indonesia")
	default:
		fmt.Println("Tidak Termasuk Tanaman Karnivora")
	}
}
