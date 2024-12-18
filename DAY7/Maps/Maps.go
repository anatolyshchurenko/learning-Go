package main

import "fmt"

func main() {

	money := map[string]int{ //создание карты
		"рубли":  1000,
		"франки": 5000,
		"юани":   100,
		"попа":   99999,
	}
	money["юани"] = 2000  //изменение юаней в карте
	delete(money, "попа") //удаление попы(

	fmt.Println(money)
}
