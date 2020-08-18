package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Print("몇명입니까? (숫자만 입력): ")
	var a, d int
	var c, input string

	fmt.Scanln(&a)
	member := make([]string, a)
	if len(member)%2 == 0 {
		for b := 0; b < a; b++ {
			fmt.Print("이름입력", b+1, ":")

			fmt.Scanln(&c)
			member[b] = c
		}

		fmt.Println("지금 멤버는:", member)

		time.Sleep(time.Second * 2)

		fmt.Println("팀을 나눕니다.")
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(member), func(i, j int) { member[i], member[j] = member[j], member[i] })

		fmt.Print("1팀은:")
		for d = len(member) / 2; d > 0; d-- {

			fmt.Print(member[d-1], " ")
		}
		fmt.Println("")
		fmt.Print("2팀은:")
		for e := len(member) / 2; e < len(member); e++ {
			fmt.Print(member[e], " ")
		}
		fmt.Println("")
		fmt.Println(" 종료하려면 엔터키 ")
		fmt.Scanln(&input)

	} else {
		fmt.Println("멤버를 나눌 수 없습니다")
		fmt.Println(" 종료하려면 엔터키 ")
		fmt.Scanln(&input)
	}
}
