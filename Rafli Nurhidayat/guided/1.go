package main

import "fmt"

func main() {
	var jam, jamKonversi int
	var waktu string

	fmt.Scan(&jam)

	switch {
	case jam == 0:
		jamKonversi = 12
		waktu = "AM"
	case jam < 12:
		jamKonversi = jam
		waktu = "AM"
	case jam == 12:
		jamKonversi = 12
		waktu = "PM"
	case jam > 12:
		jamKonversi = jam - 12
		waktu = "PM"
	}
	
	fmt.Println(jamKonversi, waktu)
}
