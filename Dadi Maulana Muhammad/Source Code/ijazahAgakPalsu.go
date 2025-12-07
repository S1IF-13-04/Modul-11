package main

import "fmt"

func main() {
	var j string
	var d, t int

	fmt.Scan(&j, &d)
	if d < 1 {
		d = 1
	}

	switch j {
	case "motor":
		t = 2000
	case "mobil":
		t = 5000
	case "truk":
		t = 8000
	default:
		fmt.Println("Jenis kendaraan tidak valid")
		return
	}

	fmt.Println("Rp", t*d)
}
