package main
//import "fmt"
/*
import(
"fmt"
"strconv"
"strings"
)

func main(){/*
	var name string = "Sebas"
	edad:= 28
	fmt.Printf("Hello, " + strings.ToUpper(name) + " "+ strconv.Itoa(edad))

	type ages int

/*	var edad ages = 28
	fmt.Printf("Hello, " + name + " you have " + edad + "years old")*/

	
/*		i, _ := strconv.Atoi("10")
		y := i * 2
		fmt.Println(y)


		s := strings.Replace("ianianian", "ni", "in", 2)
		fmt.Println(s)
	
		
		x:=7
		switch {
		  case x>3:
			fmt.Printf("1")
		  case x>5:
			fmt.Printf("2")
		  case x==7:
			fmt.Printf("3")
		  default: 
			fmt.Printf("4")
		}

		test()*/

		var x int
		var y *int
		z := 3
		y = &z
		x = &y
}

func test(){

	var xtemp int
  x1 := 0
  x2 := 1
  for x:=0; x<5; x++ {
    xtemp = x2
    x2 = x2 + x1
    x1 = xtemp
  }
  fmt.Println(x2)
}

