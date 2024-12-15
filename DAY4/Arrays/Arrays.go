package main

import "fmt"

func main() {
	//var names [3]string
	//names[0] = "Анатолий"
	//names[1] = "Артем"
	//names[2] = "Рина"

	names := [3]string{"Анатолий", "Артем", "Рина"}
	fmt.Println("всего участников: ", len(names))
	fmt.Println(names)
	fmt.Println("РПС в здании детка!")

}
