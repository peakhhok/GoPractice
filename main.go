package main

import (
	"fmt"

	"./accounts"
)

func main() {

	account := accounts.NewAccount("HO")
	account.Deposit(10)
	fmt.Println(account.Balance())
	err := account.Withdraw(20)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(account)
}
