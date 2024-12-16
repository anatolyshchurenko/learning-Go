package main

import (
	"fmt"
	"sort"
)

func main() {
	slice := []int{5, 3, 4, 2, 1} //создание слайса

	slice = append(slice, 6) //добвить в конец слайса новый элемент (в данном случае цыфра 6)
	sort.Ints(slice)         //сортировка слайса в порядке возрастания
	fmt.Println(slice)       //вывод слайса

	slicestr := []string{"g", "r", "e", "t", "d", "a"}  //cоздание слайса типа string
	fmt.Println(slicestr, ": неотсортированная версия") //выводит неотсортированный слайс
	sort.Strings(slicestr)                              //сортировка слайса по алфавиту
	fmt.Println(slicestr, ": отсортированная версия")   //выводит отсортированный слайс

	//el == slice[i]
	for _, el := range slice { // перебор элементов в массиве
		fmt.Printf("%d\n", el) //вывод номера массива и его элемента
	}

	//made by Anatoly Shchurenko

}
