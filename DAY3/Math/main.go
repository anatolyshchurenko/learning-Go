package main

import (
	"fmt"
	"math"
)

func main() {

	var number float64

	fmt.Println("введите число с которым хотите работать: ")

	fmt.Scan(&number) // Считывает ввод и сохраняет в переменную
	fmt.Println("Вы ввели число:", number)

	//a = math.Sqrt(a) //возведение в корень

	fmt.Println("не округленное число: ", math.Sqrt(number))
	fmt.Println("округленное число: ", math.Round(math.Sqrt(number))) // округление

	fmt.Println("made by Anatoly Shchurenko")

}
