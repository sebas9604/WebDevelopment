package main
import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)
func main() {
	fmt.Println("Introduzca los valores respectivamente: aceleracion, velocidad inicial, desplazamiento inicial y tiempo")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var input string = scanner.Text()
	aceleracion, err := strconv.ParseFloat(input,64)
	if err != nil {
        fmt.Println("Error al convertir la entrada a float:", err)
        return
    }
	scanner.Scan()
	input = scanner.Text()
	velocidad_inicial, err := strconv.ParseFloat(input,64)
	if err != nil {
        fmt.Println("Error al convertir la entrada a float:", err)
        return
    }	
	scanner.Scan()
	input = scanner.Text()

	desplazamiento_inicial, err := strconv.ParseFloat(input,64)
	if err != nil {
        fmt.Println("Error al convertir la entrada a float:", err)
        return
    }	
	scanner.Scan()
	input = scanner.Text()

	tiempo, err := strconv.ParseFloat(input,64)
	if err != nil {
        fmt.Println("Error al convertir la entrada a float:", err)
        return
    }	
	fmt.Println(aceleracion, velocidad_inicial, desplazamiento_inicial, tiempo)

	fn := GenDisplaceFn(aceleracion, velocidad_inicial, desplazamiento_inicial)

	fmt.Println("Desplazamiento final: ",fn(tiempo))
}

func GenDisplaceFn(a, v, d float64) func(float64) float64  {
	return func(t float64) float64{
		s := (0.5*(exp(a, t))+v+d)
		return (s)
	}
}

func exp(a, t float64) float64{
	result := 1.0
    for t > 0 {
        result *= a
        t--
    }

	return result
}

//s = ½ a t2 + vo  + so