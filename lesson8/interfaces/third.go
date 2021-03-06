package main

import "fmt"

//1. А что если создать интерфейс, в котором в принципе нет никаких требований к поведению?
type Empty interface {
}

//2. Вопрос - кто удовлетворяет этому интерфейсу? Если интерфейс пустой - то ему удволетворяет вообще кто угодно.

//3. Реализуем функцию, которая может принимать кого угодно
func Describer(pretendent interface{}) {
	fmt.Printf("Type = %T and value %v\n", pretendent, pretendent)
}

func Data(e Empty) {
	fmt.Println(e)
}

type Student struct {
	name string
}

type SuperInt int

//4. Type Assertion
func SemiGeneric(pretendent interface{}) {
	val, ok := pretendent.(int) // Пытаюсь проверить, что претендент - типа int
	fmt.Println("Value:", val, "Ok?:", ok)
}

func main() {
	strSample := "Hello world!"
	//4. Передача параметра без присваивания в промежуточную переменную
	Describer(strSample)

	intSample := 200
	Describer(intSample)

	Describer(Student{"Vova"})

	fmt.Println("________________")
	varSInt := SuperInt(15)
	fmt.Printf("Type of superint %T \n", varSInt)
	SemiGeneric(varSInt)
	fmt.Println("________________")
	//5. Работа с полу-дженериком
	SemiGeneric(10)
	SemiGeneric(Student{"Fedya"})
	SemiGeneric("Hello world")

	// dataa := "Hello world"
	// Data(dataa)
}
