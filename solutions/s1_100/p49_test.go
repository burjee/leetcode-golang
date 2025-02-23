package s1_100

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

type case_49 struct {
	input arg_49
	ans   [][]string
}

func Test_49(t *testing.T) {
	cases := []case_49{
		{arg_49{[]string{"eat", "tea", "tan", "ate", "nat", "bat"}}, [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}}},
		{arg_49{[]string{""}}, [][]string{{""}}},
		{arg_49{[]string{"a"}}, [][]string{{"a"}}},
	}

	for _, c := range cases {
		ans := groupAnagrams(c.input.strs)
		for _, strs := range ans {
			slices.Sort(strs)
		}
		slices.SortFunc(ans, func(a, b []string) int { return len(a) - len(b) })
		assert.Equal(t, c.ans, ans)
	}
}
