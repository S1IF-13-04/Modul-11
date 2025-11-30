package main
import "fmt"
func main() {
    var kendaraan string
    var durasi, tarif int
    fmt.Scan(&kendaraan)
    fmt.Scan(&durasi)

    switch {
    case durasi < 1:
        durasi = 1
    }

    switch {
    case kendaraan == "motor":
        tarif = 2000 * durasi
    case kendaraan == "mobil":
        tarif = 5000 * durasi
    case kendaraan == "truk":
        tarif = 8000 * durasi
    default:
        fmt.Println("jenis kendaraan tidak valid")
        return
    }
    fmt.Println("tarif parkir: Rp", tarif)
}