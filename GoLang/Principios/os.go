package main
import (
	"fmt"
	"os"
)
func main() {
	f, err := os.Open("/test.txt")
	if err == nil {fmt.Println("Error")}
	barr := make([]byte, 10)
	nb, err := f.Read(barr)
	if err == nil {fmt.Println("Error")}
	fmt.Println(nb)
	f.Close()
}