package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var N int
	reader := bufio.NewReader(os.Stdin)
	_, err := fmt.Fscanf(reader, "%d\n", &N)
	if err != nil || N < 1 || N > 200000 {
		fmt.Print("Read-mistake1: ", err)
		return
	}
	dict := make(map[string]int)
	for i := 0; i < N; i++ {
		var numMusic int
		_, err = fmt.Fscanf(reader, "%d\n", &numMusic)
		if err != nil {
			fmt.Print("Read-mistake2: ", err)
			return
		}
		for j := 0; j < numMusic; j++ {
			var nameMusic string
			if j == numMusic-1 {
				_, err = fmt.Fscanf(reader, "%s\n", &nameMusic)
			} else {
				_, err = fmt.Fscanf(reader, "%s ", &nameMusic)
			}
			if err != nil {
				fmt.Print("Read-mistake3: ", err)
				return
			}
			dict[nameMusic]++
		}
	}

	answerSlice := make([]string, 0, len(dict))
	for key, value := range dict {
		if value == N {
			answerSlice = append(answerSlice, key)
		}
	}
	sort.Strings(answerSlice)

	fmt.Println(len(answerSlice))

	for i := 0; i < len(answerSlice); i++ {
		if i == 0 {
			fmt.Print(answerSlice[i])
		} else {
			fmt.Print(" ", answerSlice[i])
		}
	}

}
