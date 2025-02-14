package s1_100

import (
	"fmt"
	"strings"
)

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	str_rows := make([]string, numRows)
	row := 0
	offset := 1
	for _, c := range s {
		str_rows[row] += string(c)
		row += offset
		if row == 0 {
			offset = 1
		}
		if row == numRows-1 {
			offset = -1
		}

	}
	return strings.Join(str_rows, "")
}

type arg_6 struct {
	s       string
	numRows int
}

func Run_6() {
	input := []arg_6{
		{"PAYPALISHIRING", 3},
		{"PAYPALISHIRING", 4},
		{"A", 1},
	}

	for _, arg := range input {
		fmt.Println(convert(arg.s, arg.numRows))
	}
}
