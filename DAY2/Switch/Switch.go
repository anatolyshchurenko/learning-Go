package main

import "fmt"

func main() {
	//name := "Пиппука"

	//switch name {
	//case "Пиппука":
	//	fmt.Println("Привет Пиппука!")
	//case "Анатолий":
	//	fmt.Println("Привет Создатель!")
	//case "Артем":
	//	fmt.Println("Привет Артемыч!")
	//	case "РПС":
	//	fmt.Println("РПС в здании детка!")
	//default:
	//	fmt.Println("тебе не стоило этого делать")

	//	for i := 1; i <= 3; i++ {
	//		fmt.Printf(".")
	//		time.Sleep(1 * time.Second)
	//	}

	//} ---------- Это всё лишь тренировка, пока что без ввода(

	number := 10

	switch {
	case number > 1:
		fmt.Println("Это число больше чем 1")
		fallthrough //команда которая заставляет код не останавливаться, а продолжать проверять остальные условия
	case number < 11:
		fmt.Println("Это число меньше чем 11")
	}
}
