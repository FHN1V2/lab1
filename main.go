package main


import (
	"fmt"
)

func main() {
    // Создаем новый MyArray с размером 5
    arr := NewMyArray(5)

    // Вызываем функцию Set для установки значений
    arr.ASet(0, "1")
    arr.ASet(1, "2")
    arr.ASet(2, "3")
    arr.ASet(3, "4")
    arr.ASet(4, "5")

    // Вызываем функцию Get для получения значений
    fmt.Println(arr.Get(0)) //"1"
    fmt.Println(arr.Get(1)) //"2"
    fmt.Println(arr.Get(2)) //"3"
	fmt.Println(arr.Get(3)) //"4"
	fmt.Println(arr.Get(4)) //"5"
    // Вызываем функцию Del для удаления значений
    deletedValue := arr.Del(1)
    fmt.Println(deletedValue)
    fmt.Println(arr.Get(1)) // Вывод: ""
}
