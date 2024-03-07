package main

import (
	"bufio"
	"fmt"
	"os"
	//"time"
)

func main() {
	//startTime := time.Now
	scanner := bufio.NewScanner(os.Stdin)
	slice := make([][]byte, 8)
	for i := range slice {
		slice[i] = make([]byte, 8)
	}

	for i := 0; i < 8; i++ {
		scanner.Scan()
		err := scanner.Err()
		str := scanner.Text()
		err2 := scanner.Err()
		if err != nil || err2 != nil || len(str) < 8 {
			fmt.Println("Read-mistake:", err)
			return
		}

		for j := 0; j < 8; j++ {
			if str[j] != '*' && str[j] != 'B' && str[j] != 'R' {
				fmt.Print("Not allowed symbol inside 8*8")
				return
			}
			slice[i][j] = str[j]
		}
	}

	for row, str := range slice {
		for col, char := range str {
			if char == 'R' {
				//up
				offset := 1
				for row-offset >= 0 && slice[row-offset][col] != 'B' && slice[row-offset][col] != 'R' {
					slice[row-offset][col] = '#'
					offset++
				}
				//down
				offset = 1
				for row+offset <= 7 && slice[row+offset][col] != 'B' && slice[row+offset][col] != 'R' {
					slice[row+offset][col] = '#'
					offset++
				}
				//right
				offset = 1
				for col+offset <= 7 && slice[row][col+offset] != 'B' && slice[row][col+offset] != 'R' {
					slice[row][col+offset] = '#'
					offset++
				}
				//left
				offset = 1
				for col-offset >= 0 && slice[row][col-offset] != 'B' && slice[row][col-offset] != 'R' {
					slice[row][col-offset] = '#'
					offset++
				}
			} else if char == 'B' {
				//up-right
				offset := 1
				for row-offset >= 0 && col+offset <= 7 && slice[row-offset][col+offset] != 'B' && slice[row-offset][col+offset] != 'R' {
					slice[row-offset][col+offset] = '#'
					offset++
				}
				//up-left
				offset = 1
				for row-offset >= 0 && col-offset >= 0 && slice[row-offset][col-offset] != 'B' && slice[row-offset][col-offset] != 'R' {
					slice[row-offset][col-offset] = '#'
					offset++
				}
				//down-left
				offset = 1
				for row+offset <= 7 && col-offset >= 0 && slice[row+offset][col-offset] != 'B' && slice[row+offset][col-offset] != 'R' {
					slice[row+offset][col-offset] = '#'
					offset++
				}
				//down-right
				offset = 1
				for row+offset <= 7 && col+offset <= 7 && slice[row+offset][col+offset] != 'B' && slice[row+offset][col+offset] != 'R' {
					slice[row+offset][col+offset] = '#'
					offset++
				}
			}
		}
	}

	var numEmpty int
	for _, str := range slice {
		for _, char := range str {
			if char == '*' {
				numEmpty++
			}
		}
	}

	fmt.Print(numEmpty)
	//spendTime := time.Since(startTime())
	//fmt.Printf("Time:%s", spendTime)
}
