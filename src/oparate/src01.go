package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main()  {
	sum := 0

	for i := 0; i <= 10; i++{
		sum += int(math.Pow(float64(i),2)) 
	}

	fmt.Println(sum)

	const Number float64 = 23.41

	init_number := 3.1
	count := 0

	for init_number < Number {
		init_number += rand.Float64()
		fmt.Println(init_number)
		count ++
	}

	fmt.Println(count)
		
}