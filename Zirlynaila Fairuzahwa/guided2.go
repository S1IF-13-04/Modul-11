package main
import "fmt"

func main(){
	var tanaman string
	fmt.Scan(&tanaman)
	switch tanaman {
	case "nephentes", "drosera":
		fmt.Println("Termasuk Tanaman Karnivora")
		fmt.Println("Asli Indonesia")
	case "venus", "sarracenia":
		fmt.Println("Termasuk Tanaman Karnivora")
		fmt.Println("Bukan Asli Indonesia")
	default:
		fmt.Print("Tidak Termasuk Tanaman Karnivora")
	}
}