package main

import (
	"fmt"
	"math/rand"
)

func main() {
	n := 10
	var varInt[] int = make([]int,10,10)
	varStr := make([]string,10,10)
	varFloat := make([]float32, n,n)

	fmt.Println(varStr)
	for i := 0;i<n; i++ {
		varInt[i] = rand.Intn(10) * rand.Intn(10)
	}

	fmt.Println(varInt)

	for i := 0;i<n; i++{
		varFloat[i] = float32(rand.Float64() * float64(rand.Intn(10)))
	}

	fmt.Println(varFloat)
}