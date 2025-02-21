package s1_100

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type case_28 struct {
	input arg_28
	ans   int
}

func Test_28(t *testing.T) {
	cases := []case_28{
		{arg_28{"sadbutsad", "sad"}, 0},
		{arg_28{"leetcode", "leeto"}, -1},
	}

	for _, c := range cases {
		ans := strStr(c.input.haystack, c.input.needle)
		assert.Equal(t, c.ans, ans)
	}
}
