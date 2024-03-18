package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	dict := make(map[int]int, 3010)
	for i := 0; i < 3; i++ {
		var N int
		_, err := fmt.Fscanf(reader, "%d\n", &N)
		if err != nil || N < 1 || N > 1000 {
			fmt.Print("Read-mistake1: ", err)
			return
		}
		for j := 0; j < N; j++ {
			var num int
			if j == N-1 {
				_, err = fmt.Fscanf(reader, "%d\n", &num)
			} else {
				_, err = fmt.Fscanf(reader, "%d ", &num)
			}
			if err != nil || num < 0 || num > 1_000_000_000 {
				fmt.Print("Read-mistake2: ", err)
				return
			}

			if dict[num] == 0 {
				dict[num] = i + 1
			} else if dict[num] != (i+1) && dict[num] != -1 {
				dict[num] = -1
			}
		}
	}
	answerSlice := make([]int, 0, 3010)
	for key, value := range dict {
		if value == -1 {
			answerSlice = append(answerSlice, key)
		}
	}
	sort.Ints(answerSlice)
	for i := 0; i < len(answerSlice); i++ {
		if i == 0 {
			fmt.Print(answerSlice[i])
		} else {
			fmt.Print(" ", answerSlice[i])
		}
	}
}
