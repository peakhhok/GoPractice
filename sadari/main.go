package main

import "fmt"

func main() {
	fmt.Println("몇명임?: ")
	var a int
	fmt.Scanln(&a)
	member := make([]string, a)
	for b := 0; b < a; b++ {
		fmt.Println("이름입력_")
		fmt.Scanln(member[b])
	}
	fmt.Print("지금 멤버는: ")
	fmt.Print(member)
	if len(member)%2 == 0 {
		fmt.Println("팀을 나눕니다.")

	}

}
