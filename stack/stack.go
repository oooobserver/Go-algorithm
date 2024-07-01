package stack

import (
	"fmt"
)

type Stack struct {
	items []interface{}
}

// Implement the fmt.print feature
func (s Stack) String() string {
	l := len(s.items)
	res := "---Stack Top---\n"
	for i := l - 1; i >= 0; i-- {
		res += fmt.Sprintf("     %v\n", s.items[i])
	}

	return res
}

func New() Stack {
	return Stack{}
}

func (s *Stack) Push_back(item interface{}) error {
	if _, ok := item.(error); ok {
		return fmt.Errorf("can't add error type")
	}

	s.items = append(s.items, item)
	return nil
}

func (s *Stack) Pop_off() interface{} {
	l := len(s.items)
	if l < 1 {
		return fmt.Errorf("empty stack")
	}
	tmp := s.items[l-1]
	s.items = s.items[:l-1]
	return tmp
}

func (s *Stack) Top() interface{} {
	l := len(s.items)
	if l == 0 {
		return nil
	} else {
		return s.items[l-1]
	}
}

func (s Stack) Len() int {
	return len(s.items)
}
