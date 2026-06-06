package main

import (
	"fmt"
)

func main() {
	var var_int int = 1000
	var var_str string = "Unko"

	var pointer_int *int = &var_int
	var pointer_str *string = &var_str
	
	fmt.Println(var_int)
	fmt.Println(var_str)

	fmt.Println(pointer_int)
	fmt.Println(pointer_str)

	fmt.Println(*pointer_int)
	fmt.Println(*pointer_str)

	fmt.Println(*pointer_int + 100)
	fmt.Println(*pointer_int - 100)
}
