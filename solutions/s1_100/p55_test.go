package s1_100

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type case_55 struct {
	input arg_55
	ans   bool
}

func Test_55(t *testing.T) {
	cases := []case_55{
		{arg_55{[]int{2, 3, 1, 1, 4}}, true},
		{arg_55{[]int{3, 2, 1, 0, 4}}, false},
	}

	for _, c := range cases {
		ans := canJump(c.input.nums)
		assert.Equal(t, c.ans, ans)
	}
}
