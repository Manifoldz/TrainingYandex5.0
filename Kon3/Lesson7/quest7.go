package main

import (
	"bufio"
	"fmt"
	"os"
)

type myStruct struct {
	key   int
	array []int
}

func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	_, err := fmt.Fscanf(reader, "%d\n", &n)
	if err != nil || n < 1 || n > 2000 {
		fmt.Printf("Read-mistake1 err: %v, n: %d", err, n)
		return
	}

	myMatrix := make([]myStruct, 0, n)
	var x4, y4, x5, y5 int

	for i := 0; i < n; i++ {
		var x, y int
		_, err = fmt.Fscanf(reader, "%d %d\n", &x, &y)
		if err != nil || x < -1_000_000_00 || y < -1_000_000_00 || x > 1_000_000_00 || y > 1_000_000_00 {
			fmt.Printf("Read-mistake2 err: %v, x: %d, y: %d", err, x, y)
			return
		}

		index := isInMatrix(myMatrix, x)

		if index == -1 {
			var newStruct myStruct
			newStruct.key = x
			newStruct.array = append(newStruct.array, y)
			myMatrix = append(myMatrix, newStruct)
		} else if !isInSlice(myMatrix[index].array, y) {
			myMatrix[index].array = append(myMatrix[index].array, y)
		}

		if i == 1 {
			x4, y5 = x, y
		} else if i == 0 {
			x5, y4 = x, y
		}
	}
	if len(myMatrix) == 1 && len(myMatrix[0].array) == 1 {
		fmt.Println(3)
		fmt.Println(myMatrix[0].key+1, myMatrix[0].key)
		fmt.Println(myMatrix[0].key+1, myMatrix[0].key+1)
		fmt.Println(myMatrix[0].key, myMatrix[0].key+1)
		return
	}
	var printX, printY int
	var find3 bool
	for i := 0; i < len(myMatrix)-1; i++ {
		for k := 0; k < len(myMatrix[i].array); k++ {
			x1 := myMatrix[i].key
			y1 := myMatrix[i].array[k]
			for j := i + 1; j < len(myMatrix); j++ {
				x2 := myMatrix[j].key
				//итерация по столбцу соседнему
				for l := 0; l < len(myMatrix[j].array); l++ {
					y2 := myMatrix[j].array[l]
					if (x1+x2+y1+y2)%2 != 0 || (x1+x2-y1-y2)%2 != 0 {
						continue
					}
					sum3 := (x1 + x2 + y1 + y2) / 2
					ded3 := (x1 + x2 - y1 - y2) / 2
					x4, y4 = ded3+y1, sum3-x1
					x5, y5 = ded3+y2, sum3-x2
					index4 := isInMatrix(myMatrix, x4)
					index5 := isInMatrix(myMatrix, x5)
					if !find3 {
						if index4 != -1 {
							if isInSlice(myMatrix[index4].array, y4) {
								printX = x5
								printY = y5
								find3 = true
							}
						}
						if index5 != -1 {
							if isInSlice(myMatrix[index5].array, y5) {
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
						if index4 != -1 && index5 != -1 && isInSlice(myMatrix[index4].array, y4) && isInSlice(myMatrix[index5].array, y5) {
							fmt.Println(0)
							return
						}
					}
				}
			}

		}
	}
	if find3 {
		fmt.Println(1)
		fmt.Println(printX, printY)
	} else {
		fmt.Println(2)
		fmt.Println(x4, y4)
		fmt.Println(x5, y5)
	}
}

func isInSlice(mySlice []int, check int) bool {
	for _, val := range mySlice {
		if val == check {
			return true
		}
	}
	return false
}

func isInMatrix(myMatrix []myStruct, check int) int {
	for i := 0; i < len(myMatrix); i++ {
		if myMatrix[i].key == check {
			return i
		}
	}
	return -1
}
