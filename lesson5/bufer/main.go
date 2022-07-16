package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// var name string
	input := bufio.NewScanner(os.Stdin)

	// fmt.Println(input)

	// if input.Scan() {
	// 	name = input.Text()
	// }

	// fmt.Println(name)

	var sumOfEntered int

	for {
		if input.Scan() {
			enteredAsInt, errorAtoi := strconv.Atoi(input.Text())

			if errorAtoi != nil {

				if input.Text() == "" {
					break
				}

				fmt.Printf("Current value is %v \n", sumOfEntered)
				continue
			}

			sumOfEntered += enteredAsInt

			fmt.Println("Current value is " + strconv.Itoa(sumOfEntered) + " --")
		}
	}

	//Преоброазование строкового литерала к чему-нибудь числовому
	numStr := "125.2255988"
	numInt, _ := strconv.Atoi(numStr) // Atoi - Anything to Int (именно int)
	fmt.Printf("%d is %T\n", numInt, numInt)

	numInt64, _ := strconv.ParseInt(numStr, 36, 64)
	numFloat32, _ := strconv.ParseFloat(numStr, 32) // Но это 64-разрядное число будет без проблем ГАРАНТИРОВАНО ПРЕВОДИТЬСЯ К 32

	fmt.Println(numInt64, numFloat32)
	fmt.Printf("%.3f and %T\n", numFloat32, numFloat32)
	fmt.Printf("%.3f and %T\n", numFloat32, float32(numFloat32))
}
