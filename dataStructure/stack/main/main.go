package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	MaxTop int
	Top    int
	Arr    [5]int
}

//入栈
func (s *Stack) Push(val int) (err error) {
	if s.Top == s.MaxTop-1 {
		return errors.New("stack full")
	}

	s.Top++
	s.Arr[s.Top] = val

	return
}

//出栈
func (s *Stack) Pop() (val int, err error) {
	if s.Top == -1 {
		err = errors.New("stack full")
		return
	}

	val = s.Arr[s.Top]
	s.Top--
	return val, nil
}

//栈遍历
func (s *Stack) List() {
	if s.Top == -1 {
		fmt.Println("stack full")
		return
	}

	for i := s.Top; i >= 0; i-- {
		fmt.Println(s.Arr[i])
	}
}

//理解栈/堆栈,是一个先入后出的有序列表
func main() {
	s := &Stack{
		MaxTop: 5,
		Top:    -1,
	}

	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(5)

	s.List()

	val, _ := s.Pop()
	fmt.Println("出栈 = ", val)
	val, _ = s.Pop()
	fmt.Println("出栈 = ", val)
	val, _ = s.Pop()
	fmt.Println("出栈 = ", val)
	val, _ = s.Pop()
	fmt.Println("出栈 = ", val)

	s.List()

}
