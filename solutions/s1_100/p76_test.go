package s1_100

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type case_76 struct {
	input arg_76
	ans   string
}

func Test_76(t *testing.T) {
	cases := []case_76{
		{arg_76{"ADOBECODEBANC", "ABC"}, "BANC"},
		{arg_76{"a", "a"}, "a"},
		{arg_76{"a", "aa"}, ""},
		{arg_76{"ab", "A"}, ""},
	}

	for _, c := range cases {
		ans := minWindow(c.input.s, c.input.t)
		assert.Equal(t, c.ans, ans)
	}
}
