package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var t int
	reader := bufio.NewReader(os.Stdin)

	_, err := fmt.Fscanf(reader, "%d\n", &t)
	if err != nil || t < 1 || t > 1000 {
		fmt.Print("Read-mistake1: ", err)
		return
	}

	matrix := make([][]int, t)
	for k := 0; k < t; k++ {
		var n int
		_, err = fmt.Fscanf(reader, "%d\n", &n)
		if err != nil || n < 1 || n > 100000 {
			fmt.Print("Read-mistake2: ", err)
			return
		}
		var num, startPos, min int
		matrix[k] = make([]int, 0, n)
		for i := 0; i < n; i++ {
			if i != n-1 {
				_, err = fmt.Fscanf(reader, "%d ", &num)
			} else {
				_, err = fmt.Fscanf(reader, "%d\n", &num)
			}
			if err != nil || num < 1 || num > n {
				fmt.Print("Read-mistake3: ", err)
				return
			}

			if i == 0 {
				min = num
			}

			//последний элемент
			if i == n-1 {
				if i-startPos >= num || i-startPos >= min {
					matrix[k] = append(matrix[k], i-startPos)
					matrix[k] = append(matrix[k], 1)
				} else {
					matrix[k] = append(matrix[k], i-startPos+1)
				}
				continue
			}

			//остальные
			if i-startPos >= min || i-startPos >= num {
				matrix[k] = append(matrix[k], i-startPos)
				startPos = i
				min = num
			} else if num < min {
				min = num
			}

		}
	}
	//вывод
	for _, str := range matrix {
		fmt.Println(len(str))
		for idx, num := range str {
			if idx == 0 {
				fmt.Print(num)
			} else {
				fmt.Print(" ", num)
			}

		}
		fmt.Print("\n")
	}
}
