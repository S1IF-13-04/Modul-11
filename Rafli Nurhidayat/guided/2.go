package main

import "fmt"

func main() {
	var tumbuhan string

	fmt.Scan(&tumbuhan)

	switch tumbuhan {
		case "nepenthes" :
			fmt.Println("Termasuk Tanamana Karnivora")
			fmt.Println("Asli Indonesia")
		case "venus" :
			fmt.Println("Termasuk Tanamana Karnivora")
			fmt.Println("Bukan Asli Indonesia")
		default :
			fmt.Println("Tidak Termasuk Tanamana Karnivora")
	}
}