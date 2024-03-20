package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)
	//имеющиеся
	haveMap := make(map[string][]string)
	for i := 0; i < n; i++ {
		var x1, y1, x2, y2 int
		fmt.Scanf("%d %d %d %d\n", &x1, &y1, &x2, &y2)
		p1 := fmt.Sprintf("x%dy%d", x1, y1)
		p2 := fmt.Sprintf("x%dy%d", x2, y2)

		haveMap[p1] = append(haveMap[p1], fmt.Sprintf("dx%ddy%d", x2-x1, y2-y1))
		haveMap[p2] = append(haveMap[p2], fmt.Sprintf("dx%ddy%d", x1-x2, y1-y2))
	}

	//необходимые
	needMap := make(map[string][]string)
	for i := 0; i < n; i++ {
		var x1, y1, x2, y2 int
		fmt.Scanf("%d %d %d %d\n", &x1, &y1, &x2, &y2)
		p1 := fmt.Sprintf("x%dy%d", x1, y1)
		p2 := fmt.Sprintf("x%dy%d", x2, y2)
		needMap[p1] = append(needMap[p1], fmt.Sprintf("dx%ddy%d", x2-x1, y2-y1))
		needMap[p2] = append(needMap[p2], fmt.Sprintf("dx%ddy%d", x1-x2, y1-y2))
	}

	//поиск базы(найдем самую подходящую площадку)
	max := 0
	var staticSlice []string
	var static
	for keyNeed, valNeedSlice := range needMap {
		if len(valHaveSlice) <= max {
			continue
		}
		for keyHave, valHaveSlice := range haveMap {
			if len(valHaveSlice) <= max {
				continue
			}
			commP, := commVal(valNeedSlice, valHaveSlice)
			if commP > max {
				max = commP
				staticMap
			}
		}
	}
}

func isInSlice(mySlice []string, check string) bool {
	for _, val := range mySlice {
		if strings.Compare(val, check) == 0 {
			return true
		}
	}
	return false
}

func commVal(slice1 []string, slice2 []string, staticSlice*[]string) (int) {
	var count int
	sliceCom := make([]string, 0, len(slice1))
	for _, valNeed := range slice1 {
		for _, valHave := range slice2 {
			if valNeed == valHave {
				count++
				sliceCom = append(sliceCom, valNeed)
			}
		}
	}
	return count, sliceCom
}


func 