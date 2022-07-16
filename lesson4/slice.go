package main

import "fmt"

func main() {
	//1. Слайсы (они же - срезы)
	// Слайс - это динамическая обвязка над массивом.
	startArr := [4]int{10, 20, 30, 40}
	var startSlice []int = startArr[0:2] // Слайс инициализируется пустыми квадратными скобками
	fmt.Println("Slice[0:2]:", startSlice)
	// Создали слайс, основываясь уже на существующем массиве

	//2. Создание слайса без явной инициализации массива
	secondSlice := []int{15, 20, 30, 40}
	fmt.Println("SecondSlice:", secondSlice)

	//3. Измнение элементов среза
	originArr := [...]int{30, 40, 50, 60, 70, 80}
	firstSlice := originArr[1:4] // Это набор ссылок на элементы нижележащего массива
	for i, _ := range firstSlice {
		firstSlice[i]++
	}
	fmt.Println("OriginArr:", originArr)   // И он поменяется
	fmt.Println("FirstSlice:", firstSlice) // и он поменяется

	parentArray := [5]int{1, 2, 3, 4, 5}
	var sliceOfParent1 []int = parentArray[:]
	var sliceOfParent2 []int = parentArray[1:3]
	fmt.Printf("First slice of parentArray %v \n", sliceOfParent1)  // [1 2 3 4 5]
	fmt.Printf("Second slice of parentArray %v \n", sliceOfParent2) // [2 3]

	sliceOfParent1[1]++

	fmt.Printf("First slice of parentArray %v \n", sliceOfParent1)  // [1 3 3 4 5]
	fmt.Printf("Second slice of parentArray %v \n", sliceOfParent2) // [3 3] изменение одного слайса который ссылается на родительского массива  меняет значечении и остальных слайсов которые слались на этот массив

	//6. Длина и емкость слайса
	wordsSilce := []string{"one", "two", "three"}
	fmt.Println("slice:", wordsSilce, "Length:", len(wordsSilce), "Capacity:", cap(wordsSilce))
	wordsSilce = append(wordsSilce, "four")
	fmt.Println("slice:", wordsSilce, "Length:", len(wordsSilce), "Capacity:", cap(wordsSilce))
	// Capacity (cap) или ёмкость слайса - это значение, показывающее СКОЛЬКО ЭЛЕМЕНТОВ В ПРИНЦИПЕ
	// можно добавить в срез БЕЗ ВЫДЕЛЕНИЯ ДОПОЛНИТЕЛЬНОЙ ПАМЯТИ ПОД НИЖЕЛЕЖАЩИЙ МАССИВ.

	// Допустим у нас есть срез на 3 элемента (инициализировали без явного указания массива)
	// Компилятор при создании этого среза СНАЧАЛА создал массив ровно на 3 элемента
	// После этого компилятор вернул адрес, где этот масив живет
	// Срез запомнил этот адрес и теперь ссылается на него
	// Потом
	// Начинаем деформировать слайс (увеличим длину /увеличим количество элементов)
	// Проблема - в нижележащем массиве (на котором основан слайс) все 3 ячейки. Что делать?
	// Компилятор ищет в памяти место для массива размера 3*2 (в общем случае n*2, где n - первоначальный размер)
	// После того как место найдено (в нашем случае найдено место для 6 элементов), в это место копируются
	// старые 3 элемента на свои позиции. На 4-ую позицию мы добавляем новый элемент
	// После этого компилятор возвращает нашему слайсу новый адрес в памяти, где находится массив под 6 элементов.

		// 7. Как создавать слайсы наиболее эффективно?
	// make() - это функция, позволяющая более детально создавать срезы
	sl := make([]int, 10, 15)
	// []int - тип коллекции
	// 10 - длина
	// 15 - емкость
	//Сначала инициализируется arr = [15]int
	//Затем по нему делается срез arr[0:10]
	//После чего он возаращается
	fmt.Println(sl)

	geamotricProgressSlice := []int{1}

	for i := 1; i < 1500; i++ {
		fmt.Printf("Current len of slice %v and capacity %v \n", len(geamotricProgressSlice), cap(geamotricProgressSlice))
		geamotricProgressSlice = append(geamotricProgressSlice, i)
	}

	//Важно: после выделения памяти под новый массив, ссылки со старым будут перенсены в новый
	// Пример
	numArr := [2]int{1, 2}
	numSlice := numArr[:]

	numSlice = append(numSlice, 3) // В этот момент numSlice больше не ссылается на numArr
	numSlice[0] = 10
	fmt.Println(numArr)
	fmt.Println(numSlice)



	// добавляем к parentSliceToAddChildren срез childSlice



	// parentSliceToAddChildren := []int{12, 3, 4, 5}
	// childSlice := []int{65, 51, 66}

	// Вариант 1 (not true way)

	// for _, value := range childSlice {
	// 	parentSliceToAddChildren = append(parentSliceToAddChildren, value)
	// }

	// Вариант 2 (true way)

	// parentSliceToAddChildren = append(parentSliceToAddChildren, childSlice...);

	// fmt.Printf("Parent slice %v", parentSliceToAddChildren)


		//9. Многомерный срез
		mSlice := [][]int{
			{1, 2},
			{3, 4, 5, 6},
			{10, 20, 30},
			{},
		}
		fmt.Println(mSlice)

}
