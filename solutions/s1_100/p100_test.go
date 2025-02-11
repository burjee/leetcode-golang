package s1_100

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type case_100 struct {
	input arg_100
	ans   bool
}

func Test_100(t *testing.T) {
	cases := []case_100{
		{arg_100{NewTreeNone([]*int{Ptr(1), Ptr(2), Ptr(3)}),
			NewTreeNone([]*int{Ptr(1), Ptr(2), Ptr(3)})}, true},
		{arg_100{NewTreeNone([]*int{Ptr(1), Ptr(2)}),
			NewTreeNone([]*int{Ptr(1), nil, Ptr(2)})}, false},
	}

	for _, c := range cases {
		ans := isSameTree(c.input.p, c.input.q)
		assert.Equal(t, c.ans, ans)
	}
}
