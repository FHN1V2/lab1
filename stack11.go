package main

import (
	"errors"
)



type SNode struct {
	next *SNode
	val int
}

type Stack struct {
	head *SNode
}

func (s *Stack) Spush(val int) {
	newNode := &SNode{val: val}
	newNode.next = s.head
	s.head = newNode
}

func (s *Stack) Spop() (int, error) {
	if s.head == nil {
		return 0, errors.New("stack is empty")
	}
	val := s.head.val 
	s.head = s.head.next
	return val, errors.New("")
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

