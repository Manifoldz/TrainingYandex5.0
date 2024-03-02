package main

import "fmt"

func main() {
	var treeBucketVasya, treeBucketMasha, limitVasya, limitMasha int32
	_, err := fmt.Scanf("%d %d\n%d %d\n", &treeBucketVasya, &limitVasya, &treeBucketMasha, &limitMasha)
	if err != nil {
		fmt.Println("n/a")
		return
	}
	min1, max1 := myMinMax(treeBucketVasya-limitVasya, treeBucketVasya+limitVasya)
	min2, max2 := myMinMax(treeBucketMasha-limitMasha, treeBucketMasha+limitMasha)
	if min2 > max1 || min1 > max2 {
		fmt.Print(max1 - min1 + 1 + max2 - min2 + 1)
	} else {
		fmt.Print(myMax(max1, max2) - myMin(min1, min2) + 1)
	}

}

func myMax(a int32, b int32) int32 {
	if a >= b {
		return a
	}
	return b
}

func myMin(a int32, b int32) int32 {
	if a <= b {
		return a
	}
	return b
}

func myMinMax(a int32, b int32) (int32, int32) {
	if a <= b {
		return a, b
	}
	return b, a
}
