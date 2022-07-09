package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
	"unsafe"
)

func main() {
	// Boolean (default => false)
	var firstBool bool
	fmt.Println(firstBool)

	// Boolean operands

	// --------------------------------------------------------

	// var (
	// 	falseBool bool = false
	// 	trueBool  bool = true
	// )

	falseBool, trueBool := false, true // мы можем обявлять так потому что как мы сказали что у golang "статическая строгая типизация"
	// --------------------------------------------------------

	fmt.Println("AND", falseBool && trueBool) // false
	fmt.Println("OR", falseBool || trueBool)  // true
	fmt.Println("NOT", !falseBool)            // true

	// ------------------------------------------------------------------------------------------------------------------------------------
	// Numerics. Integers
	// int8, int16, int32, int64, int (разрядность простого int системо-зависимый)
	// uint8, uint16, uint32, uint64, uint

	a := 213
	b := 321

	fmt.Println("Sum of a and b ", a+b)

	// Получаем разрадность int для нашей системы
	fmt.Printf("Type is int in my OS is %T \n", a) // он возвращает int --- то есть мы по этому пути ничего не узнаем
	// Получаем вес в байтах, нашего int

	fmt.Printf("Size of our var in byte %d", unsafe.Sizeof(a)) // и мы увидим что он нам возвращает 8 byte то есть это по меркам bite 64 то есть у него разрядность равен int64

	// ??????
	var int32Var int32 = 12
	var int64Var int64 = 12

	// fmt.Print(int32Var + int64Var) // выдаст ошибку "mismatched types int32 and int64". И это fatal error

	// исправим
	fmt.Println(int64(int32Var) + int64Var) // надо постараться перенсти тип разраядности в большую строну. Если перенести вниз то есть int64 -> int32 то он тупо может не влезить

	// уровнение разрядности косается и int
	// не смотря на то что он у нас равен по разрядности int64
	var intVar int = 12
	// fmt.Printf(int64Var + int32Var) // вот это кинет ошибку
	fmt.Println(int64Var + int64(intVar)) // его тоже все равно надо перенести на нужный разряж

	// ВСЕ ТОЖЕ САМОЕ КАСАЕТСЯ К uintN
	// только то что надо помнить про него это его нало объявлять
	// a := 123 это всегда int и может принять <0 данные

	// ------------------------------------------------------------------------------------------------------------------------------------

	// FLOAT
	// float32, float64

	someFloat := 3.61231231235 // разряд по умолчанию float64
	fmt.Printf("Type of value someFloat is %T\n", someFloat)
	someFloat2 := 3.11231234
	sumOfFloats := someFloat + someFloat2
	fmt.Printf("Sum of floats are :%v \n", sumOfFloats)
	fmt.Printf("Sum of floats are %.3f :", sumOfFloats) // вот так можно подровнять float
	// ------------------------------------------------------------------------------------------------------------------------------------

	// Numeric. Complex

	// ------------------------------------------------------------------------------------------------------------------------------------

	// Strings  --> строка для golang набор БАЙТОВ
	name := "Fedya"
	lastName := "Lavash"
	fullName := name + "-" + lastName
	fmt.Println(fullName)

	// length of string (byte)
	fmt.Println("Length of string ", len(fullName)) // 12 byte

	cyrilicName := "Иван"

	fmt.Println("Get russian string length ", len(cyrilicName)) //  8 byte

	// получение количество рунов (rune) ИЛИ по другому нами принятое длина строки
	fmt.Println("Count rune of string OR count length of string fullName", utf8.RuneCountInString(cyrilicName)) // 4 rune

	// знать есть ли подстрока в строке
	fullText, partOfText := "Lorem ipsum os fishing text to fill places of site", "ipsum"
	fmt.Println(strings.Contains(fullText, partOfText)) // bool(true)

	// rune это alias от int32

	var int32Data int32 = 'ф'
	fmt.Println(int32Data)
	int32Data = 12312312
	fmt.Println(int32Data)

	var runeData rune = 123123	
	fmt.Println(runeData)
	runeData = 'e'
	fmt.Println(runeData)
}
