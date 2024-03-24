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
	haveMap := make(map[point][]direction, n)
	var p1, p2 point
	for i := 0; i < n; i++ {
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
	//необходимые
	var col int
	needMap := make(map[point][]direction, n)
	for i := 0; i < n; i++ {
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
	max := startMatch(haveMap, needMap)

	if col > max {
		max = col
	}

	fmt.Print(n - max)
	diffTime := time.Since(startTime)
	fmt.Println()
	fmt.Print(diffTime)
}

func startMatch(haveMap map[point][]direction, needMap map[point][]direction) int {
	var max int
	banMap1 := make(map[int]map[int]bool, len(haveMap))
	for key := range haveMap {
		for key2 := range needMap {
			if banMap1[key2.x-key.x] != nil {
				if banMap1[key2.x-key.x][key2.y-key.y] {
					continue
				}
			}
			num := countSame(haveMap, needMap, key, key2)
			if num > max {
				max = num
			}
			if banMap1[key2.x-key.x] == nil {
				banMap1[key2.x-key.x] = make(map[int]bool)
			}
			banMap1[key2.x-key.x][key2.y-key.y] = true
		}
	}

	return max
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
