package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	switch {
	case n%10 == 0:
		fmt.Printf("Kategori: Bilangan Kelipatan 10\nHasil pembagian antara %d / 10 = %d\n", n, n/10)
	case n%5 == 0:
		fmt.Printf("Kategori: Bilangan Kelipatan 5\nHasil kuadrat dari %d ^2 = %d\n", n, n*n)
	case n%2 == 0:
		fmt.Printf("Kategori: Bilangan Genap\nHasil perkalian dengan bilangan berikutnya %d * %d = %d\n", n, n+1, n*(n+1))
	default:
		fmt.Printf("Kategori: Bilangan Ganjil!\nHasil penjumlahan dengan bilangan berikutnya %d + %d = %d\n", n, n+1, n+n+1)
	}
}
