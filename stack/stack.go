package stack

import (
	"fmt"
)

const MAX_SIZE = 101

// TODO: Change the implementation to allow [] of any type
type Stack struct {
	List [MAX_SIZE]int
	top  int
}

func (s *Stack) Create() *Stack {
	s.top = -1
	return s
}

func (s *Stack) Push(x int) {
	if s.top == MAX_SIZE-1 {
		fmt.Println("Error: stack overflow")
		return
	}

	s.top++
	s.List[s.top] = x
}

func (s *Stack) Pop(x int) {
	if s.top == -1 {
		fmt.Println("Error: No element to pop")
		return
	}

	s.top--
}

func (s *Stack) Top() int {
	return s.List[s.top]
}

func (s *Stack) Print() {
	fmt.Printf("Stack: ")
	for i := 0; i <= s.top; i++ {
		fmt.Printf("%d ", s.List[i])
	}
	fmt.Println()
}
