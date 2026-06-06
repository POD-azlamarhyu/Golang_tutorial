package main

import (
	"fmt"
	// "math"
)

func main() {
	arrayFloat := [] float32{0,0,0}

	arrayFloat[0] = 4.29
	arrayFloat[1] = 3.3

	var arrayString []string = [] string{"札幌", "仙台", "東京", "名古屋", "大阪", "広島", "福岡"}
	
	sliceStr := arrayString[:]
	fmt.Println(arrayFloat)
	fmt.Println(sliceStr)
}