package main

import "fmt"

func main() { // Создание функции
	//svyat := "svyat"                          // Инициализация
	//var andrey string                         // Объявление
	//andrey = "andrey"                         // Присвоение
	//fmt.Println("Hello World", svyat, andrey) // Вывод
	//
	//var str = "hello"
	//var intval = 78
	//var floatval = 56.34233
	//var boolval = true
	//fmt.Printf("%v, %v, %e, %t", str, intval, floatval, boolval) //форматированный вывод
	//fmt.Println()
	//a := 10
	//b := 5
	//if a > b { //условная конструкция if
	//	fmt.Println("a>b")
	//} else if a < b { //условная конструкция else if
	//	fmt.Println("a<b")
	//} else { //условная конструкция else
	//	fmt.Println("a=b")
	//}
	//
	//r := 8
	//// конструкция switch case
	//switch r {
	//case 9:
	//	fmt.Println("r=9")
	//case 8:
	//	fmt.Println("r=8")
	//	fallthrough
	//case 7:
	//	fmt.Println("r=7")
	//case 6, 5, 4:
	//	fmt.Println("r=6 or r=5 or r=4, but it's maybe not")
	//default:
	//	fmt.Println("value of variable is not defined")
	//}
	//
	//fmt.Println("FOR:")
	//for i := 1; i < 10; i++ { // вариант обычного for
	//	fmt.Println(i * i)
	//}
	//
	//fmt.Println("WHILE:")
	//var i = 1
	//for i < 10 { // вариант while
	//	fmt.Println(i * i)
	//	i++
	//}
	//
	//fmt.Println("FOREACH:")
	//hello := "hello"
	//for index, value := range hello { // вариант for each
	//	fmt.Printf("Index: %d, Value: %c\n", index, value)
	//}
	//
	//fmt.Println("FOREACH, but:")
	//hello1 := "hello"
	//for _, value := range hello1 { // вариант for each, но без индекса
	//	fmt.Printf(" %c ", value)
	//}
	//fmt.Printf(" %c ", 10)

	var a = add(1, 2)
	var b = add(2, 2)
	fmt.Println(a)
	fmt.Println(b)

	var a1 = add1(1, 2)
	var b1 = add1(2, 2)
	fmt.Println(a1)
	fmt.Println(b1)
}

func add(a, b int) (z int) {
	z = a + b
	return
}
func add1(a, b int) int {
	return a + b
}
