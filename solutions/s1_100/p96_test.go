package s1_100

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type case_96 struct {
	input arg_96
	ans   int
}

func Test_96(t *testing.T) {
	cases := []case_96{
		{arg_96{6}, 132},
		{arg_96{5}, 42},
		{arg_96{4}, 14},
		{arg_96{3}, 5},
		{arg_96{2}, 2},
		{arg_96{1}, 1},
	}

	for _, c := range cases {
		ans := numTrees(c.input.n)
		assert.Equal(t, c.ans, ans)
	}
}
