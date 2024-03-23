package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	var n int
	startTime := time.Now()
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d\n", &n)
	//имеющиесяs
	var x1, y1, x2, y2 int
	haveMap := make(map[string][]string)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d %d %d %d\n", &x1, &y1, &x2, &y2)
		if x1 > x2 {
			x2, x1 = x1, x2
			y2, y1 = y1, y2
		} else if y2 < y1 {
			y2, y1 = y1, y2
		}
		p1 := fmt.Sprintf("x%dy%d", x1, y1)
		haveMap[p1] = append(haveMap[p1], fmt.Sprintf("dx%ddy%d", x2-x1, y2-y1))
	}
	var col = 0
	//необходимые
	needMap := make(map[string][]string)
	for i := 0; i < n; i++ {
		var x1, y1, x2, y2 int
		fmt.Fscanf(reader, "%d %d %d %d\n", &x1, &y1, &x2, &y2)
		if x1 > x2 {
			x2, x1 = x1, x2
			y2, y1 = y1, y2
		} else if y2 < y1 {
			y2, y1 = y1, y2
		}
		p1 := fmt.Sprintf("x%dy%d", x1, y1)
		if _, ok := haveMap[p1]; ok {
			if isInSlice(haveMap, p1, fmt.Sprintf("dx%ddy%d", x2-x1, y2-y1), true) {
				col++
				continue
			}
		}
		needMap[p1] = append(needMap[p1], fmt.Sprintf("dx%ddy%d", x2-x1, y2-y1))
	}
	max := 0
	//запускаем обход с каждой точки и сохраняем самую длинный обход
	var count int
	for key := range haveMap {
		num, countTheSame := startMatch(key, haveMap, needMap)
		if num > max {
			max = num
			count = countTheSame
		} else if num == max {
			if count < countTheSame {
				count = countTheSame
			}
		}
	}

	fmt.Print(n - max - col - count)
	diffTime := time.Since(startTime)
	fmt.Print(diffTime)
}

func startMatch(key string, haveMap map[string][]string, needMap map[string][]string) (max int, countTheSame int) {
	for _, valSlice := range haveMap[key] {
		for key2 := range needMap {
			for _, valSlice2 := range needMap[key2] {
				if valSlice == valSlice2 {
					banMap1 := make(map[string]bool, len(haveMap))
					banMap2 := make(map[string]bool, len(needMap))
					num := countMatch(key, key2, valSlice, haveMap, needMap, banMap1, banMap2)
					if num > max {
						countTheSame = countSame(haveMap, needMap, key, key2) - num
						max = num
					} else if num == max {
						newTheSame := countSame(haveMap, needMap, key, key2) - num
						if newTheSame > countTheSame {
							countTheSame = newTheSame
						}
					}
				}
			}
		}
	}
	return max, countTheSame
}

func countSame(haveMap map[string][]string, needMap map[string][]string, key1, key2 string) int {
	var dX, dY int
	var count int
	var x1, y1, x2, y2 int
	fmt.Sscanf(key1, "x%dy%d", &x1, &y1)
	fmt.Sscanf(key2, "x%dy%d", &x2, &y2)
	dX = x2 - x1
	dY = y2 - y1

	for key, valueSlice := range haveMap {
		fmt.Sscanf(key, "x%dy%d", &x1, &y1)
		p1 := fmt.Sprintf("x%dy%d", x1+dX, y1+dY)
		if _, ok := needMap[p1]; ok {
			for _, val := range valueSlice {
				if isInSlice(needMap, p1, val, false) {
					count++
					continue
				}
			}
		}
	}
	return count
}

func countMatch(key1, key2, valSlice string, haveMap map[string][]string, needMap map[string][]string, banMap1 map[string]bool, banMap2 map[string]bool) int {
	//расшифруем точки
	var dx, dy int
	fmt.Sscanf(valSlice, "dx%ddy%d", &dx, &dy)
	var x1, y1, x2, y2 int
	fmt.Sscanf(key1, "x%dy%d", &x1, &y1)
	fmt.Sscanf(key2, "x%dy%d", &x2, &y2)

	banMap1[key1] = true
	banMap2[key2] = true

	key1 = fmt.Sprintf("x%dy%d", x1+dx, y1+dy)
	key2 = fmt.Sprintf("x%dy%d", x2+dx, y2+dy)

	if _, ok := banMap1[key1]; ok {
		return 0
	}

	if _, ok := banMap2[key2]; ok {
		return 0
	}
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

func isInSlice(mapDest map[string][]string, p1 string, check string, change bool) bool {
	for idx, val := range mapDest[p1] {
		if strings.Compare(val, check) == 0 {
			if change {
				if len(mapDest[p1]) == 1 {
					delete(mapDest, p1)
				} else if len(mapDest[p1])-1 != idx {
					mapDest[p1] = append(mapDest[p1][:idx], mapDest[p1][idx+1:]...)
				} else {
					mapDest[p1] = mapDest[p1][:idx]
				}
			}
			return true
		}
	}
	return false
}
