package main

import (
	"fmt"
	"math/rand"
	"time"
)

func BubbleSort(s []int) []int {
	for i := 0; i < len(s)-1; i++ {
		for j := 0; j < len(s)-i-1; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}

	return s
}

func main() {
	rand.Seed(time.Now().Unix())

	const N = 10
	s := make([]int, N)
	for i := 0; i < N; i++ {
		s[i] = rand.Intn(N)
	}
	fmt.Println(s)
	fmt.Println(BubbleSort(s))
}
