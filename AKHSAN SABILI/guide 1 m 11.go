package main
import "fmt"
func main(){
	var n,t int
	var jam string
	fmt.Scan(&n)
	switch {
	case n == 0:
		t = 12
		jam = "AM"
	case n < 12:
		t = n
		jam = "AM"
	case n > 12:
		t = n - 12
		jam = "PM"
	case n == 12:
		t = 12
		jam = "PM"
	}
	fmt.Println(t, jam)
}