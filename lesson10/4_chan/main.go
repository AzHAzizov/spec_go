package main

import "fmt"

//1. Каналы - средство для коммуникации между горутинами.
// Каналы можно рассматривать как соединетильные трубы, через которые горутины между собой общаются (аналогично тому,
// как вода течет по трубам, данные передаются через каналы)

//2. Объявление канала.
// Каналы по умолчанию имеют zeroValue - nil. Поэтому их создают через фукнцию make.

func main() {
	var a chan int // объявляем канал, через который будут передаваться данные типа int
	// fmt.Printf("Type of a is %T and value %v \n", a, a)
	if a == nil {
		fmt.Println("channel is nil, Let's define it")
		a = make(chan int)
		fmt.Printf("Type of a is %T and value %v \n", a, a)
	}
}