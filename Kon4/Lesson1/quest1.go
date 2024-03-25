package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d\n", &n)
	var arr = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d", &arr[i])
	}
	qSort(arr, 0, len(arr)-1)
	var k int
	var min, max int
	fmt.Fscanf(reader, "\n%d\n", &k)
	for i := 0; i < k; i++ {
		if i != 0 {
			fmt.Print(" ")
		}
		fmt.Fscanf(reader, "%d %d\n", &min, &max)

		var res int
		left := lbinSearch(arr, min)
		if left == -1 {
			res = 0
		} else {
			right := rbinSearch(arr, max)
			if right == -1 {
				res = 0
			} else {
				res = right - left + 1
			}
		}
		fmt.Print(res)
	}
}

func qSort(arr []int, first, last int) {
	if first < last {
		left, right := first, last
		mid := arr[(left+right)/2]

		for left <= right {
			for arr[left] < mid {
				left++
			}
			for arr[right] > mid {
				right--
			}
			if left <= right {
				arr[left], arr[right] = arr[right], arr[left]
				left++
				right--
			}
		}
		qSort(arr, first, right)
		qSort(arr, left, last)
	}
}

func rbinSearch(arr []int, srcVal int) int {
	left := 0
	right := len(arr) - 1
	for left < right {
		mid := (left + right + 1) / 2
		guess := arr[mid]
		if check(srcVal, guess) {
			left = mid
		} else {
			right = mid - 1
		}
	}
	if arr[left] > srcVal {
		left = -1
	}

	return left
}

func lbinSearch(arr []int, srcVal int) int {
	left := 0
	right := len(arr) - 1
	for left < right {
		mid := (left + right) / 2
		guess := arr[mid]
		if check(guess, srcVal) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	if arr[left] < srcVal {
		left = -1
	}
	return left
}

func check(guess, srcVal int) bool {
	return guess >= srcVal
}
