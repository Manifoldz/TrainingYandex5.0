package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//startTime := time.Now()
	var n int
	reader := bufio.NewReader(os.Stdin)
	_, err := fmt.Fscanf(reader, "%d\n", &n)
	if err != nil || n < 1 || n > 2000 {
		fmt.Printf("Read-mistake1 err: %v, n: %d", err, n)
		return
	}

	myMap := make(map[int][]int, n)
	var x4, y4, x5, y5 int

	for i := 0; i < n; i++ {
		var x, y int
		_, err = fmt.Fscanf(reader, "%d %d\n", &x, &y)
		if err != nil || x < -1_000_000_00 || y < -1_000_000_00 || x > 1_000_000_00 || y > 1_000_000_00 {
			fmt.Printf("Read-mistake2 err: %v, x: %d, y: %d", err, x, y)
			return
		}

		_, ok := myMap[x]

		if !ok {
			myMap[x] = append(myMap[x], y)
		} else if !isInSlice(myMap[x], y) {
			myMap[x] = append(myMap[x], y)
		}

		if i == 1 {
			x4, y5 = x, y
		} else if i == 0 {
			x5, y4 = x, y
		}
	}
	if len(myMap) == 1 {
		var key int
		var slice1 []int
		for key, slice1 = range myMap {
		}
		if len(slice1) == 1 {
			fmt.Println(3)
			fmt.Println(key+1, slice1[0])
			fmt.Println(key+1, slice1[0]+1)
			fmt.Println(key, slice1[0]+1)
			return
		}
	}
	var printX, printY int
	var find3 bool
	for i, slice1 := range myMap {
		for k := 0; k < len(slice1); k++ {
			x1 := i
			y1 := slice1[k]
			for j, slice2 := range myMap {
				var l int
				if i == j {
					if k == len(slice1)-1 {
						continue
					} else {
						l = k + 1
					}
				}
				x2 := j
				sum3 := x1 + x2 + y1
				ded3 := x1 + x2 - y1
				//итерация по столбцу соседнему
				for ; l < len(slice2); l++ {
					y2 := slice2[l]
					if (sum3+y2)%2 != 0 || (ded3-y2)%2 != 0 {
						continue
					}
					sum3 = (sum3 + y2) / 2
					ded3 = (ded3 - y2) / 2
					x4, y4 = ded3+y1, sum3-x1
					x5, y5 = ded3+y2, sum3-x2
					slice4, ok1 := myMap[x4]
					slice5, ok2 := myMap[x5]
					if !find3 {
						if ok1 {
							if isInSlice(slice4, y4) {
								printX = x5
								printY = y5
								find3 = true
							}
						}
						if ok2 {
							if isInSlice(slice5, y5) {
								if printX == x5 && printY == y5 {
									fmt.Println(0)
									return
								} else {
									printX = x4
									printY = y4
									find3 = true
								}
							}
						}
					} else {
						if ok1 && ok2 && isInSlice(slice4, y4) && isInSlice(slice5, y5) {
							fmt.Println(0)
							return
						}
					}
				}
			}

		}
		delete(myMap, i)
	}
	if find3 {
		fmt.Println(1)
		fmt.Println(printX, printY)
	} else {
		fmt.Println(2)
		fmt.Println(x4, y4)
		fmt.Println(x5, y5)
	}
	//diffTime := time.Since(startTime)
	//fmt.Println(diffTime)
}

func isInSlice(mySlice []int, check int) bool {
	for _, val := range mySlice {
		if val == check {
			return true
		}
	}
	return false
}
