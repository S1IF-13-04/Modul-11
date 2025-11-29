package main
import "fmt"
func main(){
	var b int
	fmt.Scan(&b)
	switch{
	case b % 10 ==0:
		fmt.Println("Kategori : Bilangan Kelipatan 10")
		fmt.Printf("Hasil pembagian antara %d / 10 = %d",b, b/10)
	case b % 5 == 0 && b != 5:
		fmt.Println("Kategori : Bilangan Kelipatan 5")
		fmt.Printf("Hasil kuadrat dari %d ^2 = %d",b, b*b)
	case b % 2 == 1:
		fmt.Println("Kategori : Bilangan Ganjil")
		fmt.Printf("Hasil penjumlahan dengan bilangan berikutnya %d + %d = %d",b, b+1, b+(b+1))
	case b % 2 == 0:
		fmt.Println("Kategori : Bilangan Genap")
		fmt.Printf("Hasil perkalian dengan bilangan berikutnya %d * %d = %d",b, b+1, b*(b+1))
	}
}