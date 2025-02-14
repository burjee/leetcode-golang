package s1_100

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type case_5 struct {
	input arg_5
	ans   string
}

func Test_5(t *testing.T) {
	cases := []case_5{
		{arg_5{"babad"}, "bab"},
		{arg_5{"cbbd"}, "bb"},
		{arg_5{"aaaa"}, "aaaa"},
		{arg_5{"aabaa"}, "aabaa"},
		{arg_5{"bb"}, "bb"},
		{arg_5{"abc"}, "a"},
	}

	for _, c := range cases {
		ans := longestPalindrome(c.input.s)
		assert.Equal(t, c.ans, ans)
	}
}
