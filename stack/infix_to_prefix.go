package stack

import (
	"regexp"
)

type Stack []rune

func (s *Stack) Push(data rune) {
	*s = append(*s, data)
}

func (s *Stack) Pop() rune {
	item := (*s)[len(*s)-1]

	*s = (*s)[:len(*s)-1]

	return item
}

func (s Stack) TOS() rune {
	return s[len(s)-1]
}

func checkOPPrecedence(operator rune) int {
	switch operator {
	case '^', '$':
		return 3

	case '*', '/':
		return 2

	case '+', '-':
		return 1
	}

	return 0
}

func InfixToPostfix(infix string) string {
	var s Stack

	var postfix string

	re := regexp.MustCompile("^[a-zA-Z0-9]*$")

	for _, value := range infix {
		switch {
		case value == '(':
			s.Push(value)

		case value == ')':

			for {
				item := s.Pop()

				if item == '(' {
					break
				}

				postfix = postfix + string(item)
			}

		case re.MatchString(string(value)):
			postfix = postfix + string(value)

		default:
		precedenceOperation:
			TOS := 0

			if len(s) != 0 {
				TOS = checkOPPrecedence(s.TOS())
			}

			if len(s) == 0 || checkOPPrecedence(value) > TOS {
				s.Push(value)
			} else {

				item := s.Pop()

				postfix = postfix + string(item)

				goto precedenceOperation
			}
		}
	}

	for len(s) != 0 {
		postfix = postfix + string(s.Pop())
	}

	return postfix
}
