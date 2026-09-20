package main

func isEven(n int) bool {
	return n%2 == 0
}
func isPositive(n int) bool {
	return n > 0
}

func sum(numbers []int, criteria func(int) bool) int {
	result := 0
	for _, n := range numbers {
		if criteria(n) {
			result += n
		}
	}
	return result
}
func main() {
	slice := []int{1, 2, 3, 4, -7, -4, 23, 52}
	sumOfEvens := sum(slice, isEven)
	println(sumOfEvens)

	sumOfPositive := sum(slice, isPositive)
	println(sumOfPositive)
}
