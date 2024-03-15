package main

import (
	"bufio"
	"fmt"
	"os"
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
		fmt.Println("Ответ: ", a+b)
	} else if operation == "-" {
		fmt.Println("Ответ: ", a-b)
	} else if operation == "*" {
		fmt.Println("Ответ: ", a*b)
	} else if operation == "/" {
		if b == 0 {
			fmt.Println("На ноль делить нельзя!")
		} else {
			fmt.Println("Ответ: ", a/b)
		}
	} else {
		fmt.Println("Такой операции не существует!")
	}

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		exit := scanner.Text()
		if exit == "q" {
			break
		} else {
			fmt.Println("Нажмите 'q' для выхода")
		}
	}
}
