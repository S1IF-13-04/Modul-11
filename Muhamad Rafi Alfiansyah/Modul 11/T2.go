package main
import "fmt"
func main(){
	var j,jam string
	var w,t int
	fmt.Scan(&j,&w,&jam)
	switch j{
	case "motor":
		t = w * 2000
	case "mobil":
		t = w * 5000
	case "truk":
		t = w * 8000
	} 
	fmt.Println("Rp",t)
}