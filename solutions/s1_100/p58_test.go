package s1_100

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type case_58 struct {
	input arg_58
	ans   int
}

func Test_58(t *testing.T) {
	cases := []case_58{
		{arg_58{"Hello World"}, 5},
		{arg_58{"   fly me   to   the moon  "}, 4},
		{arg_58{"luffy is still joyboy"}, 6},
	}

	for _, c := range cases {
		ans := lengthOfLastWord(c.input.s)
		assert.Equal(t, c.ans, ans)
	}
}
