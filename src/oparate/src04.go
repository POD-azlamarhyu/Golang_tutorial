package main

import (
	"fmt"
)

const STOCK_OWNER= 2200
const TRADING_STOCK_RATE = 35.00
const MARKET_VALUE = 250

func isCapableListed(so int,tsr float64,mv float64,cn string) string{
	var req_msg string
	if so >= STOCK_OWNER && tsr >= TRADING_STOCK_RATE && mv>=MARKET_VALUE {
		req_msg = cn+"社は上場条件を満たしています"
	}else if so >= STOCK_OWNER || tsr >= TRADING_STOCK_RATE || mv>=MARKET_VALUE {
		req_msg=cn+"社は上場条件の一部を満たしています"
	}else{
		req_msg=cn+"社は上場条件を満たせていません"
	}
	return req_msg
} 

func main() {

	defer fmt.Println(isCapableListed(10000,29.89,float64(40000),"NTTT"))
	fmt.Println(isCapableListed(1000,2.89,float64(50),"MURATA"))
	defer fmt.Println(isCapableListed(10000,50.89,float64(220000),"SAICO-MART"))
}