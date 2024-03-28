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

	var ans uint64
	for left < right {
		m1 := left + (right-left)/3
		m2 := right - (right-left)/3
		//считаем для м1
		minRowM11 := countRows(arr1, (width - m1))
		minRowM12 := countRows(arr2, m1)

		//посчитаем для м2 и если оно больше чем макс первых двух, то дальше не нужно второе проверять
		minRowM21 := countRows(arr1, (width - m2))
		minRowM22 := countRows(arr2, m2)
		//досрочный выход
		if minRowM11 == minRowM12 {
			ans = minRowM11
			break
		} else if minRowM21 == minRowM22 {
			ans = minRowM21
			break
		}

		//определим максимумы в обеих точках
		ans = max(minRowM11, minRowM12)
		minRowM21 = max(minRowM21, minRowM22)
		if minRowM21 <= ans {
			left = m1 + 1
			ans = minRowM21
		} else {
			right = m2 - 1
		}
	}

	if ans == 0 {
		ans = max(countRows(arr2, left), countRows(arr1, (width-left)))
	}

	fmt.Println(ans)
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
