package main

import "fmt"

func main() {

	//value := 10

	//if value > 10 {
	//	fmt.Println("число больше 10")
	//} else if value < 10 {
	//	fmt.Println("число меньше 10")
	//} else {
	//	fmt.Printf("число равно 10")
	//}

	name := "AnatolyShchurenko"
	raper := true

	if name == "AnatolyShchurenko" {

		if raper == true {
			fmt.Println("Анатолий трушный репер!!!")
		} else {
			fmt.Println("Анатолий не трушный репер :(")
		}
	} else {
		fmt.Println("Это не Анатолий")
	}

	// логические операторы:
	//&& - и
	//|| - или
	//! - не

}
