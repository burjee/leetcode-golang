package s1_100

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type case_6 struct {
	input arg_6
	ans   string
}

func Test_6(t *testing.T) {
	cases := []case_6{
		{arg_6{"PAYPALISHIRING", 3}, "PAHNAPLSIIGYIR"},
		{arg_6{"PAYPALISHIRING", 4}, "PINALSIGYAHRPI"},
		{arg_6{"A", 1}, "A"},
	}

	for _, c := range cases {
		ans := convert(c.input.s, c.input.numRows)
		assert.Equal(t, c.ans, ans)
	}
}
