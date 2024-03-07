package main

import "fmt"

func main() {
	var numSold, hpBarack, enemySold, numTurn int64
	_, err := fmt.Scanf("%d\n%d\n%d\n", &numSold, &hpBarack, &enemySold)
	if err != nil || numSold < 1 || hpBarack < 1 || enemySold < 1 || numSold > 5000 || hpBarack > 5000 || enemySold > 5000 {
		fmt.Print("Read-mistake:", err)
		return
	}
	if numSold >= hpBarack {
		fmt.Print(1)
		return
	}
	hpBarack -= numSold
	numTurn++

	if numSold <= enemySold && hpBarack >= numSold {
		fmt.Print(-1)
		return
	}

	numTurn = countTurn(numSold, hpBarack, enemySold, numTurn)
	if numTurn == 10000 {
		numTurn = -1
	}
	fmt.Print(numTurn)
}

func countTurn(numSold, hpBarack, enemySold, numTurn int64) int64 {
	if hpBarack < 0 {
		hpBarack = 0
	}
	if enemySold <= 0 && hpBarack <= 0 {
		return numTurn
	}
	if numSold <= 0 {
		return 10000
	}

	//fmt.Println("Turn:", numTurn)
	//сломать барак
	var ver1 int64 = 10000
	if numSold >= hpBarack {
		ver1 = countTurn(numSold-(enemySold-(numSold-hpBarack)), 0, enemySold-(numSold-hpBarack), numTurn+1)
	}

	//не ломать барак
	var ver2 int64 = 10000
	if hpBarack > 0 && numSold > enemySold {
		ver2 = countTurn(numSold, hpBarack-(numSold-enemySold), enemySold, numTurn+1)
	}

	if ver1 < ver2 {
		return ver1
	}

	return ver2
}
