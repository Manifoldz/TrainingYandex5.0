package main

import (
	"fmt"
)

func main() {
	var n uint64
	fmt.Scanf("%d\n", &n)
	left := uint64(0)
	right := n
	for left < right {
		mid := (left + right + 1) / 2
		if check(mid, n) {
			left = mid
		} else {
			right = mid - 1
		}
	}

	fmt.Print(left)
}

func check(guess, n uint64) bool {
	if guess == 1 {
		return 1 <= n
	} else if guess > 2_000_000 {
		return false
	}
	var sum = guess * (guess + 1) * (guess + 2) / 6
	if sum > n {
		return false
	}
	return sum-1+(guess+1)*guess/2 <= n
}
