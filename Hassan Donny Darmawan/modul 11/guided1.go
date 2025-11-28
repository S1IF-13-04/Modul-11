package main 
import "fmt"

func main (){
	var jam_24 int
	var ampm string
	fmt.Scan(&jam_24)
	
	switch {
	case jam_24==0:
		jam_24=12
		ampm = "am"
	case jam_24 < 12:
		ampm="am"
	case jam_24==12:
		ampm="pm"
	case jam_24>=24:
		ampm = "Jam yang dimasukan tidak valid"
	default:
		jam_24 = jam_24-12
		ampm = "pm"
	}
	fmt.Println(jam_24,ampm) 
}