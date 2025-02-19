package s1_100

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type case_17 struct {
	input arg_17
	ans   []string
}

func Test_17(t *testing.T) {
	cases := []case_17{
		{arg_17{"23"}, []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
		{arg_17{""}, []string{}},
		{arg_17{"2"}, []string{"a", "b", "c"}},
	}

	for _, c := range cases {
		ans := letterCombinations(c.input.digits)
		assert.Equal(t, c.ans, ans)
	}
}
