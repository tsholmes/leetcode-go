package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var counts [11]int
	for i := 0; i < 10000; i++ {
		counts[rand10()]++
	}
	fmt.Println(counts[1:])
}

func rand7() int {
	return rand.Intn(7) + 1
}

func rand10() int {
	res := 0
	for i := 0; i < 10; i++ {
		res += rand7()
	}
	return (res % 10) + 1
}
