package main
import "fmt"
import "strconv"
func main()  {
	var appleNum int
	fmt.Printf("Number of apples?")
	fmt.Scan(&appleNum)
	fmt.Printf(strconv.Itoa(appleNum))
	
}