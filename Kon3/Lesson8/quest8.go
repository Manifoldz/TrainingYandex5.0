package main

import (
	"bufio"
	"fmt"
	"os"
)

type point struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

type direction struct {
	dx int
	dy int
}

func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d\n", &n)

	//имеющиеся
	haveMap := make(map[point]bool, n)
	for i := 0; i < n; i++ {
		var p point
		fmt.Fscanf(reader, "%d %d %d %d\n", &p.x1, &p.y1, &p.x2, &p.y2)
		if p.x1 > p.x2 {
			p.x2, p.x1 = p.x1, p.x2
			p.y2, p.y1 = p.y1, p.y2
		} else if p.y2 < p.y1 && p.x1 == p.x2 {
			p.y2, p.y1 = p.y1, p.y2
		}
		haveMap[p] = true
	}

	//необходимые
	needMap := make(map[point]bool, n)
	for i := 0; i < n; i++ {
		var p point
		fmt.Fscanf(reader, "%d %d %d %d\n", &p.x1, &p.y1, &p.x2, &p.y2)
		if p.x1 > p.x2 {
			p.x2, p.x1 = p.x1, p.x2
			p.y2, p.y1 = p.y1, p.y2
		} else if p.y2 < p.y1 && p.x1 == p.x2 {
			p.y2, p.y1 = p.y1, p.y2
		}
		needMap[p] = true
	}

	max := countSame(haveMap, needMap)

	fmt.Print(n - max)
}

func countSame(haveMap, needMap map[point]bool) int {

	var counterMap = make(map[direction]int)

	for p1 := range haveMap {
		for p2 := range needMap {
			var dir direction
			dir.dx = p2.x2 - p1.x2
			dir.dy = p2.y2 - p1.y2
			if dir.dx == p2.x1-p1.x1 && dir.dy == p2.y1-p1.y1 {
				counterMap[dir]++
			}
		}
	}

	var max int
	for _, val := range counterMap {
		if val > max {
			max = val
		}
	}

	return max
}
