package s1_100

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type case_45 struct {
	input arg_45
	ans   int
}

func Test_45(t *testing.T) {
	cases := []case_45{
		{arg_45{[]int{2, 3, 1, 1, 4}}, 2},
		{arg_45{[]int{2, 3, 0, 1, 4}}, 2},
		{arg_45{[]int{0}}, 0},
	}

	for _, c := range cases {
		ans := jump(c.input.nums)
		assert.Equal(t, c.ans, ans)
	}
}
