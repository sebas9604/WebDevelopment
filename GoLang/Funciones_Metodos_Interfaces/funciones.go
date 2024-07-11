package main
import (
	"fmt"
)
func main() {
	PrintHello()
	add(5,6)
	y := foo(5)
	fmt.Println(y)

	a, b := foo2(3)
	fmt.Println(a,b)

	//call by reference
	x := 2
	ref(&x)
	fmt.Print(x)

	//array arguments
	ar := [3]int{1, 2, 3}
	fmt.Println(arrs(ar))

	//array pointers
	ap := [3]int{1, 2, 3}
	arrp(&ap)
	fmt.Println(ap)

	//slices instead
	si := []int{1, 2, 3}
	sli(si)
	fmt.Println(si)
}

func PrintHello(){
	fmt.Println("Hello World")
}

func add(x , y int){
	fmt.Println(x+y)
}

func foo(x int)int{
	return x + 1
}

func foo2(x int)(int, int){
	return x, x + 1
}

//call by reference
func ref(y *int){
	*y = *y + 1
}

//array arguments
func arrs(x [3]int)int{
	return x[0]
}

//array pointers

func arrp(x *[3]int) {
	(*x)[0] = (*x)[0]+1
}

//slices instead
func sli(sli []int){
	sli[0] = sli[0] + 1
}