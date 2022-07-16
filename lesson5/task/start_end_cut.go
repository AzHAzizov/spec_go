package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewScanner(os.Stdin)

	if input.Scan() {
		text := input.Text()
		textAsRune := []rune(text)

		var removePostion string = "start"

		for {

			if len(textAsRune) <= 2 {
				if len(textAsRune) == len([]rune(text)) {
					fmt.Println(text)
				}
				break
			}

			if removePostion == "start" {
				textAsRune = textAsRune[2:]
				removePostion = "end"
			} else if removePostion == "end" {
				textAsRune = textAsRune[:len(textAsRune)-2]
				removePostion = "start"
			}

			fmt.Println(string(textAsRune))

		}

	}
}
