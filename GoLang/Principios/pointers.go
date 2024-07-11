package main

import "fmt"

func main(){
	//& para la direccion y * para el valor

	var x int = 1
	var y int
	var ip *int //ip is pointer to int

	ip = &x // ip now points to x
	y = *ip //y is now 1


	//-----------------------

	ptr := new(int)
	*ptr = 3
}