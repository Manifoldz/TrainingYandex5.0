package main

import (
	"fmt"
	"math"
)

func main() {
	var K int8
	_, err := fmt.Scanf("%d\n", &K)
	if err != nil || K < 1 || K > 100 {
		fmt.Print("Read-mistake: ", err)
		return
	}

	var max1, max2, min1, min2 int64
	_, err = fmt.Scanf("%d %d\n", &max1, &max2)
	if err != nil || int64(math.Abs(float64(max2))) > 1000000000 || int64(math.Abs(float64(max1))) > 1000000000 {
		fmt.Print("Read-mistake2: ", err)
		return
	}

	min1, min2 = max1, max2
	var temp1, temp2 int64
	for i := 0; int8(i) < K-1; i++ {
		_, err = fmt.Scanf("%d %d\n", &temp1, &temp2)
		if err != nil || int64(math.Abs(float64(temp2))) > 1000000000 || int64(math.Abs(float64(temp1))) > 1000000000 {
			fmt.Print("Read-mistake2: ", err)
			return
		}
		if temp1 > max1 {
			max1 = temp1
		} else if temp1 < min1 {
			min1 = temp1
		}

		if temp2 > max2 {
			max2 = temp2
		} else if temp2 < min2 {
			min2 = temp2
		}
	}

	fmt.Print(min1, min2, max1, max2)
}
