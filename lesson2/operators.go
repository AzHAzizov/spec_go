package main

import (
	"fmt"
	"strconv"
)

func main() {

	// ------------------------------------------------------------------------------------------------------

	// // зарашиваем у пользователя число и проверяем его на четность
	// var number int
	// fmt.Println("Please enter number: ")
	// fmt.Scan(&number)

	// if number%2 == 0 {
	// 	fmt.Println("Число четное")
	// } else {
	// 	fmt.Println("Число НЕ четное")
	// }

	// ------------------------------------------------------------------------------------------------------

	// // сравниваем цвета который ввел пользователь
	// var color string
	// fmt.Println("Can you please enter your hair color: ")
	// fmt.Scan(&color)

	// if strings.Compare(color, "green") == 0 {
	// 	fmt.Println("Congrotulations green hair mans win the game")
	// } else if strings.Compare(color, "brown") == 0 {
	// 	fmt.Println("We recommend you paint your hair color to green")
	// } else {
	// 	fmt.Printf("Hair color %v is not play in our tournoment", color)
	// }

	// ------------------------------------------------------------------------------------------------------

	// // обявление переменного сразу внутри условии
	// if age := askUserAge(); age > 18 {
	// 	fmt.Println("Welcome to party")
	// } else {
	// 	fmt.Println("You can not enter to party")
	// }

	// ------------------------------------------------------------------------------------------------------

	// // не рабочий стек кода
	// if false {
	// 	fmt.Println("Line is true")
	// }
	// else { // вот так вниз отпускать запрещается, потому что компелятор за нас подставит в 53 строке ;
	// 	fmt.Println("Line is false");
	// }

	// ------------------------------------------------------------------------------------------------------

	return

}

func askUserAge() int {
	var age string
	fmt.Println("Please enter your age: ")
	fmt.Scan(&age)

	ageInt, err := strconv.Atoi(age)
	if err == nil {
		return ageInt
	}
	return 0
}
