// This code snippet defines a Stack data structure with methods to push, pop, and peek at the top element. It uses a linked list implementation to store the elements.
package core

type Stack struct {
	Top *Node
}

func NewStack() *Stack {
	return &Stack{Top: nil}
}

func (s *Stack) Push(data string) {
	newNode := &Node{Data: data, Next: s.Top}
	s.Top = newNode
}

func (s *Stack) Pop() string {
	if s.Top == nil {
		return ""
	}

	data := s.Top.Data
	s.Top = s.Top.Next
	return data
}

func (s *Stack) Peek() string {
	if s.Top == nil {
		return ""
	}
	return s.Top.Data
}
