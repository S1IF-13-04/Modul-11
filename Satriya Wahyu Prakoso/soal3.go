package main
import "fmt"
func main() {
	var n int
	fmt.Scanln(&n)
	switch {
	case n%10 == 0:
		kelipatan10 := n / 10
		fmt.Print("Kategori: Bilangan Kelipatan 10\n")
		fmt.Printf("Hasil pembagian antara %d / 10 = %d\n", n, kelipatan10)
	case n%5 == 0 && n != 5:
		kelipatan5 := n * n
		fmt.Print("Kategori: Bilangan Kelipatan 5\n")
		fmt.Printf("Hasil kuadrat dari %d ^2 = %d\n", n, kelipatan5)
	case n%2 != 0:
		ganjil := n + n + 1
		fmt.Print("Kategori: Bilangan Ganjil\n")
		fmt.Printf("Hasil pertambahan dengan bilangan berikutnya %d + %d = %d\n",
			n, n+1, ganjil)
	case n%2 == 0:
		genap := n * (n + 1)
		fmt.Print("Kategori: Bilangan Genap\n")
		fmt.Printf("Hasil perkalian dengan bilangan berikutnya %d * %d = %d\n",
			n, n+1, genap)
	}
}