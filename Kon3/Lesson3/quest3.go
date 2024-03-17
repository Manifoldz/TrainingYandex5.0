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
	if err != nil || N < 1 || N > 200000 {
		fmt.Print("Read-mistake1: ", err)
		return
	}
	dict := make(map[int]int)
	for i := 0; i < N; i++ {
		var num int
		if i == N-1 {
			_, err = fmt.Fscanf(reader, "%d\n", &num)
		} else {
			_, err = fmt.Fscanf(reader, "%d ", &num)
		}
		if err != nil || num < 0 || num > 100000 {
			fmt.Print("Read-mistake2: ", err)
			return
		}
		dict[num]++
	}

	var max int
	for key, value := range dict {
		if value+dict[key+1] > max {
			max = value + dict[key+1]
		}
	}

	fmt.Print(N - max)

}
