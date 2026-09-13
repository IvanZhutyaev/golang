package main

import "fmt"

func main() { // Создание функции
	svyat := "svyat"                          // Инициализация
	var andrey string                         // Объявление
	andrey = "andrey"                         // Присвоение
	fmt.Println("Hello World", svyat, andrey) // Вывод

	var str = "hello"
	var intval = 78
	var floatval = 56.34233
	var boolval = true
	fmt.Printf("%v, %v, %e, %t", str, intval, floatval, boolval) //форматированный вывод
	fmt.Println()
	a := 10
	b := 5
	if a > b { //условная конструкция if
		fmt.Println("a>b")
	} else if a < b { //условная конструкция else if
		fmt.Println("a<b")
	} else { //условная конструкция else
		fmt.Println("a=b")
	}
}
