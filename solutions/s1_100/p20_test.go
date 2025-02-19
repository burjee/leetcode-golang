package s1_100

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type case_20 struct {
	input arg_20
	ans   bool
}

func Test_20(t *testing.T) {
	cases := []case_20{
		{arg_20{"()"}, true},
		{arg_20{"()[]{}"}, true},
		{arg_20{"(]"}, false},
		{arg_20{"([])"}, true},
		{arg_20{"("}, false},
	}

	for _, c := range cases {
		ans := isValid(c.input.s)
		assert.Equal(t, c.ans, ans)
	}
}
