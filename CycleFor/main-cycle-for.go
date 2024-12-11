package main

import "fmt"

func main() {
	//for i := 0; i < 10; i++ {
	//	fmt.Println(i)
	//	time.Sleep(1 * time.Second)
	//}

	for i := 0; i < 13; i++ {
		if i == 10 {
			fmt.Println("число 10")
			//continue
			break

		}
		fmt.Println(i)
	}
	fmt.Println("СТОП")
}
