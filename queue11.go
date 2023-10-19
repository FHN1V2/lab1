package main

import (
	"fmt"
)

// Queue представляет собой структуру данных очередь.
type Queue struct {
	items []interface{}
}

// Enqueue добавляет элемент в конец очереди.
func (q *Queue) Qenqueue(item interface{}) {
	q.items = append(q.items, item)
}

// Dequeue удаляет и возвращает элемент из начала очереди.
func (q *Queue) Qdequeue() interface{} {
	if len(q.items) == 0 {
		return nil
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

// Size возвращает количество элементов в очереди.
func (q *Queue) Size() int { //не выводить 
	return len(q.items)
}

// IsEmpty возвращает true, если очередь пуста, и false в противном случае.
func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

func main() {
	queue := Queue{}

	queue.Qenqueue(1)
	queue.Qenqueue("ab")
	queue.Qenqueue(2)
	queue.Qenqueue("cd")
	

	fmt.Println("Размер:", queue.Size())

	fmt.Println("Извлечение элементов:")
	for !queue.IsEmpty() {
		fmt.Println(queue.Qdequeue())
	}
}
