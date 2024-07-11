package main
import (
	"fmt"
	"encoding/json"
	"os"
	"bufio"
)
func main() {
	var nombre string
	var direccion string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Ingrese su nombre:")
	scanner.Scan()
	nombre = scanner.Text()
	fmt.Println("Ingrese una direccion:")
	scanner.Scan()
	direccion = scanner.Text()


	var mapa = map[string]string{"nombre": nombre, "direccion": direccion}

	
	jsonData, err := json.Marshal(mapa)
	if err != nil{
		fmt.Println("Error en el marshal", err)
		return
	}

	fmt.Println(string(jsonData))

}