package main

import "fmt"

func main() {
	var kendaraan string
	var motor, mobil, truk, durasi int
	fmt.Scan(&kendaraan, &durasi)
	
	motor = durasi * 2000
	mobil = durasi * 5000
	truk = durasi * 8000

	switch kendaraan {
	case "motor":
		fmt.Println("Rp", motor)
	case "mobil":
		fmt.Println("Rp", mobil)
	case "truk":
		fmt.Println("Rp", truk)
	}
}
