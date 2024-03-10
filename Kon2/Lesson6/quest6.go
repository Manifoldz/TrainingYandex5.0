package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//первая строка - количество секторов n
	var n int
	_, err := fmt.Scanf("%d\n", &n)
	if err != nil || n < 3 || n > 100 {
		fmt.Print("Read-mistake1: ", err)
		return
	}

	//вторая строка последовательность секторов по часовой
	array := make([]int, 0, n)
	var num int
	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < n; i++ {
		if i == n-1 {
			_, err = fmt.Fscanf(reader, "%d\n", &num)
		} else {
			_, err = fmt.Fscanf(reader, "%d", &num)
		}
		if err != nil || num < 1 || num > 1000 {
			fmt.Print("Read-mistake2: ", err, num)
			return
		}
		array = append(array, num)
	}

	//третья строка
	var a, b, k int
	_, err = fmt.Fscanf(reader, "%d %d %d\n", &a, &b, &k)
	if err != nil || a < 1 || a > b || b < 1 || b > 1_000_000_000 || k < 1 || k > 1_000_000_000 {
		fmt.Print("Read-mistake3: ", err)
		return
	}

	//логика
	var max int

	if a%k == 0 {
		a = (a/k - 1)
	} else {
		a = (a / k)
	}

	if b%k == 0 {
		b = (b/k - 1)
	} else {
		b = (b / k)
	}

	if b-a >= n {
		a = 0
		b = n - 1
	}

	for i := a; i <= b; i++ {

		if array[i%n] > max {
			max = array[i%n]
		}
		if array[(n-i%n)%n] > max {
			max = array[(n-i%n)%n]
		}
	}
	fmt.Print(max)
}
