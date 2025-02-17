package s1_100

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type case_12 struct {
	input arg_12
	ans   string
}

func Test_12(t *testing.T) {
	cases := []case_12{
		{arg_12{3749}, "MMMDCCXLIX"},
		{arg_12{58}, "LVIII"},
		{arg_12{1994}, "MCMXCIV"},
	}

	for _, c := range cases {
		ans := intToRoman(c.input.num)
		assert.Equal(t, c.ans, ans)
	}
}
