package main

import (
	"fmt"
	"math/rand"
	"time"
)

func QuickSort(a []int, left int, right int) {
	if right > left {
		pivot := Partition(a, left, right)
		QuickSort(a, left, pivot-1)
		QuickSort(a, pivot+1, right)
	}
}

func Partition(a []int, left int, right int) int {
	lower := left
	pivot := a[left]
	for left < right {
		for a[left] <= pivot {
			left++
		}
		for a[right] > pivot {
			right--
		}
		if left < right {
			a[right], a[left] = a[left], a[right]
		}
	}
	a[lower] = a[right]
	a[right] = pivot
	return right
}

func main() {
	rand.Seed(time.Now().Unix())

	const N = 10000
	s := make([]int, N)
	for i := 0; i < N; i++ {
		s[i] = rand.Intn(N)
	}
	fmt.Println(s)

	start := time.Now()

	QuickSort(s, 0, len(s)-1)

	fmt.Println(s)

	end := time.Now()

	fmt.Printf("%fs\n", (end.Sub(start)).Seconds())
}
