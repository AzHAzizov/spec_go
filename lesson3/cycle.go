package main

import (
	"fmt"
	"strings"
)

func main() {

	// for init; condition; post {
	// init - блок инициализации переменных цикла
	// condition - условие (если верно - то тело цикла выполняется, если нет - то цикл завершается)
	// post - изменение переменной цикла (инкрементарное действие, декрементарное действие)
	// }

	/*
		for i := 0; i < 10; i++ {
			fmt.Printf("Curent number is %v \n", i)
		}
	*/

	//Важный момент : в качестве init может быть использовано выражение присваивания ТОЛЬКО через :=

	for i := 1; i < 1500; i *= 7 {
		if i%3 == 0 {
			continue
		}

		if i > 1300 {
			break
		}

		fmt.Printf("Current defficult number is : %v \n", i)
	}

	//Иногда бывает плохо. С лейблами по лучше. Лейблы - это синтаксический сахар

outer:
	for i := 0; i <= 2; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("i:%d and j:%d and sum i+j=%d\n", i, j, i+j)
			if i == j {
				break outer // Хочу чтобы вообще все циклы (внешние тоже остановились)
			}
		}
	}

	// while

	var loopVar int = 2
	// while (loopVar < 10) {
	// 	.....
	// 	loopVar++
	// }

	// for ; loopVar < 10; {
	for loopVar < 10 {
		fmt.Printf("Current number in while : %v \n", loopVar)
		loopVar++
	}

	// while true

	var password string
outer2:
	for {
		fmt.Print("Insert password: ")
		fmt.Scan(&password)
		if strings.Contains(password, "1234") {
			fmt.Println("Weak password . Try again")
		} else {
			fmt.Println("Password Accepted")
			break outer2
		}
	}


		//3. Цикл с множественными переменными цикла
		for x, y := 0, 1; x <= 10 && y <= 12; x, y = x+1, y+2 {
			fmt.Printf("%d + %d = %d\n", x, y, x+y)
		}
}
