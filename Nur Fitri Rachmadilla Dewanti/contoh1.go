package main
import "fmt"
func main() {
	var jam int
	fmt.Scan(&jam)

	switch {
	case jam == 0:
		fmt.Println("12 AM")
	case jam == 12:
		fmt.Println("12 PM")
	case jam > 0 && jam < 12:
		fmt.Printf("%d AM\n", jam)
	case jam > 12 && jam <= 23:
		fmt.Printf("%d PM\n", jam-12)
	default:
		fmt.Println("jam tdk valid")				
	}
}