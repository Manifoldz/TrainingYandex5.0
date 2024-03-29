package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var width, numWord1, numWord2 uint64
	fmt.Fscanf(reader, "%d %d %d\n", &width, &numWord1, &numWord2)
	arr1 := make([]uint64, numWord1)
	var minWidth1, minWidth2 uint64
	for i := uint64(0); i < numWord1; i++ {
		var temp uint64
		if i == numWord1-1 {
			fmt.Fscanf(reader, "%d\n", &temp)
		} else {
			fmt.Fscanf(reader, "%d ", &temp)
		}
		if temp > minWidth1 {
			minWidth1 = temp
		}
		if i == 0 {
			arr1[i] = temp
		} else {
			arr1[i] = arr1[i-1] + temp
		}
	}
	arr2 := make([]uint64, numWord2)
	for i := uint64(0); i < numWord2; i++ {
		var temp uint64
		if i == numWord2-1 {
			fmt.Fscanf(reader, "%d\n", &temp)
		} else {
			fmt.Fscanf(reader, "%d ", &temp)
		}
		if temp > minWidth2 {
			minWidth2 = temp
		}
		if i == 0 {
			arr2[i] = temp
		} else {
			arr2[i] = arr2[i-1] + temp
		}
	}
	left := minWidth2
	right := width - minWidth1

	if left > right {
		left, right = right, left
	}

	minRowM11 := countRows(arr1, (width - left))
	minRowM12 := countRows(arr2, left)

	//досрочный выход
	if minRowM11 == minRowM12 {
		fmt.Println(minRowM11)
		return
	}

	saveFirst := max(minRowM11, minRowM12)

	var isFirstMore bool
	if minRowM11 > minRowM12 {
		isFirstMore = true
	}
	//fmt.Println(left, right, saveFirst)
	var minCommon uint64 = 18446744073709551615
	for left < right {
		mid := (left + right) / 2
		minRowM11 = countRows(arr1, (width - mid))
		minRowM12 = countRows(arr2, mid)
		minLocal := max(minRowM11, minRowM12)
		if minLocal < minCommon {
			minCommon = minLocal
		}
		if isFirstMore != (minRowM11 >= minRowM12) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	minLocal := max(countRows(arr1, (width-left)), countRows(arr2, left))
	if minLocal < minCommon {
		minCommon = minLocal
	}

	fmt.Println(min(minCommon, saveFirst))
}

func countRows(arr []uint64, width uint64) (numRows uint64) {
	step := width
	left := 0
	start := 0
	for left != len(arr)-1 {
		left = start
		right := len(arr) - 1
		for left < right {
			mid := (left + right + 1) / 2
			if arr[mid]+uint64(mid-start) <= width {
				left = mid
			} else {
				right = mid - 1
			}
		}
		width = width + step - (width - arr[left])
		numRows++
		start = left + 1
	}

	return
}

func max(a, b uint64) uint64 {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b uint64) uint64 {
	if a > b {
		return b
	} else {
		return a
	}
}
