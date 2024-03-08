package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	var N, K int64
	_, err := fmt.Scanf("%d %d\n", &N, &K)
	if err != nil || N < 1 || N > 100000 || K < 1 || K > 100 {
		fmt.Print("Read-mistake:", err)
		return
	}

	var diff int64
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	num_arr := make([]int64, 0, N+2)

	num_arr = append(num_arr, -1)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		num, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil || num > 1000000000 || num < 1 {
			fmt.Print("Convert-mistake: ", num)
			return
		}
		num_arr = append(num_arr, num)
	}
	num_arr = append(num_arr, -1)

	for i := int64(1); num_arr[i] != -1; i++ {
		for j := int64(1); num_arr[i+j] != -1 && j <= K; j++ {
			curr_diff := num_arr[i+j] - num_arr[i]
			if curr_diff > diff {
				diff = curr_diff
			}
		}
	}

	fmt.Print(diff)

}
