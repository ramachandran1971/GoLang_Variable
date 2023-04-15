package main
import (
	"fmt"
	"reflect"
)

func main() {
	var i = 10
	var s = "Trichy"

	fmt.Println(reflect.TypeOf(i))
	fmt.Println(reflect.TypeOf(s))
}
/*
Output
int
string
*/
