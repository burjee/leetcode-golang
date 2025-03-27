package s1_100

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type case_98 struct {
	input arg_98
	ans   bool
}

func Test_98(t *testing.T) {
	cases := []case_98{
		{arg_98{NewTreeNone([]*int{Ptr(2), Ptr(1), Ptr(3)})}, true},
		{arg_98{NewTreeNone([]*int{Ptr(5), Ptr(1), Ptr(4), nil, nil, Ptr(3), Ptr(6)})}, false},
	}

	for _, c := range cases {
		ans := isValidBST(c.input.root)
		assert.Equal(t, c.ans, ans)
	}
}
