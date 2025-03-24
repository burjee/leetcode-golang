package s1_100

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type case_72 struct {
	input arg_72
	ans   int
}

func Test_72(t *testing.T) {
	cases := []case_72{
		{arg_72{"horse", "ros"}, 3},
		{arg_72{"intention", "execution"}, 5},
		{arg_72{"horse", "rose"}, 2},
		{arg_72{"horse", "horse"}, 0},
		{arg_72{"prosperity", "properties"}, 4},
	}

	for _, c := range cases {
		ans := minDistance(c.input.word1, c.input.word2)
		assert.Equal(t, c.ans, ans)
	}
}
