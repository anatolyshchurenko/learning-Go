package main

import "fmt"

func made() {
	fmt.Println("made by Anatoly Shchrenko")
}

func math_func(a float32, b float32) (float32, float32, float32, float32) {

	sum := a + b //сумма
	sub := a - b //разность
	mul := a * b //умножение
	div := a / b //деление

	return sum, sub, mul, div //возвращаю

}

func main() {

	sum, sub, mul, div := math_func(12, 8)

	fmt.Println(sum, sub, mul, div)
	made()

}
