package s1_100

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type case_11 struct {
	input arg_11
	ans   int
}

func Test_11(t *testing.T) {
	cases := []case_11{
		{arg_11{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}}, 49},
		{arg_11{[]int{1, 1}}, 1},
	}

	for _, c := range cases {
		ans := maxArea(c.input.height)
		assert.Equal(t, c.ans, ans)
	}
}
