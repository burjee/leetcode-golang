package s1_100

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type case_14 struct {
	input arg_14
	ans   string
}

func Test_14(t *testing.T) {
	cases := []case_14{
		{arg_14{[]string{"flower", "flow", "flight"}}, "fl"},
		{arg_14{[]string{"dog", "racecar", "car"}}, ""},
	}

	for _, c := range cases {
		ans := longestCommonPrefix(c.input.strs)
		assert.Equal(t, c.ans, ans)
	}
}
