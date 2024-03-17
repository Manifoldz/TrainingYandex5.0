package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	buffer := make([]byte, 100001) // установка буфера для больших строк
	scanner.Buffer(buffer, 100001)
	scanner.Scan()
	text := scanner.Text()
	dict := make(map[rune]int, 30)
	textRune := []rune(text)
	if len(textRune) > 100000 || len(textRune) < 1 {
		fmt.Print("Mistake1")
		return
	}
	for i := 0; i < len(textRune); i++ {
		if textRune[i] < 'a' || textRune[i] > 'z' {
			fmt.Print("Mistake2")
			return
		}
		dict[textRune[i]]++
	}
	scanner.Scan()
	text = scanner.Text()
	textRune = []rune(text)
	for i := 0; i < len(textRune); i++ {
		value, ok := dict[textRune[i]]
		if ok && value > 0 {
			dict[textRune[i]]--
		} else {
			fmt.Print("NO")
			return
		}
	}

	for _, value := range dict {
		if value != 0 {
			fmt.Print("NO")
			return
		}
	}
	fmt.Print("YES")
}
