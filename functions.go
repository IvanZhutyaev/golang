package main

import "fmt"

//	func isEven(n int) bool {
//		return n%2 == 0
//	}
//
//	func isPositive(n int) bool {
//		return n > 0
//	}
//
//	func sum(numbers []int, criteria func(int) bool) int {
//		result := 0
//		for _, n := range numbers {
//			if criteria(n) {
//				result += n
//			}
//		}
//		return result
//	}
func main() {
	//slice := []int{1, 2, 3, 4, -7, -4, 23, 52}
	//sumOfEvens := sum(slice, isEven)
	//println(sumOfEvens)
	//
	//sumOfPositive := sum(slice, isPositive)
	//println(sumOfPositive)

	//f := selectFn(1)
	//fmt.Println(f(3, 4))
	//f1 := selectFn(3)
	//fmt.Println(f1(3, 4))

	d := 5
	var p *int
	p = &d
	fmt.Println("Before:", d)
	changeValue(p)
	fmt.Println("After:", d)

}
func selectFn(n int) func(int, int) int {
	if n == 1 {
		return add
	} else if n == 2 {
		return subtract
	} else {
		return multiply
	}
}

func multiply(i int, i2 int) int { return i * i2 }

func subtract(i int, i2 int) int { return i - i2 }

func add(i int, i2 int) int { return i + i2 }

func changeValue(x *int) {
	*x = (*x) * (*x)
}
