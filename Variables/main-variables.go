package main

import "fmt"

func main() {
	//var age int  ---------первый способ
	//age = 16
	//var age = 16 ---------второй способ
	age := 16                   // int - числа
	name := "Anatoly Shchrenko" //string - строка
	raper := true

	fmt.Println(name, age, raper)

	// еще способы создания переменных
	//var test float32 -- создание переменной типа float
	//var test uint -- создание переменной типа uint (от нуля до бесконечность, ТОЛЬКО ПОЛОЖИТЕЛЬНЫЕ ЧИСЛА)
}
