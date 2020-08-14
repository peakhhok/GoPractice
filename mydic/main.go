package main

import (
	"fmt"

	"./indic"
)

func main() {
	dictionary := indic.Dictionary{}
	word := "hi"
	definition := "say hi to you"
	err := dictionary.Add(word, definition)
	if err != nil {
		fmt.Println(err)
	}
	hello, _ := dictionary.Search(word)
	fmt.Println(hello)
	err2 := dictionary.Add(word, definition)
	if err2 != nil {
		fmt.Println(err2)
	}

}
