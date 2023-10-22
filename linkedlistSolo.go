package main

import (
    "fmt"
)

// Структура, представляющая элемент списка
type LSNode struct {
    Data string     // Значение элемента
    Next *LSNode  // Указатель на следующий элемент
}

// Структура, представляющая односвязный список
type LinkedList struct {
    Head *LSNode // Указатель на голову списка
}

// Функция для создания нового элемента списка
func NewLSNode(data string) *LSNode {
    return &LSNode{
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
func (ll *LinkedList) InsertBegin(data string) {
    NewLSNode := NewLSNode(data)
    NewLSNode.Next = ll.Head
    ll.Head = NewLSNode
}

// Функция для добавления элемента в конец списка
func (ll *LinkedList) InsertAtEnd(data string) {
    NewLSNode := NewLSNode(data)
    if ll.Head == nil {
        ll.Head = NewLSNode
        return
    }
    current := ll.Head
    for current.Next != nil {
        current = current.Next
    }
    current.Next = NewLSNode
}

// Функция для удаления элемента из списка по значению
func (ll *LinkedList) DeleteNode(data string) {
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
        fmt.Printf("%s -> ", current.Data)
        current = current.Next
    }
    fmt.Println("nil")
}


