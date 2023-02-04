package stack

func checkOPPrecedence(operator rune) int {
	switch operator {
	case '^':
		fallthrough
	case '$':
		return 3

	case '*':
		fallthrough
	case '/':
		return 2

	case '+':
		fallthrough
	case '-':
		return 1
	}

	return 0
}

func infixToPostfix(infix string) string {

	s := new(Stack)

	var postfix string

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

		case value >= 65 && value <= 90 || value >= 97 && value <= 122:
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
