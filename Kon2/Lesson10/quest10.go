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
	array := make([][]int, m)
	for i := range array {
		array[i] = make([]int, n)
	}

	var counter int
	var symbol byte
	for i := 0; i < m; i++ {
		for j := 0; j <= n; j++ {
			_, err := fmt.Fscanf(reader, "%c", &symbol)
			if err != nil || (symbol != '#' && symbol != '.' && j != n) || (j == n && symbol != '\n') {
				fmt.Print("Read-mistake2: ", err)
				return
			}
			if symbol == '#' {
				array[i][j] = 1
			}
		}
	}
	parsedArray := make([][]int, m)
	for i := range parsedArray {
		parsedArray[i] = make([]int, n)
		copy(parsedArray[i], array[i])
	}

	parsedArray, counter = parseArr(parsedArray, m, n)
	if counter < 2 {
		fmt.Print("NO")
		return
	}

	if !ProcessHoriz(&parsedArray, m, n, counter, false) {
		array = transposeArr(&array, m, n)
		array, _ = parseArr(array, n, m)
		if !ProcessHoriz(&array, n, m, counter, true) {
			fmt.Print("NO")
		}
	}
}

func ProcessHoriz(array *[][]int, m, n, counter int, reverse bool) bool {
	//обработка горизонтальная
	num := 0 //изначально 0, после при 1-а, 2-б
	var isEnd bool
	for i := 0; i < m && !isEnd; i++ {
		for j := 0; j < n && !isEnd; j++ {
			if (*array)[i][j] != 0 {
				if counter-(*array)[i][j] == 0 && num == 0 {
					counter -= (*array)[i][j]
					for k := 1; k < (*array)[i][j]; k++ {
						(*array)[i][j-k] = 1
					}
					(*array)[i][j] = 2
					num = 3
				} else {
					var add bool
					//самая первая проходка
					if num == 0 {
						num = 1
					}
					counter -= (*array)[i][j]
					if i == m-1 || (counter-(*array)[i+1][j] == 0 || (*array)[i+1][j] != (*array)[i][j]) {
						add = true
					}
					for k := (*array)[i][j] - 1; k >= 0; k-- {
						(*array)[i][j-k] = num
					}
					if add {
						num++
					}
				}
			}
			if num == 3 {
				isEnd = true
				if counter != 0 {
					return false
				}
			}
		}
	}
	if reverse {
		reverseArr := transposeArr(array, m, n)
		printArr(&reverseArr, n, m)
	} else {
		printArr(array, m, n)
	}
	return true
}

func printArr(array *[][]int, m, n int) {
	//вывод ответа
	fmt.Println("YES")
	var answer string
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			switch (*array)[i][j] {
			case 1:
				answer = "a"
			case 2:
				answer = "b"
			default:
				answer = "."
			}
			fmt.Print(answer)
		}
		if i != m-1 {
			fmt.Println()
		}
	}
}

func parseArr(array [][]int, m, n int) (arrayNew [][]int, counter int) {
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if array[i][j] == 1 {
				counter++
				if j != 0 {
					array[i][j] = array[i][j-1] + 1
					array[i][j-1] = 0
				} else {
					array[i][j] = 1
				}
			}
		}
	}
	return array, counter
}

func transposeArr(array *[][]int, m, n int) (arrayNew [][]int) {
	arrayNew = make([][]int, n)
	for i := range arrayNew {
		arrayNew[i] = make([]int, m)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			arrayNew[j][i] = (*array)[i][j]
		}
	}

	return
}
