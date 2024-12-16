package main

import "fmt"

func main() {
	array := [3][2]int{{1, 2}, {3, 4}, {5, 6}}
	fmt.Println(array)       //все массивы
	fmt.Println(array[1][1]) //вывод второй массива, второго элемента
}
