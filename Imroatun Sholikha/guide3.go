package main

import "fmt"

func main() {
	var jenis string
	var durasi int
	var tarif int

	fmt.Print("Masukan jenis: ")
	fmt.Scan(&jenis)
	fmt.Print("Durasi: ")
	fmt.Scan(&durasi)

	switch jenis {
	case "motor":
		switch {
		case durasi <= 2:
			tarif = 7000
		default:
			tarif = 9000
		}

	case "mobil":
		switch {
		case durasi <= 2:
			tarif = 15000
		default:
			tarif = 20000
		}

	case "truk":
		switch {
		case durasi <= 2:
			tarif = 25000
		default:
			tarif = 35000
		}

	default:
		tarif = 0
	}

	fmt.Println("Tarif Kendaraan", tarif)
}
