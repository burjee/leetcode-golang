package s1_100

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type case_54 struct {
	input arg_54
	ans   []int
}

func Test_54(t *testing.T) {
	cases := []case_54{
		{arg_54{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}}, []int{1, 2, 3, 6, 9, 8, 7, 4, 5}},
		{arg_54{[][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}}, []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}},
	}

	for _, c := range cases {
		ans := spiralOrder(c.input.matrix)
		assert.Equal(t, c.ans, ans)
	}
}
