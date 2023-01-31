package stack

import "fmt"

type Stack struct {
	data []interface{}
	size int
}

func (s *Stack) Push(data interface{}) {
	s.data = append(s.data, data)

	s.size++
}

func (s *Stack) Pop() interface{} {

	if s.size == 0 {
		fmt.Println("Stack is empty")
		return nil
	}

	item := s.data[s.size-1]

	s.data = s.data[:s.size-1]

	s.size--

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

	fmt.Printf("\n%s", convertToPostfix("(a*(b-c/(d*e-f)+g))^h"))

}
