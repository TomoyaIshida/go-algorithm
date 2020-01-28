package main

import (
	"fmt"
)

func BinarySearch(t int, s []int) int {
	left := 0
	right := len(s) - 1
	for left <= right {
		mid := (left + right) / 2
		if s[mid] == t {
			return mid + 1
		} else if s[mid] < t {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target := 5
	fmt.Println(slice)
	fmt.Println(target)

	result := BinarySearch(target, slice)
	var str string
	if result == -1 {
		str = fmt.Sprintf("%dは見つかりません", target)
	} else {
		str = fmt.Sprintf("%dは%d番目です", target, result)
	}
	fmt.Println(str)
}
