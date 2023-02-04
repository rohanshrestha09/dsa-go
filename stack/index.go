package stack

import "fmt"

type Stack struct {
	data   []interface{}
	Length int
	TOS    interface{}
}

func (s *Stack) Push(data interface{}) {
	s.data = append(s.data, data)

	s.TOS = data

	s.Length++
}

func (s *Stack) Pop() interface{} {

	if s.Length == 0 {
		fmt.Println("Stack is empty")
		return nil
	}

	item := s.data[s.Length-1]

	s.data = s.data[:s.Length-1]

	s.Length--

	if s.Length != 0 {
		s.TOS = s.data[len(s.data)-1]
	}

	return item
}

func (s *Stack) Display() {
	for _, val := range s.data {
		fmt.Printf("%v\t", val)
	}
}

func Run() {

	// s := new(Stack)

	// s.Push(10)

	// s.Push(200)

	// s.Push("Nice")

	// s.Pop()

	// s.Display()

	fmt.Printf("\n%s", infixToPostfix("A+(B*C-(D/E$F)*G)*H"))

}
