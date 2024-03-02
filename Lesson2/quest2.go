package main

import "fmt"

func main() {
	var scoreF1, scoreF2, scoreS1, scoreS2 int8
	var is1atHome int8 //1-Yes, 2-No
	if _, err := fmt.Scanf("%d:%d\n%d:%d\n%d", &scoreF1, &scoreF2, &scoreS1, &scoreS2, &is1atHome); err != nil {
		fmt.Print("Read-mistake")
		return
	} else if !isValidScore(scoreF1) || !isValidScore(scoreF2) || !isValidScore(scoreS1) || !isValidScore(scoreS2) || (is1atHome != 1 && is1atHome != 2) {
		fmt.Print("Not valid input")
		return
	}
	//если итак выигрывает нужно 0
	if scoreF1+scoreS1 > scoreF2+scoreS2 {
		fmt.Print(0)
		return
	}
	// если в гостях забила больше, главное сравнять счет
	if (scoreF1 > scoreS2 && is1atHome == 2) || (scoreS1 > scoreF2 && is1atHome == 1) {
		fmt.Print(scoreF2 + scoreS2 - scoreF1 - scoreS1)
		return
	}

	fmt.Print(scoreF2 + scoreS2 - scoreF1 - scoreS1 + 1)
	return

}

func isValidScore(score int8) bool {
	if score < 0 || score > 5 {
		return false
	}
	return true
}
