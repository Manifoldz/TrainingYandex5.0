package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var m, n int
	reader := bufio.NewReader(os.Stdin)
	_, err := fmt.Fscanf(reader, "%d %d\n", &m, &n)
	if err != nil || n < 1 || m < 1 || n > 200 || m > 200 {
		fmt.Print("Read-mistake1: ", err)
		return
	}
	array := make([][]byte, m+2)
	for i := range array {
		array[i] = make([]byte, n+2)
	}

	var symbol byte
	var counter int
	for i := 1; i < m+1; i++ {
		for j := 1; j <= n+1; j++ {
			_, err := fmt.Fscanf(reader, "%c", &symbol)
			if err != nil || (symbol != '#' && symbol != '.' && j != n) || (j == n && symbol != '\n') {
				fmt.Print("Read-mistake2: ", err)
				return
			}
			if symbol == '#' {
				counter++
			}
			if j != n+1 {
				array[i][j] = symbol
			}
		}
	}
	sliceOffsetY := []int{-1, -1, -1, 0, 1, 0, 1, 1}
	sliceOffsetX := []int{-1, 0, 1, -1, -1, 1, 1, 0}

	var isExistA bool
	var isExistB bool
	var isEnd bool
	//обработка матрицы
	for i := 0; i < m && !isEnd; i++ {
		for j := 0; j < n && !isEnd; j++ {
			if array[i][j] == '#' {
				//если последняя ячейка
				if counter < 2 {
					if isExistA && !isExistB {
						array[i][j] = 'b'
						isEnd = true
					} else if !isExistA {
						fmt.Print("NO")
						return
					}
				}
				// если первая ячейка то пишем в нее 'a' и выходим
				if !isExistA {
					array[i][j] = 'a'
					counter--
					continue
				}
				//иначе отдаем функции подсчет
				array[i][j] = neighbours(&array, &sliceOffsetY, &sliceOffsetX, i, j)
			}
		}
	}

	//вывод ответа
	fmt.Println("YES")
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Print(string(array[i][j]))
		}
		if i != m-1 {
			fmt.Println()
		}
	}
}

func neighbours(array *[][]byte, sliceY, sliceX *[]int, i, j int) (answer byte) {
	var regA, regB int
	var crossA, crossB int
	for k := 0; k < 8; k++ {
		if (*array)[i+(*sliceY)[k]][j+(*sliceX)[k]] == 'a' {
			if k%2 == 1 {
				crossA++
			} else {
				regA++
			}
		} else if (*array)[i+(*sliceY)[k]][j+(*sliceX)[k]] == 'b' {
			if k%2 == 1 {
				crossB++
			} else {
				regB++
			}
		}
	}

	if regA+crossA > 1 {
		return 'b'
	}

	if

	if 

	return
}
