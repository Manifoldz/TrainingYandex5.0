package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	fmt.Scanf("%d\n", &n)
	//имеющиеся
	var x1, y1, x2, y2 int
	haveMap := make(map[string][]string)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d %d %d %d\n", &x1, &y1, &x2, &y2)
		p1 := fmt.Sprintf("x%dy%d", x1, y1)
		p2 := fmt.Sprintf("x%dy%d", x2, y2)

		haveMap[p1] = append(haveMap[p1], fmt.Sprintf("dx%ddy%d", x2-x1, y2-y1))
		haveMap[p2] = append(haveMap[p2], fmt.Sprintf("dx%ddy%d", x1-x2, y1-y2))
	}
	var col = 0
	//необходимые
	needMap := make(map[string][]string)
	for i := 0; i < n; i++ {
		var x1, y1, x2, y2 int
		fmt.Scanf("%d %d %d %d\n", &x1, &y1, &x2, &y2)
		p1 := fmt.Sprintf("x%dy%d", x1, y1)
		p2 := fmt.Sprintf("x%dy%d", x2, y2)
		if _, ok := haveMap[p1]; ok {
			if isInSlice(haveMap, p1, fmt.Sprintf("dx%ddy%d", x2-x1, y2-y1), true) {
				if _, ok := haveMap[p2]; ok {
					if isInSlice(haveMap, p2, fmt.Sprintf("dx%ddy%d", x1-x2, y1-y2), true) {
						col++
						continue
					}
				}
			}
		}
		needMap[p1] = append(needMap[p1], fmt.Sprintf("dx%ddy%d", x2-x1, y2-y1))
		needMap[p2] = append(needMap[p2], fmt.Sprintf("dx%ddy%d", x1-x2, y1-y2))
	}
	max := 0
	//theSame(haveMap, needMap)

	//запускаем обход с каждой точки и сохраняем самую длинный обход
	for key := range haveMap {
		num := startMatch(key, haveMap, needMap)
		if num > max {
			max = num
		}
	}

	if col > max {
		max = col
	}

	fmt.Print(n - max)

}

func startMatch(key string, haveMap map[string][]string, needMap map[string][]string) (max int) {
	var countTheSame int
	for _, valSlice := range haveMap[key] {
		for key2 := range needMap {
			for _, valSlice2 := range needMap[key2] {
				if valSlice == valSlice2 {
					banMap1 := make(map[string]bool, len(haveMap))
					banMap2 := make(map[string]bool, len(needMap))
					num := countMatch(key, key2, valSlice, haveMap, needMap, banMap1, banMap2)
					if num > max {
						countTheSame = countSame(haveMap, needMap, key, key2)
						max = num
					} else if num == max {
						newTheSame := countSame(haveMap, needMap, key, key2)
						if newTheSame > countTheSame {
							max = num
							countTheSame = newTheSame
						}
					}
				}
			}
		}
	}
	return max + countTheSame
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
					fmt.Sscanf(val, "dx%ddy%d", &x2, &y2)
					p2 := fmt.Sprintf("x%dy%d", x1+x2+dX, y1+y2+dY)
					if _, ok := needMap[p2]; ok {
						if isInSlice(needMap, p2, fmt.Sprintf("dx%ddy%d", x1+dX, y1+dY), false) {
							count++
							continue
						}
					}
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
