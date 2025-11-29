package main

import "fmt"

func main() {
	var x int

	fmt.Scan(&x)

	switch {
	case x%10 == 0:
		fmt.Print("Kategori: Kelipatan 10\nHasil pembagian antara ", x, "dengan 10 adalah ", x/10)
	case x%5 == 0 && x != 5:
		fmt.Print("Kategori: Kelipatan 5\nHasil kuadrat dari ", x, " ^2 adalah ", x*x)
	case x%2 == 0:
		fmt.Print("Kategori: Bilangan Genap\nHasil perkalian dengan bilangan berikutnya adalah ", x*(x+1))
	default:
		fmt.Print("Kategori: Bilangan Ganjil\nHasil penjumlahan dengan bilangan berikutnya adalah ", x+(x+1))
	}
}
