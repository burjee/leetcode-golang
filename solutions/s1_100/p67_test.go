package s1_100

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type case_67 struct {
	input arg_67
	ans   string
}

func Test_67(t *testing.T) {
	cases := []case_67{
		{arg_67{"11", "1"}, "100"},
		{arg_67{"1010", "1011"}, "10101"},
	}

	for _, c := range cases {
		ans := addBinary(c.input.a, c.input.b)
		assert.Equal(t, c.ans, ans)
	}
}
