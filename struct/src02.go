package main

import (
	"fmt"
)

type BankAccount struct{
	bank_code string
	bank_name string
	branch_code int
	branch string
	account_code int
	name string
}

func initAccount(bank_code string,bank_name string,branch string,account_code int) *BankAccount{
	account := new(BankAccount)
	account.bank_code = bank_code
	account.bank_name = bank_name
	account.branch = branch
	account.account_code = account_code

	return account
}

func (p BankAccount) inputPin(pin string) string{
	var msg string
	if pin == "3199" {
		msg = "正しいPINです"
	}else{
		msg = "PINが間違っています"
	}
	return msg
}

func main() {
	var jiro BankAccount

	jiro.bank_code = "8391"
	jiro.bank_name = "四菱UFO銀行"

	fmt.Println(jiro)

	misaki := BankAccount{bank_code: "9999",bank_name: "青い反社銀行",branch_code: 998, branch: "東東京支店"}

	fmt.Println(misaki)

	alice  := initAccount("x6931","Bank of Amex","Sapporo",2934001)
	fmt.Println(*alice)
	
	fmt.Println(alice.inputPin("3184"))
}