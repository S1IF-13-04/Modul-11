package main

import (
	"fmt"
	"math"
)

func main() {
	var bil, bil2, hasil, ganjil, genap, kelipatan5, kelipatan10 int
	fmt.Scan(&bil)
	ganjil = (bil - 1) % 2
	genap = bil % 2
	kelipatan5 = bil % 5
	kelipatan10 = bil % 10

	switch {
	case kelipatan10 == 0:
		hasil = bil / 10
		fmt.Println("Kategori: Bilangan Kelipatan 10")
		fmt.Println("Hasil penjumlahan dengan bilangan berikutnya ", bil, "/ 10 =", hasil)
	case kelipatan5 == 0 && bil != 5:
		var hasilTemp = float64(hasil)
		hasilTemp = math.Pow(float64(bil), 2)
		fmt.Println("Kategori: Bilangan Kelipatan 5")
		fmt.Println("Hasil penjumlahan dengan bilangan berikutnya ", bil, "^ 2 =", hasilTemp)
	case ganjil == 0:
		bil2 = bil + 1
		hasil = bil + bil2
		fmt.Println("Kategori: Bilangan Ganjil")
		fmt.Println("Hasil penjumlahan dengan bilangan berikutnya ", bil, "+", bil2, "=", hasil)
	case genap == 0:
		bil2 = bil + 1
		hasil = bil * bil2
		fmt.Println("Kategori: Bilangan Genap")
		fmt.Println("Hasil penjumlahan dengan bilangan berikutnya ", bil, "*", bil2, "=", hasil)
	}
}
