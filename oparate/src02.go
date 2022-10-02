package main

import (
	"fmt"
)

const STOCK_OWNER= 2200
const TRADING_STOCK_RATE = 35.00
const MARKET_VALUE = 250

func isCapableListed(so int,tsr float64,mv float64) string{
	var req_msg string
	if so >= STOCK_OWNER && tsr >= TRADING_STOCK_RATE && mv>=MARKET_VALUE {
		req_msg = "上場条件を満たしています"
	}else if so >= STOCK_OWNER || tsr >= TRADING_STOCK_RATE || mv>=MARKET_VALUE {
		req_msg="上場条件の一部を満たしています"
	}else{
		req_msg="上場条件を満たせていません"
	}
	return req_msg
} 

func main() {

	stock_owner := 3888
	trading_stock_rate := 34.99
	market_value := 20000

	fmt.Println(isCapableListed(stock_owner,trading_stock_rate,float64(market_value)))

	is_auth := true

	if is_auth == true {
		fmt.Println("適正ユーザです")
	}else{
		fmt.Println("不正ユーザです")
	}
	
}