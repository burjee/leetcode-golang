package s1_100

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type case_79 struct {
	input arg_79
	ans   bool
}

func Test_79(t *testing.T) {
	cases := []case_79{
		{arg_79{[][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCCED"}, true},
		{arg_79{[][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "SEE"}, true},
		{arg_79{[][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCB"}, false},
		{arg_79{[][]byte{{'a'}}, "a"}, true},
		{arg_79{[][]byte{{'a', 'b'}}, "ab"}, true},
	}

	for _, c := range cases {
		ans := exist(c.input.board, c.input.word)
		assert.Equal(t, c.ans, ans)
	}
}
