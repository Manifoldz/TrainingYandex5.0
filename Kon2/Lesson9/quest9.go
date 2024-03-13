package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

func main() {
	var N int
	reader := bufio.NewReader(os.Stdin)
	_, err := fmt.Fscanf(reader, "%d\n", &N)
	if err != nil || N < 1 || N > 100 {
		fmt.Print("Read-mistake1: ", err)
		return
	}

	sliceY := make([]int, 0, N)
	sliceX := make([]int, 0, N)
	var input int
	for i := 0; i < N*2; i++ {
		if i%2 == 1 {
			_, err = fmt.Fscanf(reader, "%d\n", &input)
		} else {
			_, err = fmt.Fscanf(reader, "%d ", &input)
		}
		if err != nil || input < 1 || input > N {
			fmt.Print("Read-mistake2: ", err)
			return
		}
		if i%2 == 1 {
			sliceX = append(sliceX, input)
		} else {
			sliceY = append(sliceY, input)
		}
	}
	sort.Ints(sliceX)
	sort.Ints(sliceY)
	median := sliceX[N/2]
	var sum int
	for i := 0; i < N; i++ {
		sum += int(math.Abs(float64(sliceX[i] - median)))
		sum += int(math.Abs(float64(sliceY[i] - i - 1)))
	}
	fmt.Print(sum)
}
