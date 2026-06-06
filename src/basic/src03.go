package main

import (
	"fmt"
	"math/rand"
	"math"
)

var val_bool bool
var val_int = 1000
var val_float = 9.81

var val_str string = "SDGsを進めよう!"

func main() {
	var val_loc_str string = "スタンダード & プアーズ 500"
	const PI = math.Pi 

	y := 4*rand.Intn(3) + 3*rand.Intn(10)

	
	fmt.Printf("%v %v %v %q\n",val_bool,val_int,val_float,val_loc_str)
	fmt.Printf("y = %v\n",y)

	z := math.Cos(PI)
	fmt.Printf("z = %v\n",z)
}