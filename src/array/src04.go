package main

import (
	"fmt"
	"math/rand"
)

func initArrayInt () *[10]int{
	var array[10] int
	for i := 0; i < len(array); i++ {
		array[i] = rand.Intn(10)
	}

	return &array
}

func main()  {
	// var arrayStr[10] string

	arrayInt := initArrayInt()

	fmt.Println(*arrayInt)
	for i:=0;i<len(arrayInt);i++{
		fmt.Println((*arrayInt)[i])
	}
}

