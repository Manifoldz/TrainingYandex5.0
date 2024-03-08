package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var N int
	_, err := fmt.Scanf("%d\n", &N)
	if err != nil || N < 2 || N > 1000 {
		fmt.Print("Read-mistake: ", err)
		return
	}

	reader := bufio.NewReader(os.Stdin)

	var sum uint64
	var max int

	for i := 0; i < N; i++ {
		var num int
		if i == N-1 {
			_, err = fmt.Fscanf(reader, "%d\n", &num)
		} else {
			_, err = fmt.Fscanf(reader, "%d ", &num)
		}

		if err != nil || num < 1 || num > 1000 {
			fmt.Print("Read-mistake2: ", err)
			return
		}

		if num > max {
			sum += uint64(max)
			max = num
		} else {
			sum += uint64(num)
		}
	}

	var l uint64
	if uint64(max) > sum {
		l = uint64(max) - sum
	} else {
		l = uint64(max) + sum
	}
	fmt.Print(l)
}
