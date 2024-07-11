package main

import "fmt"

func main()  {
	
	type Person struct{
		name string
		addr string
		phone string	
	}

	p1 := Person(name: "Sebas", addr: "Piñu", phone: "005")

	barr, err := json.Marshal(p1)
}