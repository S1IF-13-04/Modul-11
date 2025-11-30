package main
import "fmt"
func main() {
    var n int
    fmt.Scan(&n)

    switch {
    case n%10 == 0:
        fmt.Println("kategori: bilangan kelipatan 10")
        fmt.Printf("hasil pembagian antara %d / 10 = %d\n", n, n/10)
    case n%5 == 0 && n != 5:
        fmt.Println("kategori: bilangan kelipatan 5")
        fmt.Printf("hasil kuadrat dari %d ^2 = %d\n", n, n*n)
    case n%2 == 0:
        fmt.Println("kategori: bilangan genap")
        fmt.Printf("hasil perkalian dengan bilangan berikutnya %d * %d = %d\n", n, n+1, n*(n+1))
    default:
        fmt.Println("kategori: bilangan ganjil")
        fmt.Printf("hasil penjumlahan dengan bilangan berikutnya %d + %d = %d\n", n, n+1, n+(n+1))
    }
}