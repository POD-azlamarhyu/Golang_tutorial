package main

import (
	"fmt"
)

func main() {
	var account_level int
	var is_account_level string

	account_level = 6

	switch account_level {
	case 1:
		is_account_level = "不正アカウントです"
	case 2:
		is_account_level = "違反アカウントです"
	case 3:
		is_account_level = "制限アカウントです"
	case 4:
		is_account_level = "ノーマルアカウントです"
	case 5:
		is_account_level = "上位アカウントです"
	case 6:
		is_account_level = "スーパーアカウントです"
	case 7:
		is_account_level = "マスターアカウントです"
	case 8:
		is_account_level = "アドミンアカウントです"
	default:
		is_account_level = "アカウントが存在しません"
	}

	fmt.Println(is_account_level)
}