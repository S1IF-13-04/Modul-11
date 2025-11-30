package main

import "fmt"

func main() {
	var nama string
	fmt.Scan(&nama)

	switch nama {
	case "nepenthes":
		fmt.Println("Termasuk tanaman karnivor \nAsli Indonesia")
	case "venus":
		fmt.Println("Termasuk Tanaman Karnivora \nBukan Asli Indonesia")
	default:
		fmt.Println("Bukan Termasuk tanaman Karnivora")
	}
}
