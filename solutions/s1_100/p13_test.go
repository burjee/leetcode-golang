package s1_100

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type case_13 struct {
	input arg_13
	ans   int
}

func Test_13(t *testing.T) {
	cases := []case_13{
		{arg_13{"III"}, 3},
		{arg_13{"LVIII"}, 58},
		{arg_13{"MCMXCIV"}, 1994},
	}

	for _, c := range cases {
		ans := romanToInt(c.input.s)
		assert.Equal(t, c.ans, ans)
	}
}
