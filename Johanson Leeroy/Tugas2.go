package main

import "fmt"

func main() {
	var kendaraan, j string
	var jam, biaya int

	fmt.Scan(&kendaraan, &jam, &j)
	if jam == 0 {
		jam = 1
	}
	switch kendaraan {
	case "motor":
		biaya = 2000 * jam
	case "mobil":
		biaya = 5000 * jam
	case "truk":
		biaya = 8000 * jam
	}
	fmt.Println("Rp", biaya)
}
