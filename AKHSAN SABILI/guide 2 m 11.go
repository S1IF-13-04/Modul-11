package main
import "fmt"
func main(){
	var tanam string
	fmt.Scan(&tanam)
	switch tanam{
	case "nepenthes":
		fmt.Println("Termasuk tanaman karnivora asli indonesia")
	case "venus":
		fmt.Println("Termasuk tanaman karnivora tidak asli indonesia")
	default:
		fmt.Println("tidak termasuk tanaman karnivora")
	}
}