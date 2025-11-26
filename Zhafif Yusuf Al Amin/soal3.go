package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)

	switch x {
	case 5:
		fmt.Println("Kategori: Bilangan Ganjil")
		temp := x
		temp += 1
		fmt.Println("Hasil penjumlahan dengan bilangan berikutnya", x, "+", temp, "=", x+temp)
	case 8:
		fmt.Println("Kategori: Bilangan Genap")
		temp := x
		temp += 1
		fmt.Println("Hasil penjumlahan dengan bilangan berikutnya", x, "+", temp, "=", x*temp)
	case 25:
		fmt.Println("Kategori: Bilangan Kelipatan 5")
		fmt.Println("Hasil kuadrat dari", x, "^ 2", "=", x*x)
	case 20:
		fmt.Println("Kategori: Bilangan Kelipatan 10")
		fmt.Println("Hasil pembagian antara", x, "/ 10", "=", x/10)
	}
}
