package main

import (
	"fmt"
	// "math"
)

func main() {
	arrayFloat := [] float32{0,0,0}

	arrayFloat[0] = 4.29
	arrayFloat[1] = 3.3

	var arrayString []string = [] string{"東北電力","七十七銀行","TTK","仙台市営地下鉄","JR東日本仙台支社","やまや","サトー商会"}
	
	sliceStr := arrayString[:]
	fmt.Println(arrayFloat)
	fmt.Println(sliceStr)
}