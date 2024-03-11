package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// cчитаем общее количество
	var num int
	reader := bufio.NewReader(os.Stdin)
	_, err := fmt.Fscanf(reader, "%d\n", &num)
	if err != nil || num < 1 || num > 5_00_000 {
		fmt.Print("Read-mistake1: ", err)
		return
	}
	//считаем последовательно характеристики ягод в слайсы
	matrix := make([][]int, num)
	for i := range matrix {
		matrix[i] = make([]int, 3)
	}

	for i := 0; i < num; i++ {
		var up, down int
		_, err = fmt.Fscanf(reader, "%d %d\n", &up, &down)
		if err != nil || up < 0 || up > 1_000_000_000 || down < 0 || down > 1_000_000_000 {
			fmt.Print("Read-mistake2: ", err)
			return
		}
		/* сортируем объекты по правилу
		если отрицательная дельта то сортируем по вверх только с отрицательными

		иначе если очередная дельта больше то свап
						если равен то смотрим на 0 столбец
								если меньше то свап
		*/
		flagAdded := false

		for j := 0; j < i && !flagAdded; j++ {
			if up-down < 0 && up > matrix[j][0] && matrix[j][1] < 0 {
				displacement(i, j, &matrix, true)
				addInstead(j, up, down, i, &matrix)
				flagAdded = true
			} else if up-down > matrix[j][1] && up-down >= 0 {
				displacement(i, j, &matrix, true)
				addInstead(j, up, down, i, &matrix)
				flagAdded = true
			} else if up-down == matrix[j][1] && up-down >= 0 {
				if up < matrix[j][0] {
					displacement(i, j, &matrix, true)
					addInstead(j, up, down, i, &matrix)
					flagAdded = true
				}
			}
		}
		if !flagAdded {
			addInstead(i, up, down, i, &matrix)
		}
	}
	//максимальный up-down с положительным down нужно поставить после последнего положительного смещения
	max := -1
	var max_id int
	startMinus := -1
	for i := 0; i < num; i++ {
		if matrix[i][1] >= 0 && matrix[i][0]-matrix[i][1] >= max {
			max = matrix[i][0] - matrix[i][1]
			max_id = i
		}
		if matrix[i][1] <= 0 && startMinus == -1 {
			startMinus = i - 1
		}
	}
	if startMinus == -1 {
		startMinus = num - 1
	}
	if max_id > startMinus {
		tempUp := matrix[max_id][0]
		tempDown := tempUp - matrix[max_id][1]
		tempNo := matrix[max_id][2]
		displacement(max_id, startMinus, &matrix, true)
		addInstead(startMinus+1, tempUp, tempDown, tempNo, &matrix)
	} else if max_id < startMinus {
		tempUp := matrix[max_id][0]
		tempDown := tempUp - matrix[max_id][1]
		tempNo := matrix[max_id][2]
		displacement(max_id, startMinus, &matrix, false)
		addInstead(startMinus, tempUp, tempDown, tempNo, &matrix)
	}

	show(&matrix)
	//showTest(&matrix)
}

func addInstead(pos int, up, down, i int, matrix *[][]int) {
	(*matrix)[pos][0] = up
	(*matrix)[pos][1] = up - down
	(*matrix)[pos][2] = i
}

func displacement(i, j int, matrix *[][]int, isDown bool) {
	if isDown {
		for k := i; k > j; k-- {
			(*matrix)[k][0] = (*matrix)[k-1][0]
			(*matrix)[k][1] = (*matrix)[k-1][1]
			(*matrix)[k][2] = (*matrix)[k-1][2]
		}
	} else {
		for k := i; k < j; k++ {
			(*matrix)[k][0] = (*matrix)[k+1][0]
			(*matrix)[k][1] = (*matrix)[k+1][1]
			(*matrix)[k][2] = (*matrix)[k+1][2]
		}
	}
}

func show(matrix *[][]int) {
	var curr, max int
	for _, row := range *matrix {
		if curr+row[0] > max {
			max = curr + row[0]
		}
		curr += row[1]
	}
	fmt.Println(max)
	for idx, row := range *matrix {
		if idx == 0 {
			fmt.Print(row[2] + 1)
		} else {
			fmt.Print(" ", row[2]+1)
		}
	}
}

// func showTest(matrix *[][]int) {
// 	for idxRow, row := range *matrix {
// 		fmt.Print("\n", idxRow, ": ")
// 		for _, val := range row {
// 			fmt.Print(val, " ")
// 		}
// 	}
// }
