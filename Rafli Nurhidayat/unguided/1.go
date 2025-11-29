package main

import "fmt"

func main() {
	var kadar float64

	fmt.Scan(&kadar)

	switch {
		case kadar >= 6.5 && kadar <= 8.6 :
			fmt.Print("Air Layak Minum")
		case kadar < 6.5 && kadar > 0 || kadar > 8.6 && kadar <= 14 :
			fmt.Print("Air Tidak Layak Minum")
		case kadar < 0 || kadar > 14 :
			fmt.Print("Nilai pH tidak valid. Nilai pH harus antara 0 dan 14.")
		default :
			fmt.Print("Nilai pH tidak valid.")
	}
}