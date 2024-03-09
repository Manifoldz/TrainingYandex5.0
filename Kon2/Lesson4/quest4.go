package main

import "fmt"

func main() {
	var N int
	fmt.Scanf("%d\n", &N)
	var x, y int
	matrix := make([][]int, 10)
	for i := 0; i < 10; i++ {
		matrix[i] = make([]int, 10)
	}

	for i := 0; i < N; i++ {
		_, err := fmt.Scanf("%d %d\n", &x, &y)
		if err != nil || x < 1 || y < 1 || x > 8 || y > 8 {
			fmt.Print("Read-mistake: ", err)
			return
		}
		matrix[x][y] = 1
	}

	var neighbourY [4]int = [4]int{1, -1, 0, 0}
	var neighbourX [4]int = [4]int{0, 0, -1, 1}
	var sum int
	for i := 1; i < 9; i++ {
		for j := 1; j < 9; j++ {
			if matrix[i][j] == 1 {
				for k := 0; k < 4; k++ {
					sum += 1 - matrix[i+neighbourY[k]][j+neighbourX[k]]
				}
			}
		}
	}

	fmt.Print(sum)
}
