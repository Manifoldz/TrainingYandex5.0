package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type point struct {
	x int
	y int
}

type direction struct {
	dx int
	dy int
}

func main() {
	var n int
	startTime := time.Now()
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d\n", &n)
	//имеющиеся
	haveMap := make(map[point][]direction)
	for i := 0; i < n; i++ {
		var p1, p2 point
		fmt.Fscanf(reader, "%d %d %d %d\n", &p1.x, &p1.y, &p2.x, &p2.y)
		if p1.x > p2.x {
			p2.x, p1.x = p1.x, p2.x
			p2.y, p1.y = p1.y, p2.y
		} else if p2.y < p1.y && p1.x == p2.x {
			p2.y, p1.y = p1.y, p2.y
		}
		dir := direction{p2.x - p1.x, p2.y - p1.y}
		haveMap[p1] = append(haveMap[p1], dir)
	}
	var col = 0
	//необходимые
	needMap := make(map[point][]direction)
	for i := 0; i < n; i++ {
		var p1, p2 point
		fmt.Fscanf(reader, "%d %d %d %d\n", &p1.x, &p1.y, &p2.x, &p2.y)
		if p1.x > p2.x {
			p2.x, p1.x = p1.x, p2.x
			p2.y, p1.y = p1.y, p2.y
		} else if p2.y < p1.y && p1.x == p2.x {
			p2.y, p1.y = p1.y, p2.y
		}
		dir := direction{p2.x - p1.x, p2.y - p1.y}
		if _, ok := haveMap[p1]; ok {
			if isInSlice(haveMap, p1, dir, true) {
				col++
				continue
			}
		}
		needMap[p1] = append(needMap[p1], dir)
	}
	max := 0
	//запускаем обход с каждой точки и сохраняем самый длинный обход
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
	fmt.Println()
	fmt.Print(diffTime)
}

func startMatch(key point, haveMap map[point][]direction, needMap map[point][]direction) (max int, countTheSame int) {
	for _, valSlice := range haveMap[key] {
		for key2 := range needMap {
			for _, valSlice2 := range needMap[key2] {
				if valSlice == valSlice2 {
					banMap1 := make(map[point]bool, len(haveMap))
					banMap2 := make(map[point]bool, len(needMap))
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

func countSame(haveMap map[point][]direction, needMap map[point][]direction, key1, key2 point) int {
	var count int
	dX := key2.x - key1.x
	dY := key2.y - key1.y

	for key, valueSlice := range haveMap {
		p1 := point{key.x + dX, key.y + dY}
		if _, ok := needMap[p1]; ok {
			for _, val := range valueSlice {
				if isInSlice(needMap, p1, val, false) {
					count++
				}
			}
		}
	}
	return count
}

func countMatch(key1, key2 point, valSlice direction, haveMap map[point][]direction, needMap map[point][]direction, banMap1 map[point]bool, banMap2 map[point]bool) int {
	//расшифруем точки
	banMap1[key1] = true
	banMap2[key2] = true

	newKey1 := point{key1.x + valSlice.dx, key1.y + valSlice.dy}
	newKey2 := point{key2.x + valSlice.dx, key2.y + valSlice.dy}

	if _, ok := banMap1[newKey1]; ok {
		return 0
	}

	if _, ok := banMap2[newKey2]; ok {
		return 0
	}
	num := 1
	for _, valSlice1 := range haveMap[newKey1] {
		for _, valSlice2 := range needMap[newKey2] {
			if valSlice1 == valSlice2 {
				num += countMatch(newKey1, newKey2, valSlice1, haveMap, needMap, banMap1, banMap2)
			}
		}
	}
	return num
}

func isInSlice(mapDest map[point][]direction, p1 point, srcDir direction, change bool) bool {
	for idx, dir := range mapDest[p1] {
		if dir == srcDir {
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
