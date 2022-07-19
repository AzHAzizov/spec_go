package main

import "fmt"

type Employee struct {
	name     string
	position string
	salary   int
	curency  string
}

//1. Методы - функции, привязанные к определенным структурам
//func (s Struct) MethodName(parameters type) type {}
//   (s Struct) ----   Reciever - получатель метода

func (e Employee) DisplayInfo() {
	fmt.Println("Name: ", e.name)
	fmt.Println("Position: ", e.position)
	fmt.Println("Salary: ", e.salary)
	fmt.Println("Curency: ", e.curency)
}

func main() {
	var newEployee Employee = Employee{}
	newEployee.DisplayInfo()
}
