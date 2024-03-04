package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	numString, err2 := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil || err2 != nil || numString < 1 || numString > 100000 {
		fmt.Print("Read-mistake-numString")
		return
	}

	var numPress uint64
	var i int64
	//считываем строки с количеством пробелов и сразу считаем нажатия
	for i = 0; i < numString; i++ {
		scanner.Scan()
		err := scanner.Err()
		numSpace, err2 := strconv.ParseUint(scanner.Text(), 10, 64)
		if err != nil || err2 != nil || numSpace < 0 || numSpace > 1000000000 {
			fmt.Printf("Read-mistake #%d", i)
			return
		}
		mul := numSpace / 4
		numPress += mul
		numSpace -= mul * 4
		if numSpace > 1 {
			numPress += 2
		} else if numSpace == 1 {
			numPress++
		}
	}

	fmt.Printf("%d", numPress)
}
