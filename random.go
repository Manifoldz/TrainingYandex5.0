package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i < 100000; i++ {
		fmt.Print(rand.Intn(2000000001)-1000000000, " ")
	}
}
