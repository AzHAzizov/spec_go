package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	input := bufio.NewScanner(os.Stdin)
	var lastWord string

	for {
		if input.Scan() {
			inputedText := string(input.Text())

			if inputedText == "" {
				fmt.Println("Game over")
				break
			}

			fmt.Println(inputedText)

			if lastWord != "" {

				lastWordAsRune := []rune(lastWord)
				lastWordLasRune := lastWordAsRune[len(lastWordAsRune)-1]

				correntWordAsRune := []rune(inputedText)
				currentWordLastRune := correntWordAsRune[0]

				fmt.Println(lastWord)

				if lastWordLasRune != currentWordLastRune {
					fmt.Println("Game over")
					break
				}

			}

			lastWord = inputedText
			fmt.Println("__________________________________________")

		}
	}

}
