package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var N int
	reader := bufio.NewReader(os.Stdin)
	_, err := fmt.Fscanf(reader, "%d\n", &N)
	if err != nil || N < 1 || N > 100 {
		fmt.Print("Read-mistake1: ", err)
		return
	}

	matrix := make([][]int, N)
	for i := range matrix {
		matrix[i] = make([]int, N)
	}
	var input int
	for i := 0; i < N; i++ {
		for j := 0; j < 2; j++ {
			if j == 1 {
				_, err = fmt.Fscanf(reader, "%d\n", &input)
			} else {
				_, err = fmt.Fscanf(reader, "%d ", &input)
			}
			if err != nil || input < 1 || input > N {
				fmt.Print("Read-mistake2: ", err)
				return
			}
			matrix[i][j] = input
		}
	}
}
