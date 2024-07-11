package main

func main()  {
	var idMap map[string]int
	idMap = make(map[string]int)


	//definir con valores

	ipMap := map[string]int{
		"sebas":28
	}

	for key, val := range ipMap{
		fmt.Println(key, val)
	}
}