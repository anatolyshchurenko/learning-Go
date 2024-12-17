package main

import "fmt"

func main() {

	age := 16
	name := "Анатолий Щуренко"
	poshalko := 148.8
	rapper := true

	fmt.Printf("мне сейчас %d \n", age)                   //вывод возраст и переход на новую строку
	fmt.Printf("меня зовут %s \n", name)                  //вывод имени и переход на новую строку
	fmt.Printf("ПОСХАЛКО!!! -----> %f \n", poshalko)      //<----- ПОСХАЛКО и переход на новую строку
	fmt.Printf("Анатолий Щуренко реально %t )))", rapper) //вывод буллевой переменной

	// \n -- переход на новую строку
	// %d -- для int
	// %s -- для string
	// %f -- для float
	// %t -- для bool
}
