package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var N, K int
	reader := bufio.NewReader(os.Stdin)
	_, err := fmt.Fscanf(reader, "%d %d\n", &N, &K)
	if err != nil || N < 1 || N > 100000 || K < 1 || K > 100000 {
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
		if err != nil || num < -1_000_000_000 || num > 1_000_000_000 {
			fmt.Print("Read-mistake2: ", err)
			return
		}

		if dict[num] != 0 && (i+1)-dict[num] <= K {
			fmt.Print("YES")
			return
		} else {
			dict[num] = i + 1
		}
	}

	fmt.Print("NO")
}
