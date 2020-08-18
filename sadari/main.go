package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Print("몇명임?: ")
	var a int
	fmt.Scanln(&a)
	member := make([]string, a)
	for b := 0; b < a; b++ {
		fmt.Println("이름입력", b, ":")

		fmt.Scan(member[b])
	}
	fmt.Print("지금 멤버는: ")
	fmt.Println(member)
	time.Sleep(time.Second * 2)
	if len(member)%2 == 0 {
		fmt.Println("팀을 나눕니다.")
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(member), func(i, j int) { member[i], member[j] = member[j], member[i] })
		fmt.Println(member)

	}

}
