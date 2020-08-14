package main

import (
	"fmt"

	"./indic"
)

func main() {
	dictionary := indic.Dictionary{}
	baseword := "hi"
	dictionary.Add(baseword, "First")
	dictionary.Search(baseword)
	dictionary.Delete(baseword)
	word, err := dictionary.Search(baseword)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(word)
}
