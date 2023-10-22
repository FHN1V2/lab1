package main


import (
	"fmt"
)

func main() {
    // Создаем новый MyArray с размером 5
    size:=5
    arr := NewMyArray(size)

    // Вызываем функцию Set для установки значений
    arr.ASet(0, "1")
    arr.ASet(1, "2")
    arr.ASet(2, "3")
    arr.ASet(3, "4")
    arr.ASet(4, "5")

    // Вызываем функцию Get для получения значений
    fmt.Println(arr.AGet(0)) //"1"
    fmt.Println(arr.AGet(1)) //"2"
    fmt.Println(arr.AGet(2)) //"3"
	fmt.Println(arr.AGet(3)) //"4"
	fmt.Println(arr.AGet(4)) //"5"
    // Вызываем функцию Del для удаления значений
    deletedValue := arr.ADel(1)
    fmt.Println(deletedValue)
    fmt.Println("")
    fmt.Println(arr.AGet(1)) // Вывод: ""

    // Создаем экземпляр HashMap
	hmap := HashMap{}

	// Вставляем пару ключ-значение
	err := hmap.Insert("n1", "a")
	if err != nil {
		fmt.Println("Ошибка при вставке:", err)
	}

	err = hmap.Insert("n2", "b")
	if err != nil {
		fmt.Println("Ошибка при вставке:", err)
	}

	err = hmap.Insert("n3", "c")
	if err != nil {
		fmt.Println("Ошибка при вставке:", err)
	}

	// Выводим все элементы через цикл
	for i := 0; i < len(hmap.table); i++ {
		if hmap.table[i] != nil {
			fmt.Printf("Ключ: %s, Значение: %s\n", hmap.table[i].key, hmap.table[i].value)
		}
	}
}
