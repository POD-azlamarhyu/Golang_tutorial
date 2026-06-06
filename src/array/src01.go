package main

import (
	"fmt"
	"math/rand"
)

var globalArray[3] int = [3]int {1,3,2} 

func main()  {
	var size int = 10
	var arrayInt[10] int

	arrayString := [...] string{"札幌", "仙台", "東京", "名古屋", "大阪", "広島", "福岡"}

	fmt.Println(globalArray)
	
	for i := 0; i< size; i++{
		arrayInt[i] = rand.Intn(100)
	}

	fmt.Println(arrayInt)
	for i := 0; i< len(arrayString); i++{
		fmt.Println(arrayString[i])
	}
}

