package main

import (
	"fmt"
)

func LinearSearch(t int, s []int) int {
	for n := 0; n < len(s); n++ {
		if s[n] == t {
			return n + 1
		}
	}
	return -1
}

func main() {
	slice := []int{1, 3, 5, 4, 2}
	target := 2
	fmt.Println(slice)
	fmt.Println(target)

	result := LinearSearch(target, slice)
	var str string
	if result == -1 {
		str = fmt.Sprintf("%dは見つかりません", target)
	} else {
		str = fmt.Sprintf("%dは%d番目です", target, result)
	}
	fmt.Println(str)
}
