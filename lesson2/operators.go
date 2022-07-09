package main

import (
	"fmt"
	"strings"
)

func main() {

	// // зарашиваем у пользователя число и проверяем его на четность
	// var number int
	// fmt.Println("Please enter number: ")
	// fmt.Scan(&number)

	// if number%2 == 0 {
	// 	fmt.Println("Число четное")
	// } else {
	// 	fmt.Println("Число НЕ четное")
	// }


	// сравниваем цвета который ввел пользователь 
	var color string
	fmt.Println("Can you please enter your hair color: ")
	fmt.Scan(&color)

	if strings.Compare(color, "green") == 0 {
		fmt.Println("Congrotulations green hair mans win the game")
	} else if strings.Compare(color, "brown") == 0 {
		fmt.Println("We recommend you paint your hair color to green")
	} else {
		fmt.Printf("Hair color %v is not play in our tournoment", color)
	}

}
