package s1_100

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type case_3 struct {
	input arg_3
	ans   int
}

func Test_3(t *testing.T) {
	cases := []case_3{
		{arg_3{"abcabcbb"}, 3},
		{arg_3{"bbbbb"}, 1},
		{arg_3{"pwwkew"}, 3},
	}

	for _, c := range cases {
		ans := lengthOfLongestSubstring(c.input.s)
		assert.Equal(t, c.ans, ans)
	}
}
