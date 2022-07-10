package main

import "fmt"

func main() {

	var array [5]int
	fmt.Println(array) // [0 0 0 0 0] дерет тот default который указали как тип, у нас он int по этому они 0

	var arrayStr [5]string
	fmt.Println(arrayStr) // [    ] а вот тут пустые строки

	arrayStr[0] = "A"
	arrayStr[1] = "B"
	arrayStr[2] = "C"
	arrayStr[3] = "D"
	arrayStr[4] = "E"

	fmt.Println(arrayStr) // [A B C D E]

	users := [5]string{"User1", "User2", "User3", "User4", "User5"}
	fmt.Println(users) // [User1 User2 User3 User4 User5]

	arrWithValues := [...]int{1, 2, 3, 4, 5, 6, 7, 8} // соберет значении их суш. данных
	fmt.Println(arrWithValues)                        // [1 2 3 4 5 6 7 8]

	//6. Массив и его размер - это две составляющие одного типа (Размер массив - часть типа)
	// var aArr [5]int
	// var bArr [6]int
	// aArr[0] = 100
	// bArr = aArr  -----> error

	// 7. Итерирование по массиву
	floatArr := [...]float64{12.5, 13.5, 15.2, 10.0, 12.0}
	for i := 0; i < len(floatArr); i++ {
		fmt.Printf("%d element of arr is %.2f\n", i, floatArr[i])
	}

	// 8. Итерирование по массиву через оператор range

	var sumOfFloarArr float64
	for id, value := range floatArr {
		fmt.Printf("elem with key %v has value %.2f \n", id, value)
		sumOfFloarArr += value
	}
	fmt.Printf("Sum of float64 list %v \n", sumOfFloarArr)

	// 9. Игнорирование id в range based for цикле
	for _, val := range floatArr {
		fmt.Printf("%.2f value WO id\n", val)
	}

	classes := [2][3]string{
		{"class A", "class B", "class C"}, {"class D", "class E", "class F"},
	}
	for _, classesId := range classes {
		for _, classId := range classesId {
			fmt.Printf("Current class id is \"%v\" \n", classId)
		}
	}
}
