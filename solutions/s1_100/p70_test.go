package s1_100

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type case_70 struct {
	input arg_70
	ans   int
}

func Test_70(t *testing.T) {
	cases := []case_70{
		{arg_70{2}, 2},
		{arg_70{3}, 3},
	}

	for _, c := range cases {
		ans := climbStairs(c.input.n)
		assert.Equal(t, c.ans, ans)
	}
}
