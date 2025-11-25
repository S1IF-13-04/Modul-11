package main

import "fmt"

func main() {
	var kendaraan string
	var durasi int
	fmt.Print("Masukkan jenis kendaraan (Motor/mobil/truk) :")
	fmt.Scan(&kendaraan)
	fmt.Print("Masukkan durasi parkir (dalam jam) :")
	fmt.Scan(&durasi)

	switch kendaraan {
	case "Motor":
		if durasi >= 1 && durasi <= 2 {
			fmt.Println("Tarif parkir : Rp 7000")
		} else if durasi > 2 {
			fmt.Println("Tarif parkir : Rp 9000")
		}
	case "Mobil":
		if durasi >= 1 && durasi <= 2 {
			fmt.Println("Tarif parkir : Rp 15.000")
		} else if durasi > 2 {
			fmt.Println("Tarif parkir : Rp 20.000")
		}
	case "Truk":
		if durasi >= 1 && durasi <= 2 {
			fmt.Println("Tarif parkir : Rp 25.000")
		} else if durasi > 2 {
			fmt.Println("Tarif parkir : Rp 35.000")
		}
	default:
		fmt.Println("Tarif parkir : Rp 0")
	}
}
