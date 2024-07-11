package main

import (
	"fmt"
//	"strconv"
	"strings"
)

func main()  {
	fmt.Println("Write a word:")
	var word string
	fmt.Scan(&word)
	fmt.Println(word)
	word = strings.TrimSpace(word)
	//word = strings.ToLower(strings.ReplaceAll(word, " ", ""))
fmt.Println(word)
		if strings.HasPrefix(word, "i") && strings.Contains(word, "a") && strings.HasSuffix(word,"n"){
			fmt.Print("Encontrado")
		}else{
			fmt.Print("No Encontrado")
		}
	
}