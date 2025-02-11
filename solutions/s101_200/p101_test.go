package s101_200

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type case_101 struct {
	input arg_101
	ans   bool
}

func Test_101(t *testing.T) {
	cases := []case_101{
		{arg_101{NewTreeNone([]*int{Ptr(1), Ptr(2), Ptr(2), Ptr(3), Ptr(4), Ptr(4), Ptr(3)})}, true},
		{arg_101{NewTreeNone([]*int{Ptr(1), Ptr(2), Ptr(2), nil, Ptr(3), nil, Ptr(3)})}, false},
	}

	for _, c := range cases {
		ans := isSymmetric(c.input.root)
		assert.Equal(t, c.ans, ans)
	}
}
