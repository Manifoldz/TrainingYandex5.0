package main

import (
	"fmt"
)

func main() {
	var n, up, down uint64
	fmt.Scan(&n)
	row := lbinArSearch(0, n)    // в каком диагональном ряду лежит значение
	offset := arithProg(row) - n // на сколько отстает от максимума в ряду

	if row%2 == 0 {
		down = row - offset
		up = offset + 1
	} else {
		up = row - offset
		down = offset + 1
	}

	fmt.Printf("%d/%d", up, down)
}

func lbinArSearch(first, val uint64) uint64 {
	var last = val
	for first < last {
		mid := (first + last) / 2
		if isMore(mid, val) {
			last = mid
		} else {
			first = mid + 1
		}
	}
	return first
}

func isMore(check, val uint64) bool {
	if check > 4_000_000_000 {
		//fmt.Println("To many")
		return true
	}
	max := arithProg(check)

	return max >= val
}

func arithProg(val uint64) uint64 {
	return (1 + val) * val / 2
}
