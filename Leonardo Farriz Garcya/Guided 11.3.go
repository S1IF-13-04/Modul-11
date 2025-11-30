package main

import "fmt"

func main(){
	var jenis string
	var waktu,harga int
	fmt.Scan(&jenis, &waktu)

	switch jenis {
	case "motor":
		harga = 7000
		if waktu > 2{
			harga += 2000
		}
		fmt.Printf("tarif parkir: %d", harga)

	case "mobil":
		harga = 15000
		if waktu > 2{
			harga += 5000
		}
		fmt.Printf("tarif parkir: %d", harga)

	case "truk":
		harga = 25000
		if waktu > 2{
			harga += 10000
		}
		fmt.Printf("tarif parkir: %d", harga)

	default:
		fmt.Printf("tidak falid")
	}
}