package main

import (
	"fmt"
	"strings"
)

func main() {

	// var someInt int64 = 8231231231231123123
	// fmt.Printf("Type of var is %T \n", someInt)

	// fmt.Printf("Type of var is %T and size %d bytes \n", someInt, unsafe.Sizeof(someInt))

	// var name string = "Имя"
	// fmt.Printf("Real length of string is %d \n", utf8.RuneCountInString(name))

	// longString, needStringPart := "Hello my name is NAME", "name"
	// fmt.Println(strings.Contains(longString, needStringPart))

	// ----------------------------------------------------------------------

	// var oldPassword = "123"
	// var newPassword string
	// fmt.Scan(&newPassword)

	// if strings.Compare(newPassword, oldPassword) == 0 {
	// 	fmt.Println("New pass is same with old")
	// } else {
	// 	fmt.Println("Password was change")
	// }

	// ----------------------------------------------------------------------

	fmt.Print(1)
	fmt.Print(2)
	fmt.Print(4)
	// os.Exit(2)
	panic(2)
	fmt.Print(5)
	fmt.Print(6)

	if name := getName(); strings.Compare("admin", name) == 0 {
		fmt.Println("get name of admin")
	} else {
		fmt.Println("got simple name")
	}

}

func getName() string {
	return "admin"
}
