package s1_100

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type case_62 struct {
	input arg_62
	ans   int
}

func Test_62(t *testing.T) {
	cases := []case_62{
		{arg_62{3, 7}, 28},
		{arg_62{3, 2}, 3},
	}

	for _, c := range cases {
		ans := uniquePaths(c.input.m, c.input.n)
		assert.Equal(t, c.ans, ans)
	}
}
