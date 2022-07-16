package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {


	// Link: https://contest.yandex.ru/contest/25667/problems/B/

	input := bufio.NewScanner(os.Stdin)
	if input.Scan() {
		inputText := input.Text()
		if inputText == "" {
			fmt.Println("НЕ СОГЛАСЕН")
			return
		}
		textRune := []rune(inputText)
		var constructedRune []rune
		constructedRune = append(constructedRune, textRune[0])
		constructedRune = append(constructedRune, textRune[len(textRune)-1])

		constructedString := strings.ToLower(string(constructedRune))

		if constructedString == "да" {
			fmt.Println("СОГЛАСЕН")
			return
		}

		fmt.Println("НЕ СОГЛАСЕН")
	}
}
