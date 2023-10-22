package main


import (
	"fmt"
)

func main() {


	//Stack
					
	stack := Stack{}

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	//fmt.Println(stack.Peek()) // 3
	fmt.Println(stack.Pop()) // 3
	fmt.Println(stack.Pop()) // 2
	//fmt.Println(stack.IsEmpty()) // true
	//fmt.Println(stack.Peek()) // 1
	fmt.Println(stack.Pop()) // 1
	fmt.Println(stack.Pop())
	//fmt.Println(stack.IsEmpty()) // true


	//Queue
	queue := Queue{}

	queue.Qenqueue(1)
	queue.Qenqueue("ab")
	queue.Qenqueue(2)
	queue.Qenqueue("cd")
	

	fmt.Println("Размер:", queue.QSize())

	fmt.Println("Извлечение элементов:")
	for !queue.IsEmpty() {
		fmt.Println(queue.Qdequeue())
	}

					//Array
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


	//LsList
	list1 := NewLinkedList()

    list1.InsertBegin("1")
    list1.InsertBegin("2")
    list1.InsertAtEnd("3")
    list1.InsertAtEnd("4")

    fmt.Println("Список:")
    list1.Display()

    list1.DeleteNode("2")
    list1.DeleteNode("4")

    fmt.Println("Список после удаления элементов:")
    list1.Display()

	    // Создание нового двусвязного списка

	list2 := NewDoublyLinkedList()

		// Добавление элементов
	list2.Dladd("1")
	list2.Dladd("2")
	list2.Dladd("3")
	list2.Dladd("4")
	
		// Вывод содержимого
	fmt.Println("Содержимое списка:")
	list2.DLDisplay()
	
		// Удаление элемента
	list2.DlDel(2)
	fmt.Println("Содержимое списка после удаления элемента 2:")
	list2.DLDisplay()
	
}