package main
import (
	"fmt"
)
//variables as functions
var funcVar func(int)int
	func incFn1(x int) int{
		return x + 1
	}

//functions as arguments
func applyIt(afunct func (int) int, val int) int{
	return afunct(val)
}
func incFn(x int)int{return x+1}
func decFn(x int)int{return x-1}


func main() {

	//variables as functions
	funcVar = incFn1
	fmt.Println(funcVar(1))

	//functions as arguments
fmt.Println(applyIt(incFn,2))
fmt.Println(applyIt(decFn,2))

	//Anonyomous function
	v := applyIt(func (x int)int{return x*x}, 2)
	fmt.Print(v)
}