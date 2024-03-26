package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, m uint64
	fmt.Fscanf(reader, "%d %d\n", &n, &m)
	arr := make([]uint64, n+1)
	var temp uint64
	arr[0] = 0
	for i := uint64(1); i <= n; i++ {
		if i == n {
			fmt.Fscanf(reader, "%d\n", &temp)
		} else {
			fmt.Fscanf(reader, "%d ", &temp)
		}
		arr[i] = temp + arr[i-1]
	}

	var l, s uint64
	for i := uint64(0); i < m; i++ {
		fmt.Fscanf(reader, "%d %d\n", &l, &s)
		fmt.Println(lbinSearch(arr, l, s))
	}
}

func lbinSearch(arr []uint64, l, s uint64) int {
	left := l
	right := uint64(len(arr) - 1)
	for left < right {
		mid := (left + right) / 2
		if arr[mid]-arr[mid-l] >= s {
			right = mid
		} else {
			left = mid + 1
		}
	}
	if arr[left]-arr[left-l] != s {
		return -1
	}
	return int(left - l + 1)
}
