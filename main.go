package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

var operation string
var a, b float64

// Заставка
func miniature() {
	content, err := ioutil.ReadFile("miniature.txt")
	// the file is inside the local directory
	if err != nil {
		fmt.Println("Err")
	}

	fmt.Println(string(content))
}

// Ввод знака операции
func inputOperationUser() {
	fmt.Println("Введи операцию:\n'+' - сложение\n'-' - вычитание\n'*' - умножение\n'/' - деление")
	fmt.Scan(&operation)
}

// Ввод переменных
func inputVariablesUser() {
	fmt.Println("Введи первую переменную:")
	fmt.Scan(&a)

	fmt.Println("Введи вторую переменную:")
	fmt.Scan(&b)
}

// Математические вычесления
func calc() {
	if operation == "+" {
		fmt.Println("Ответ:", a+b)
	} else if operation == "-" {
		fmt.Println("Ответ:", a-b)
	} else if operation == "*" {
		fmt.Println("Ответ:", a*b)
	} else if operation == "/" {
		if b == 0 {
			fmt.Println("На ноль делить нельзя!")
		} else {
			fmt.Println("Ответ:", a/b)
		}
	} else {
		fmt.Print("Неверная операция!")
	}
}

// Обработчик выхода из программы
func exit() {
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

// Основная функция
func main() {
	miniature()
	inputOperationUser()
	inputVariablesUser()
	calc()

	exit()
}
