package s1_100

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type case_68 struct {
	input arg_68
	ans   []string
}

func Test_68(t *testing.T) {
	cases := []case_68{
		{arg_68{[]string{"This", "is", "an", "example", "of", "text", "justification."}, 16}, []string{"This    is    an", "example  of text", "justification.  "}},
		{arg_68{[]string{"What", "must", "be", "acknowledgment", "shall", "be"}, 16}, []string{"What   must   be", "acknowledgment  ", "shall be        "}},
		{arg_68{[]string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"}, 20}, []string{"Science  is  what we", "understand      well", "enough to explain to", "a  computer.  Art is", "everything  else  we", "do                  "}},
	}

	for _, c := range cases {
		ans := fullJustify(c.input.words, c.input.maxWidth)
		assert.Equal(t, c.ans, ans)
	}
}
