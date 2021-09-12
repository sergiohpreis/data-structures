package stack

import (
	"fmt"
)

type Stack []interface{}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Pop() (interface{}, bool) {
	if s.IsEmpty() {
		return nil, false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}

func (s *Stack) Push(x interface{}) {
	*s = append(*s, x)
}

func (s *Stack) Top() interface{} {
	return (*s)[len(*s)-1]
}

func (s *Stack) Print() {
	for _, element := range *s {
		fmt.Printf("%d ", element)
	}
	fmt.Println()
}
