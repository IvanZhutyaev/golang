# GOLANG
**Если переменная не используется, то go будет ругаться**
### Операции
```a:="something"``` **- операция присвоения**<br>
```var b string"``` **- операция объявление**<br>
```b="somethin1"``` **- операция присвоения**<br>

### Вывод
```fmr.Println(b)``` **вывод**
```fmr.Printf(b)``` **форматированный вывод**


### Значение по умолчанию
**Если переменной не присвоено значение, то она имеет значение по умолчанию, которое оперделено для её типа:<br>Для числовых типов: число 0<br>Для логического типа: false<br>Для строк: ""(пустая строка, а не null)**

### Константы
```const pi float64 = 3.1415```<br><br>
**или**<br>
```go
const(
	pi float64 = 3.1415
	e float64 = 2.7182
)
```

### iota-идентификатор
**Go поддерживает идентификатор iota, который увеличивает с объявлением каждой новой константы Всякий раз, когда компилятор получает блок const, идентификатор iota сбрасывается до нуля и увеличивается на единицу после каждой непустой строки в блоке const**

### Условные конструкции
```go
a := 10
b := 5
if a > b { //условная конструкция if
    fmt.Println("a>b")
} else if a < b { //условная конструкция else if
    fmt.Println("a<b")
} else { //условная конструкция else
    fmt.Println("a=b")
}
```
```go
a:=5
// конструкция switch case
switch(a){ 
    case 9: fmt.Println("a=9")
    case 8: fmt.Println("a=8"); fallthrough
    case 7: fmt.Println("a=7")
    case 6,5,4: 
		fmt.Println("a=6 or a=5 or a=4, but it's maybe not")
	default:	
    fmt.Println("value of variable is not defined")
}
```
**fallthrough используется для перехода к следующему случаю**

### Циклы 
<h4>В GO есть только цикл for!</h4>

```go
for i:=1; i<10; i++{ // вариант обычного for
	fmt.Println(i*i)
}
```
```go
var i = 1
for i<10{ // вариант while
	fmt.Println(i*i)
	i++
}
```
```go
str:="hello" 
for index, value :=range str{ // вариант for each
	fmt.Printf("Index: %d, Value: %c\n", index, value)
}
```
```go
str:="hello" 
for _, value :=range str{ // вариант for each, без индекса
	fmt.Printf(" %c: ",value)
}
fmt.Printf(" %c: ",10)
```

### Массивы
```var numbers [5]int```**- объявление массива**<br>
```var numbers [5] int = [5]int{1,2,3,4,5}```**- инициализация массива**

**len() - длина массива**

## Функции и их параметры
```go
func hello(){} //создание функции
func increment(x int){} //функция с парметром
func add(numbers ...int){} //функция с несколькими параметрами
```
### Пример возвращаемой функции
```go
//1
func add(a, b int) (z int) {
	z = a + b
	return
}

//2
func add1(a, b int) int {
z = a + b
return z
}
```
**Функция может возвращать несколько значений**