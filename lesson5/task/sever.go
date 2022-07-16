package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewScanner(os.Stdin)

	// for {

	if input.Scan() {

		text := input.Text()
		if text == "" {
			return
		}

		textAsRune := []rune(text)
		var combinedTextAsRune []rune

		for key, value := range textAsRune {
			if key%2 == 0 {
				combinedTextAsRune = append(combinedTextAsRune, value, value, value)
			}
		}

		fmt.Println(string(combinedTextAsRune))

	}
	// }
}
