package s1_100

import (
	"fmt"
)

func isValid(s string) bool {
	brackets := []rune{}
	for _, ch := range s {
		switch ch {
		case '(':
			brackets = append(brackets, ')')
		case '[':
			brackets = append(brackets, ']')
		case '{':
			brackets = append(brackets, '}')
		default:
			if bracket := pop(&brackets); bracket == nil || *bracket != ch {
				return false
			}
		}
	}
	return len(brackets) == 0
}

func pop(arr *[]rune) *rune {
	l := len(*arr)

	if l == 0 {
		return nil
	}

	e := (*arr)[l-1]
	*arr = (*arr)[:l-1]
	return &e
}

type arg_20 struct {
	s string
}

func Run_20() {
	input := []arg_20{
		{"()"},
		{"()[]{}"},
		{"(]"},
		{"([])"},
		{"("},
	}

	for _, arg := range input {
		fmt.Println(isValid(arg.s))
	}
}
