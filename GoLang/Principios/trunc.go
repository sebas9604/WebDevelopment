package main

import (
	"fmt"
	"strconv"
)

func main()  {
	var num string
	fmt.Printf("Introduzca un numero de punto flotante: ")
	fmt.Scanln(&num)

	floatNum, err := strconv.ParseFloat(num, 64)
	if err != nil {
		fmt.Println("Error: por favor, introduce un número válido.")
		return
	}
	intNum := int(floatNum)

	fmt.Printf("La versión truncada del número es: %d\n", intNum)
}

/*
package main

import "fmt"

func main() {
	var flt float32
	fmt.Printf("Enter a floating point number.\n")
	fmt.Scan(&flt)
	var i int = int(flt)
	fmt.Print(i)
}
*/