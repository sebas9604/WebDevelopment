package main
import (
	"fmt"
)
func main() {
	v := MyInt(3)
	fmt.Println(v.Double())
}

type MyInt int

func (mi MyInt) Double() int{
	return int(mi*2)
}