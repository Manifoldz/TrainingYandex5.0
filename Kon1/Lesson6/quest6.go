package main

import (
	"bufio"
	"fmt"
	"os"
	//"time"
)

func main() {
	//startTime := time.Now()
	var n int64
	reader := bufio.NewReader(os.Stdin)

	_, err := fmt.Fscanf(reader, "%d\n", &n)
	if err != nil || n < 2 || n > 100000 {
		fmt.Print("Read-mistake#01")
		return
	}

	var temp1 int64
	_, err = fmt.Fscanf(reader, "%d ", &temp1)
	if err != nil || temp1 < -1000000000 || temp1 > 1000000000 {
		fmt.Printf("Read-mistake#%d", 10)
		return
	}
	var str string
	for i := int64(1); i < n; i++ {
		var temp2 int64
		if i != n-1 {
			_, err = fmt.Fscanf(reader, "%d ", &temp2)
		} else {
			_, err = fmt.Fscanf(reader, "%d\n", &temp2)
		}

		if err != nil || temp1 < -1000000000 || temp1 > 1000000000 {
			fmt.Printf("Read-mistake#%d", i)
			return
		}

		if temp1%2 == 0 {
			str += "+"
			if temp2%2 != 0 {
				temp1 = 1
			}
		} else {
			if temp2%2 == 0 {
				str += "+"
			} else {
				str += "x"
			}
		}
	}

	fmt.Print(str)
	//diffTime := time.Since(startTime)
	//fmt.Print("Time:", diffTime)

}
