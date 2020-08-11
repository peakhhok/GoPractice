package main

import (
	"fmt"

	"./accounts"
)

func main() {

	account := accounts.NewAccount("HO")

	fmt.Println(account)
}
