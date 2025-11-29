package main

import "fmt"

func main() {
	var jenis string
	var durasi int
	fmt.Print("input jenis kendaraan:")
	fmt.Scan(&jenis)
	fmt.Print("input durasi parkir:")
	fmt.Scan(&durasi)

	switch {
	case jenis == "motor" && durasi >= 1 && durasi <= 2:
		fmt.Print("Rp7.000")
	case jenis == "motor" && durasi > 2:
		fmt.Print("Rp9.000")
	case jenis == "mobil" && durasi >= 1 && durasi <= 2:
		fmt.Print("Rp15.000")
	case jenis == "mobil" && durasi > 2:
		fmt.Print("Rp20.000")
	case jenis == "truk" && durasi >= 1 && durasi <= 2:
		fmt.Print("Rp25.000")
	case jenis == "truk" && durasi > 2:
		fmt.Print("Rp35.000")
	default:
		fmt.Print("Jenis kendaraan atau durasi parkir tidak valid")
	}
}
