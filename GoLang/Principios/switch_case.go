package main
import (
	"fmt"
)

func main()  {
	switch x {
	case 1:
		fmt.Printf("Case 1")

	case 2:
		fmt.Printf("Case 2")

	default:
		fmt.Printf("No Case")
	}


	//Tagless switch

switch {
	case x > 1:
		fmt.Printf("Case 1")
	case x < -1:
		fmt.Printf("Case 2")
	default:
		fmt.Printf("No Case")

}
}

