package main

import (
	"fmt"
	"strconv"
)
func main() {
	x := make([]int, 0, 10)
	x = fill_array(x)
	//fmt.Println(x)
	BubbleSort(&x)
	fmt.Println("Ordenado\n", x)

	Swap(x, 2)
	fmt.Println("Swap\n", x)

}

func Swap(slice []int, i int) {
    if i < 0 || i >= len(slice)-1 {
        fmt.Println("Índice fuera de rango.")
        return
    }
    // Intercambiar los elementos
    slice[i], slice[i+1] = slice[i+1], slice[i]
}

func BubbleSort(x *[]int)  {
	var aux int
	for i := 0; i < 10; i++{
		for j := 0; j < 10; j++{
			if (*x)[i] < (*x)[j]{
				aux = (*x)[i]
				(*x)[i] = (*x)[j]
				(*x)[j] = aux
			}
		}
	}
}

func fill_array(x []int)[]int{
	var nume string
	for i:=0;i<10;i++{
			
		fmt.Print("Ingrese un valor:")
		fmt.Scan(&nume)
		num, err := strconv.Atoi(nume)
        if err != nil {
            fmt.Println("Entrada inválida. Por favor, introduce un entero.")
            continue
        }
		
		x = append(x, num)

		fmt.Println(x)
		
		
	 }
	 	return x
}