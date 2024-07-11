package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main()  {
	 x := make([]int, 0, 3)
	var cond string
//	i := 0
	 for {
		fmt.Print("Ingrese un valor:")
		fmt.Scan(&cond)
		if cond == "X"{
			break
		}
		num, err := strconv.Atoi(cond)
        if err != nil {
            fmt.Println("Entrada inválida. Por favor, introduce un entero.")
            continue
        }
		
		x = append(x,num)

		sort.Ints(x)

		fmt.Print(x)
		
	 }

	
}