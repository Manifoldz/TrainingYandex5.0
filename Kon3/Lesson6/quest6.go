package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//настройка считывания
	scanner := bufio.NewScanner(os.Stdin)
	bufer := make([]byte, 1000001)
	scanner.Buffer(bufer, 1000001)

	//первый ввод записываем в слайс и сортируем там же
	if !scanner.Scan() {
		fmt.Print("Read-mistake1: ", scanner.Err())
		return
	}
	sliceStr := strings.Split(scanner.Text(), " ")
	qSortLen(sliceStr, 0, len(sliceStr)-1)

	//второй слайс считываем
	if !scanner.Scan() {
		fmt.Print("Read-mistake2: ", scanner.Err())
		return
	}
	sliceStr2 := strings.Split(scanner.Text(), " ")

	for idx, str := range sliceStr2 {
		var isPrinted bool
		for _, strDict := range sliceStr {
			if strings.HasPrefix(str, strDict) {
				if idx != 0 {
					fmt.Printf(" %s", strDict)
				} else {
					fmt.Printf("%s", strDict)
				}
				isPrinted = true
				break
			}
		}
		if !isPrinted {
			if idx != 0 {
				fmt.Printf(" %s", str)
			} else {
				fmt.Printf("%s", str)
			}
		}
	}
}

func qSortLen(sliceStr []string, first, last int) {
	if first < last {
		left, right := first, last
		middle := len(sliceStr[(left+right)/2])
		for left <= right {
			for len(sliceStr[left]) < middle {
				left++
			}
			for len(sliceStr[right]) > middle {
				right--
			}
			if left <= right {
				sliceStr[left], sliceStr[right] = sliceStr[right], sliceStr[left]
				left++
				right--
			}
		}
		qSortLen(sliceStr, first, right)
		qSortLen(sliceStr, left, last)
	}
}
