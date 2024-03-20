package main

import (
	"fmt"
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
	max := 0

	//запускаем обход с каждой точки и сохраняем самую длинный обход
	for key := range haveMap {
		num := startMatch(key, haveMap, needMap)
		if num > max {
			max = num
		}
	}

	fmt.Print(n - max)

}

func startMatch(key string, haveMap map[string][]string, needMap map[string][]string) int {
	max := 0
	for _, valSlice := range haveMap[key] {
		for key2 := range needMap {
			for _, valSlice2 := range needMap[key2] {
				if valSlice == valSlice2 {
					//var backupNeed map[string][]string
					//copy(backupHave, haveMap)
					//copy(backupNeed, needMap)
					banMap1 := make(map[string]int, len(haveMap))
					banMap2 := make(map[string]int, len(needMap))
					banMap1[key]++
					banMap2[key2]++
					num := countMatch(key, key2, valSlice, haveMap, needMap, banMap1, banMap2)
					if num > max {
						max = num
					}
				}
			}
		}
	}
	return max
}

func countMatch(key1, key2, valSlice string, haveMap map[string][]string, needMap map[string][]string, banMap1 map[string]int, banMap2 map[string]int) int {
	//расшифруем точки
	var dx, dy int
	fmt.Sscanf(valSlice, "dx%ddy%d", &dx, &dy)
	var x1, y1, x2, y2 int
	fmt.Sscanf(key1, "x%dy%d", &x1, &y1)
	fmt.Sscanf(key2, "x%dy%d", &x2, &y2)

	key1 = fmt.Sprintf("x%dy%d", x1+dx, y1+dy)
	key2 = fmt.Sprintf("x%dy%d", x2+dx, y2+dy)

	if countTheSame(haveMap[key1]) == banMap1[key1] {
		return 0
	}

	if countTheSame(needMap[key2]) == banMap2[key2] {
		return 0
	}
	banMap1[key1]++
	banMap2[key2]++
	num := 1
	for _, valSlice1 := range haveMap[key1] {
		for _, valSlice2 := range needMap[key2] {
			if valSlice1 == valSlice2 {
				num += countMatch(key1, key2, valSlice1, haveMap, needMap, banMap1, banMap2)
			}
		}
	}
	return num
}

func countTheSame(srcSlice []string) int {
	max := 1
	var tempMap = make(map[string]int)
	for _, value := range srcSlice {
		tempMap[value]++
	}
	for _, value := range tempMap {
		if value > max {
			max = value
		}
	}
	return max
}
