package main

import "fmt"
import "strconv"
func main(){

	for i:=0;i<10;i++{
		if i == 5 {break}
		if i == 3 {continue}
	fmt.Printf(strconv.Itoa(i))
	}
}


