package s1_100

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type case_33 struct {
	input arg_33
	ans   int
}

func Test_33(t *testing.T) {
	cases := []case_33{
		{arg_33{[]int{4, 5, 6, 7, 0, 1, 2}, 0}, 4},
		{arg_33{[]int{4, 5, 6, 7, 0, 1, 2}, 3}, -1},
		{arg_33{[]int{1}, 3}, -1},
		{arg_33{[]int{4, 5, 6, 7, 8, 1, 2, 3}, 8}, 4},
	}

	for _, c := range cases {
		ans := search(c.input.nums, c.input.target)
		assert.Equal(t, c.ans, ans)
	}
}
