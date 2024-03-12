package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, m int
	reader := bufio.NewReader(os.Stdin)
	_, err := fmt.Fscanf(reader, "%d %d\n", &n, &m)
	if err != nil || n < 2 || m < 2 || n > 1000 || m > 1000 {
		fmt.Print("Read-mistake1: ", err)
		return
	}

	matrix := make([][]int, n)
	for _, row := range matrix {
		row = make([]int, m)
	}
	var input int
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if j == m-1 {
				_, err = fmt.Fscanf(reader, "%d\n", &input)
			} else {
				_, err = fmt.Fscanf(reader, "%d ", &input)
			}
			if err != nil || input < 1 || input > 1_000_000_000 {
				fmt.Print("Read-mistake2: ", err)
				return
			}
			matrix[i][j] = input
		}
	}

	//заполним массив максимумов в строках
	sliceRow := make([]int, n)
	maxRow
	for i := 0; i < n; i++ {
		sliceRow[i] = findRowMax(&matrix[i])

	}

	fmt.Print(maxRow+1, maxCol+1)
}

func findRowMax(row *[]int) int {
	max := (*row)[0]
	max_idx := 0
	for i := 1; i < len(*row); i++ {
		if (*row)[i] > max {
			max = (*row)[i]
			max_idx = i
		}
	}
	(*row)[max_idx] *= -1
	return max
}

func findColMax(matrix *[][]int, numCol int, isAll bool) int {
	max := (*matrix)[0][numCol]
	max_idx := 0
	for i := 1; i < len(*matrix); i++ {
		if (*matrix)[i][numCol] > max {
			max = (*matrix)[i][numCol]
			max_idx = i
		}
	}
	(*matrix)[max_idx][numCol] *= -1
	return max
}
