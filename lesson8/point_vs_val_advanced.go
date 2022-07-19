package main

import "fmt"

type Rectangle struct {
	length, width int
}

func (r *Rectangle) Area() int {
	return r.length * r.width
}

func Area(r *Rectangle) int {
	return r.length * r.width
}

func (r *Rectangle) SetWith(newWidth int) {
	r.width = newWidth
}

func main() {
	//Значение исходное
	rectangleAsValue := Rectangle{10, 15}
	//Ссылка на исходное значение
	rectangleAsPointer := &rectangleAsValue

	println("Call area of function &rectangle:", Area(rectangleAsPointer))
	println("Call area of method &rectangle:", rectangleAsValue.Area())
	println("Call area of method &rectangle:", rectangleAsPointer.Area())

	//1. Вызываем метод у value - исходного значения
	fmt.Println("Call Area method for rectangle:", rectangleAsValue.Area())
	//2. Вызываем функцию с параметром value - исходное значение
	//Area(rectangleAsValue)

	rectangleAsValue.SetWith(45)
	fmt.Println("get area of rectangleValue after set new With", rectangleAsValue.Area())

	rectangleAsPointer.SetWith(65)
	fmt.Println("get area of rectanglePointer after set new With", rectangleAsPointer.Area())

}
