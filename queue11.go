package main


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
func (q *Queue) QSize() int { //не выводить 
	return len(q.items)
}

// IsEmpty возвращает true, если очередь пуста, и false в противном случае.
func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

