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
	for i := range matrix {
		matrix[i] = make([]int, m)
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

	//найдем максимум в строках
	rowMax1, _, _ := findMaxRow(&matrix, -1, -1)

	//найдем максимум в столбцах
	_, colMax1, _ := findMaxCol(&matrix, rowMax1)

	//найдем максимум1вар в матрице
	_, _, max1 := findMaxRow(&matrix, rowMax1, colMax1)

	//2найдем максимум в столбцах
	rowMax2, colMax3, _ := findMaxCol(&matrix, -1)

	//2найдем максимум в строках
	_, colMax2, _ := findMaxRow(&matrix, rowMax2, -1)

	//2найдем максимум2вар в матрице
	_, _, max2 := findMaxRow(&matrix, rowMax2, colMax2)

	//3найдем максимум в строках
	rowMax3, _, _ := findMaxRow(&matrix, -1, colMax3)

	//3найдем максимум3вар в матрице
	_, _, max3 := findMaxRow(&matrix, rowMax3, colMax3)

	//сравним максимумы и выведем лучшее сочетание
	if max1 < max2 {
		if max3 < max1 {
			fmt.Print(rowMax3+1, colMax3+1)
		} else {
			fmt.Print(rowMax1+1, colMax1+1)
		}
	} else {
		if max3 < max2 {
			fmt.Print(rowMax3+1, colMax3+1)
		} else {
			fmt.Print(rowMax2+1, colMax2+1)
		}
	}
}

func findMaxRow(matrix *[][]int, skipRow int, skipCol int) (row, col, max int) {
	for i := 0; i < len(*matrix); i++ {
		if i == skipRow {
			continue
		}
		for j := 0; j < len((*matrix)[0]); j++ {
			if (*matrix)[i][j] > max && j != skipCol {
				max = (*matrix)[i][j]
				row = i
				col = j
			}
		}
	}
	return
}

func findMaxCol(matrix *[][]int, skipRow int) (row, col, max int) {
	for i := 0; i < len((*matrix)[0]); i++ {
		for j := 0; j < len(*matrix); j++ {
			if (*matrix)[j][i] > max && j != skipRow {
				max = (*matrix)[j][i]
				col = i
				row = j
			}
		}
	}
	return
}
