package main

import (
    "fmt"
)

// Структура, представляющая элемент списка
type Node struct {
    Data int     // Значение элемента
    Next *Node  // Указатель на следующий элемент
}

// Структура, представляющая односвязный список
type LinkedList struct {
    Head *Node // Указатель на голову списка
}

// Функция для создания нового элемента списка
func NewNode(data int) *Node {
    return &Node{
        Data: data,
        Next: nil,
    }
}

// Функция для создания нового односвязного списка
func NewLinkedList() *LinkedList {
    return &LinkedList{
        Head: nil,
    }
}

// Функция для добавления элемента в начало списка
func (ll *LinkedList) InsertAtBeginning(data int) {
    newNode := NewNode(data)
    newNode.Next = ll.Head
    ll.Head = newNode
}

// Функция для добавления элемента в конец списка
func (ll *LinkedList) InsertAtEnd(data int) {
    newNode := NewNode(data)
    if ll.Head == nil {
        ll.Head = newNode
        return
    }
    current := ll.Head
    for current.Next != nil {
        current = current.Next
    }
    current.Next = newNode
}

// Функция для удаления элемента из списка по значению
func (ll *LinkedList) DeleteNode(data int) {
    if ll.Head == nil {
        return
    }
    if ll.Head.Data == data {
        ll.Head = ll.Head.Next
        return
    }
    current := ll.Head
    for current.Next != nil && current.Next.Data != data {
        current = current.Next
    }
    if current.Next != nil {
        current.Next = current.Next.Next
    }
}

// Функция для вывода элементов списка
func (ll *LinkedList) Display() {
    current := ll.Head
    for current != nil {
        fmt.Printf("%d -> ", current.Data)
        current = current.Next
    }
    fmt.Println("nil")
}

func main() {
    list := NewLinkedList()

    list.InsertAtBeginning(1)
    list.InsertAtBeginning(2)
    list.InsertAtEnd(3)
    list.InsertAtEnd(4)

    fmt.Println("Список:")
    list.Display()

    list.DeleteNode(2)
    list.DeleteNode(4)

    fmt.Println("Список после удаления элементов:")
    list.Display()
}
