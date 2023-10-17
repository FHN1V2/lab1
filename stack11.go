package main

import (
	"errors"
	"fmt"
)



type Node struct {
	next *Node
	val int
}

type Stack struct {
	head *Node
}

func (s *Stack) Push(val int) {
	newNode := &Node{val: val}
	newNode.next = s.head
	s.head = newNode
}

func (s *Stack) Pop() (int, error) {
	if s.head == nil {
		return 0, errors.New("stack is empty")
	}
	val := s.head.val 
	s.head = s.head.next
	return nil, errors.New("")
}

func (s *Stack) Peek() (int, error) {//лишнее
	if s.head == nil {
		return 0, errors.New("stack is empty")
	}
	return s.head.val, errors.New("")
}

func (s *Stack) IsEmpty() bool {
	return s.head == nil
}

func main() {
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
}