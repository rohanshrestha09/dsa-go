package stack

import "regexp"

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

func infixToPostfix(infix string) string {

	s := new(Stack)

	var postfix string

	re := regexp.MustCompile("^[a-zA-Z0-9]*$")

	for _, value := range infix {
		switch {
		case value == '(':
			s.Push(value)

		case value == ')':
			var item rune

			for {
				item = s.Pop().(rune)

				if item == '(' {
					break
				}

				postfix = postfix + string(item)
			}

		case re.MatchString(string(value)):
			postfix = postfix + string(value)

		default:
		precedenceOperation:
			TOS := (func() int {
				if s.Length != 0 {
					return checkOPPrecedence(s.TOS.(rune))
				}
				return 0
			})()

			if s.Length == 0 || checkOPPrecedence(value) > TOS {
				s.Push(value)
			} else {

				item := s.Pop().(rune)

				postfix = postfix + string(item)

				goto precedenceOperation
			}
		}
	}

	for s.Length != 0 {
		postfix = postfix + string(s.Pop().(rune))
	}

	return postfix
}
