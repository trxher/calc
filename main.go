package main

import (
	"fmt"
)

func main() {
	fmt.Println("Введи операцию: '+' '-' '*' '/'")
	var operation string
	fmt.Scan(&operation)

	fmt.Println("Введи первую переменную:")
	var a float64
	fmt.Scan(&a)

	fmt.Println("Введи вторую переменную:")
	var b float64
	fmt.Scan(&b)

	if operation == "+" {
		fmt.Print("Ответ: ", a+b)
	} else if operation == "-" {
		fmt.Print("Ответ: ", a-b)
	} else if operation == "*" {
		fmt.Print("Ответ: ", a*b)
	} else if operation == "/" {
		if b == 0 {
			fmt.Print("На ноль делить нельзя!")
		} else {
			fmt.Print("Ответ: ", a/b)
		}
	} else {
		fmt.Print("Такой операции не существует!")
	}
}
