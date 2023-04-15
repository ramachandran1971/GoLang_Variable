package main
import "fmt"
var s = "Trichy"
func main() {
	fmt.Println(s)
	x := true
	if x {
		y := 1
		if x != false {
			fmt.Println(s)
			fmt.Println(x)
			fmt.Println(y)
		}
	}
	fmt.Println(x)
}
/*
Output
Trichy
trichy
true
1
true
*/
