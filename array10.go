package main


import "fmt"

type MyArray struct {
	data   []string
	length int
}

func NewMyArray(size int) MyArray {
	return MyArray{
		data:   make([]string, size),
		length: size,
	}
}

func (arr *MyArray) Aset(index int, value string) {
	if index >= 0 && index < arr.length {
		arr.data[index] = value
	}
}

func (arr *MyArray) Aget(index int) string {
	if index >= 0 && index < arr.length {
		fmt.Println(index)
		return arr.data[index]
	}
	return "Error"
}

func (arr *MyArray) ARadd(value string) {
	arr.data = append(arr.data, value)
	arr.length++
}


func (arr *MyArray) Adel(index int) string {
	if index >= 0 && index < arr.length {
		deletedValue := arr.data[index] // Сохраняем значение элемента
		copy(arr.data[index:], arr.data[index+1:]) // Смещаем элементы влево на одну позицию
		arr.data = arr.data[:len(arr.data)-1]
		arr.length-- // Уменьшаем длину массива
		return deletedValue // Возвращаем удаленное значение
	}
	return "index not found" 
}


func (arr MyArray) PrintArray() {
    for i := 0; i < arr.length; i++ {
        fmt.Printf("%s ", arr.data[i])
    }
    fmt.Println() // Перейти на следующую строку после вывода массива
}
