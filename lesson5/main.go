package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	name := "Hello world"
	fmt.Println(name)

	//1. Строка - это байтовый слайс со своими особенносятми при обращении
	//к нижелажащему массиву
	word := "Simple word"
	fmt.Printf("We are print string %s\n", word)

	// Какие значения байт сейчас находятся в слайсе word?
	fmt.Printf("Bytes: ")
	for i := 0; i < len(word); i++ {
		fmt.Printf("%x ", word[i]) //%x - формат представления 16-ти ричного байта
	}
	fmt.Println()
	// Каким образом получать доступ к отдельно стоящим символам?
	fmt.Printf("Characters: ")
	for i := 0; i < len(word); i++ {
		fmt.Printf("%c ", word[i]) //%c - формат представления символа
	}
	fmt.Println()

	//2. Строки в Go - хранятся как наборы UTF-8символов. Каждый символ, вообще говоря, может занимать
	// больше чем 1 байт

	//3. Руна (Rune) - это стандартный встроенный тип в Go (alias над int32), позволяющий хранить
	//единый неделимый UTF символ ВНЕ ЗАВИСИМОСТИ ОТ ТОГО сколько байт он занимает

	fmt.Printf("Runes: ")
	runeSlice := []rune(word) // Преобразование слайса байт к слайсу рун []byte(sliceRune)
	// for i := 0; i < len(runeSlice); i++ {
	// 	fmt.Printf("%c ", runeSlice[i])
	// }

	for _, value := range runeSlice {
		fmt.Printf("%c", value)
	}

	println()

	//4. Итерирование по строке с использованием рун
	cyrilWord := "привет мир"
	for id, runeVal := range cyrilWord { // id - это индекс байта, с КОТОРОГО НАЧИНАЕТСЯ РУНА runeVal
		fmt.Printf("%c starts at postion %d\n", runeVal, id)
	}

	//5. Создание строки из слайса байт
	myByteSlice := []byte{0x40, 0x41, 0x42, 0x43} // Исходное представление байтов
	myStr := string(myByteSlice)
	fmt.Println(myStr)

	myDecimalByteSlice := []byte{100, 101, 102, 103} // Синтаксический сахар - можно использовать десятичное представление байтов
	myDecimalStr := string(myDecimalByteSlice)
	fmt.Println(myDecimalStr)

	// ----------------------------------------------------------------------------------------------
	// OWN PRACTICE --> строка => руна => байт  => строка

	ownString := "My best string ИЛИ"
	ownStringAsRune := []rune(ownString)
	fmt.Println(ownStringAsRune)
	fmt.Println(len(ownStringAsRune))

	fmt.Println("Rune as string : ", string(ownStringAsRune))

	var stringToByte []byte
	for _, data := range ownStringAsRune {
		stringToByte = append(stringToByte, byte(data))
	}

	ownStringByteToString := string(stringToByte)
	fmt.Println(ownStringByteToString)
	// ----------------------------------------------------------------------------------------------

	//7. Длина и емкость строки

	// Длина len() - количество байт в слайсе
	fmt.Println("Length of Вася:", len("Вася"), "bytes")
	// Длина RuneCounter - количество !рун!
	fmt.Println("Length of Вася:", utf8.RuneCountInString("Вася"), "runes") // Это тоже самое что : len([]rune("Вася"))

	// Вычисление емкости строки - бессмысленно, т.к. строка базовый тип
	fmt.Println(cap([]rune("Вася")))

	//8. Сравнение строк == и !=. Начиная с go 1.6
	word1, word2 := "Вася", "Петя"
	if word1 == word2 {
		fmt.Println("Equal")
	} else {
		fmt.Println("Not equal")
	}


	//9. Конкатенация
	word3 := word1 + word2
	fmt.Println(word3)

	//10. Строитель строк (java -> StringBuiler)
	firstName := "Alex"
	secondName := "Johnson"
	fullName := fmt.Sprintf("%s #### %s", firstName, secondName)
	fmt.Println("FullName:", fullName)


	//11. Строки не изменяемые
	// fullName[0] = "Q"

}
