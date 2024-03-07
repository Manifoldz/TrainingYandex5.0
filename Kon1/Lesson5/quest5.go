package main

import (
	"fmt"
	//"time"
)

func main() {
	//startTime := time.Now()
	var n, k, d uint64
	_, err := fmt.Scanf("%d %d %d", &n, &k, &d)
	if err != nil || n < 1 || k < 1 || d < 1 || n > 1000000000 || k > 1000000000 || d > 100000 {
		fmt.Println("Read-mistake:", err)
		return
	}
	flag := true
	var i uint64
	var num0 uint64
	for i < d && flag {
		n *= 10
		flag = false
		for j := 0; j <= 9; j++ {
			if (n+uint64(j))%k == 0 {
				flag = true
				n = n + uint64(j)
				break
			}
		}
		if n%10 == 0 {
			num0 = d - i - 1
			break
		}
		i++
	}

	if flag {
		fmt.Print(n)
		for num0 > 0 {
			fmt.Print(0)
			num0--
		}
	} else {
		fmt.Print(-1)
	}
	// endTime := time.Since(startTime)
	// fmt.Println("Time:", endTime)
}
