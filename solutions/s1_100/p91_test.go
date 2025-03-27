package s1_100

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type case_91 struct {
	input arg_91
	ans   int
}

func Test_91(t *testing.T) {
	cases := []case_91{
		{arg_91{"12"}, 2},
		{arg_91{"226"}, 3},
		{arg_91{"06"}, 0},
	}

	for _, c := range cases {
		ans := numDecodings(c.input.s)
		assert.Equal(t, c.ans, ans)
	}
}
