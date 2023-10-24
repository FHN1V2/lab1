package main

import (
	"fmt"
)

func main() {

/*
	//Stack
					
	stack := Stack{}

	stack.Spush(1)
	stack.Spush(2)
	stack.Spush(3)
	//fmt.Println(stack.Peek()) // 3
	fmt.Println(stack.Spop()) // 3
	fmt.Println(stack.Spop()) // 2
	//fmt.Println(stack.IsEmpty()) // true
	//fmt.Println(stack.Peek()) // 1
	fmt.Println(stack.Spop()) // 1
	fmt.Println(stack.Spop())
	//fmt.Println(stack.IsEmpty()) // true


	//Queue
	queue := Queue{}

	queue.Qadd("1")
	queue.Qadd("ab")
	queue.Qadd("2")
	queue.Qadd("cd")
	

	fmt.Println("Размер:", queue.Qsize())

	fmt.Println("Извлечение элементов:")
	for !queue.IsEmpty() {
		fmt.Println(queue.Qdell())
	}

					//Array
    // Создаем новый MyArray с размером 5
    size:=5
    arr := NewMyArray(size)

    // Вызываем функцию Set для установки значений
    arr.Aset(0, "1")
    arr.Aset(1, "2")
    arr.Aset(2, "3")
    arr.Aset(3, "4")
    arr.Aset(4, "5")

    // Вызываем функцию Get для получения значений
    fmt.Println(arr.Aget(0)) //"1"
    fmt.Println(arr.Aget(1)) //"2"
    fmt.Println(arr.Aget(2)) //"3"
	fmt.Println(arr.Aget(3)) //"4"
	fmt.Println(arr.Aget(4)) //"5"
    // Вызываем функцию Del для удаления значений
    deletedValue := arr.Adel(1)
    fmt.Println(deletedValue)
    fmt.Println("")
    fmt.Println(arr.Aget(1)) // Вывод: ""

    // Создаем экземпляр HashMap
	hmap := HashMap{}

	// Вставляем пару ключ-значение
	err := hmap.Hadd("n1", "a")
	if err != nil {
		fmt.Println("Ошибка при вставке:", err)
	}

	err = hmap.Hadd("n2", "b")
	if err != nil {
		fmt.Println("Ошибка при вставке:", err)
	}

	err = hmap.Hadd("n3", "c")
	if err != nil {
		fmt.Println("Ошибка при вставке:", err)
	}

	// Выводим все элементы через цикл
	for i := 0; i < len(hmap.table); i++ {
		if hmap.table[i] != nil {
			fmt.Printf("Ключ: %s, Значение: %s\n", hmap.table[i].key, hmap.table[i].value)
		}
	}


	//LsList
	list1 := NewLinkedList()

    list1.SLaddS("1")
    list1.SLaddS("2")
    list1.SLaddE("3")
    list1.SLaddE("4")

    fmt.Println("Список:")
    list1.SLdisplay()

    list1.SLdel("2")
    list1.SLdel("4")

    fmt.Println("Список после удаления элементов:")
    list1.SLdisplay()

	    // Создание нового двусвязного списка

	list2 := NewDoublyLinkedList()

		// Добавление элементов
	list2.Dladd("1")
	list2.Dladd("2")
	list2.Dladd("3")
	list2.Dladd("4")
	
		// Вывод содержимого
	fmt.Println("Содержимое списка:")
	list2.DLdisplay()
	
		// Удаление элемента
	list2.Dldel(2)
	fmt.Println("Содержимое списка после удаления элемента 2:")
	list2.DLdisplay()
*/
	stack:=Stack{}
	queue:=Queue{}
	array:=MyArray{}
	var Isinput bool
	Isinput =true
	for Isinput{
		var input string
		var index int
		
		
		fmt.Scanf("%s", &input)
		switch input{
		case "SPUSH":
			fmt.Scanf("%s",&input)
			stack.Spush(input)
		case "SPOP":
			fmt.Println(stack.Spop())
		case "QPUSH":
			fmt.Scanf("%s %s",&input)
			queue.Qadd(input)
		case "QPOP":
			fmt.Println(queue.Qdell())
		case "APUSH":
			fmt.Scanf("%s",&input)
			array.ARadd(input)
		
		case "AGET":
			fmt.Scanf("%d",index)
		fmt.Println(array.Aget(index))
	}
}
}