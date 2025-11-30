package main

import "fmt"

func main(){
	var jam,waktu  int
	fmt.Scan(&jam)

	switch{
	case jam == 0 :
		waktu = 12
		fmt.Printf("%d am", waktu)

	case jam < 12:
		fmt.Printf("%d am", jam)

	case jam == 12:
		fmt.Printf("%d pm", jam)	

	case jam > 12 && jam < 24:
		waktu = jam % 12
		fmt.Printf("%d pm", waktu)

	default:
		fmt.Printf("tidak valid")
	}
}