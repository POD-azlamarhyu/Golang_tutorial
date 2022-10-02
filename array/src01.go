package main

import (
	"fmt"
	"math/rand"
)

var globalArray[3] int = [3]int {1,3,2} 

func main()  {
	var size int = 10
	var arrayInt[10] int

	arrayString := [...] string{"JR北海道","ニトリ","セイコーマート","北海道拓殖銀行","雪印","サッポロビール","ツルハホールディングス","札幌証券取引所"}

	fmt.Println(globalArray)
	
	for i := 0; i< size; i++{
		arrayInt[i] = rand.Intn(100)
	}

	fmt.Println(arrayInt)
	for i := 0; i< len(arrayString); i++{
		fmt.Println(arrayString[i])
	}
}

