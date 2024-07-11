package main
import (
	"fmt"
)
func main() {
	fmt.Println(getMax(1,2,3,4))
	vslice := []int{1, 5, 8, 6}
	fmt.Println(getMax(vslice...))

	fmt.Println("Deferred:")
	//deferred function
	defer fmt.Println("Bye!")
	fmt.Println("Hello")

}

func getMax(vals ...int) int{
	maxV := -1
	for _, v := range vals{
		if v > maxV {
			maxV = v
		}
	}

	return maxV
}

//deferred function
