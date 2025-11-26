package main

import "fmt"

func main() {
	var jenis string
	var jam int
	fmt.Scanln(&jenis)
	fmt.Scanln(&jam)

	switch jenis {
	case "motor":
		fmt.Print("Rp ", jam*2000)
	case "mobil":
		fmt.Print("Rp ", jam*5000)
	case "truk":
		fmt.Print("Rp ", jam*8000)
	}
}
