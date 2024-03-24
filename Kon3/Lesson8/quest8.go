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
	haveMap := make(map[point]map[direction]bool, n)
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
		if haveMap[p1] == nil {
			haveMap[p1] = make(map[direction]bool)
		}
		haveMap[p1][dir] = true
	}
	//необходимые
	var col int
	needMap := make(map[point]map[direction]bool, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d %d %d %d\n", &p1.x, &p1.y, &p2.x, &p2.y)
		if p1.x > p2.x {
			p2.x, p1.x = p1.x, p2.x
			p2.y, p1.y = p1.y, p2.y
		} else if p2.y < p1.y && p1.x == p2.x {
			p2.y, p1.y = p1.y, p2.y
		}
		dir := direction{p2.x - p1.x, p2.y - p1.y}
		if haveMap[p1][dir] {
			delete(haveMap[p1], dir)
			col++
			continue
		}
		if needMap[p1] == nil {
			needMap[p1] = make(map[direction]bool)
		}
		needMap[p1][dir] = true
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

func startMatch(haveMap map[point]map[direction]bool, needMap map[point]map[direction]bool) int {
	var max int
	banMap1 := make(map[int]map[int]bool, len(haveMap))
	for key1 := range haveMap {
		for key2 := range needMap {
			dX := key2.x - key1.x
			dY := key2.y - key1.y

			if banMap1[dX] != nil {
				if banMap1[dX][dY] {
					continue
				}
			}

			num := countSame(haveMap, needMap, dX, dY)
			if num > max {
				max = num
			}
			if banMap1[dX] == nil {
				banMap1[dX] = make(map[int]bool)
			}
			banMap1[dX][dY] = true
		}
	}

	return max
}

func countSame(haveMap, needMap map[point]map[direction]bool, dX, dY int) int {
	var count int

	for key, dirMap := range haveMap {
		p1 := point{key.x + dX, key.y + dY}
		if needMap[p1] != nil {
			for dirKey := range dirMap {
				if needMap[p1][dirKey] {
					count++
				}
			}
		}
	}
	return count
}
