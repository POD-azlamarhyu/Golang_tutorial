package main

import (
	"fmt"
	"math"
)

func getCircleSurface(x float64) float64{
	return x*math.Pi
}

func main(){
	fmt.Println(getCircleSurface(4.0))
}