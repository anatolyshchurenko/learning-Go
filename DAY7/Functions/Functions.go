package main

import "fmt"

func made() {
	fmt.Println("made by Anatoly Shchrenko")
}

func math_func(a int, b int) (int, int, int, int) {

	sum := a + b
	sub := a - b
	mul := a * b
	div := a / b

	return sum, sub, mul, div

}

func main() {

	sum, sub, mul, div := math_func(12, 8)

	fmt.Println(sum, sub, mul, div)
	made()

}
