package s1_100

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type case_39 struct {
	input arg_39
	ans   [][]int
}

func Test_39(t *testing.T) {
	cases := []case_39{
		{arg_39{[]int{2, 3, 6, 7}, 7}, [][]int{{2, 2, 3}, {7}}},
		{arg_39{[]int{2, 3, 5}, 8}, [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}}},
		{arg_39{[]int{2}, 1}, [][]int{}},
	}

	for _, c := range cases {
		ans := combinationSum(c.input.candidates, c.input.target)
		assert.Equal(t, c.ans, ans)
	}
}
