package s1_100

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type case_9 struct {
	input arg_9
	ans   bool
}

func Test_9(t *testing.T) {
	cases := []case_9{
		{arg_9{121}, true},
		{arg_9{-121}, false},
		{arg_9{10}, false},
	}

	for _, c := range cases {
		ans := isPalindrome(c.input.x)
		assert.Equal(t, c.ans, ans)
	}
}
