package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// cчитаем общее количество
	//startTime := time.Now()
	var num int
	reader := bufio.NewReader(os.Stdin)
	_, err := fmt.Fscanf(reader, "%d\n", &num)
	if err != nil || num < 1 || num > 5_00_000 {
		fmt.Print("Read-mistake1: ", err)
		return
	}
	//считаем последовательно характеристики ягод в слайсы
	matrixPos := make([]int, 0, num)
	matrixNepos := make([]int, 0, num)

	maxPosUp, maxPosDown, maxPosNum := -1, 0, -1
	maxNepos, maxNeposNum := -1, -1
	var currHeight, maxHeight int // высоты
	for i := 0; i < num; i++ {
		var up, down int
		_, err = fmt.Fscanf(reader, "%d %d\n", &up, &down)
		if err != nil || up < 0 || up > 1_000_000_000 || down < 0 || down > 1_000_000_000 {
			fmt.Print("Read-mistake2: ", err)
			return
		}
		/* При считывании сразу будем искать максимальое положительное и отрицательное число,
		при смене будем выводить сразу старое число.
		Сразу выводим положительные числа, нулевые и отрицательные сохраняем в массив, далее выводим макс полож.,
		макс неположительное, и проходимся по массиву печатаю сразу
		*/
		diff := up - down
		if diff > 0 { // положительные
			if maxPosNum == -1 { //первая запись
				maxPosDown = down
				maxPosUp = up
				maxPosNum = i
			} else if down >= maxPosDown { //выше по дельте
				matrixPos = append(matrixPos, maxPosNum)
				if currHeight+maxPosUp > maxHeight {
					maxHeight = currHeight + maxPosUp
				}
				currHeight += maxPosUp - maxPosDown
				maxPosDown = down
				maxPosUp = up
				maxPosNum = i
			} else {
				matrixPos = append(matrixPos, i)
				if currHeight+up > maxHeight {
					maxHeight = currHeight + up
				}
				currHeight += diff
			}
		} else if up > maxNepos { //неположительные выше макса
			if maxNeposNum != -1 { //не впервые
				matrixNepos = append(matrixNepos, maxNeposNum)
			}
			maxNepos = up
			maxNeposNum = i
		} else { //остальные неположительные
			matrixNepos = append(matrixNepos, i)
		}
	}

	if maxPosNum != -1 {
		matrixPos = append(matrixPos, maxPosNum)
		if currHeight+maxPosUp > maxHeight {
			maxHeight = currHeight + maxPosUp
		}
		currHeight += maxPosUp - maxPosDown
	}
	if maxNeposNum != -1 {
		matrixPos = append(matrixPos, maxNeposNum)
		if currHeight+maxNepos > maxHeight {
			maxHeight = currHeight + maxNepos
		}
	}
	matrixPos = append(matrixPos, matrixNepos...)

	fmt.Println(maxHeight)
	for idx, val := range matrixPos {
		if idx != 0 {
			fmt.Print(" ", val+1)
		} else {
			fmt.Print(val + 1)
		}
	}

	//diffTime := time.Since(startTime)
	//fmt.Println(diffTime)
}
