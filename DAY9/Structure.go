package main

import "fmt"

func main() {
	info := anatoly_shchurenko("Анатолий Шуренко", 16)
	fmt.Println(info)
}

type Toljan struct {
	myname string
	age    int
}

func anatoly_shchurenko(name string, age int) Toljan {
	return Toljan{myname: name, age: age}
}
