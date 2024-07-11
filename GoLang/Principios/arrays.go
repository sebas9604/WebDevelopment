package main

import (
	"fmt"
)

func main()  {
	var x [5]int
	x[0] = 2
	fmt.Print(x[1])


	//Array literal
	var y = [5]int{1, 2, 3, 4, 5}
	fmt.Print(y)
	z := [...]int{1, 2, 3, 4}
	fmt.Print(z)

fmt.Println("")
	for i, v := range y{
		fmt.Printf("\nind %d, val %d", i, v)
	}
}