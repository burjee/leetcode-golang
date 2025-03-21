package s1_100

import (
	"fmt"
	"strings"
)

func fullJustify(words []string, maxWidth int) []string {
	ans := []string{}

	this_line := []string{}
	line_length := 0
	for _, word := range words {
		if line_length+len(this_line)+len(word) > maxWidth {
			remain_spaces := maxWidth - line_length
			fill_length := max(len(this_line)-1, 1)

			i := 0
			for remain_spaces > 0 {
				this_line[i] += " "
				i = (i + 1) % fill_length
				remain_spaces -= 1
			}

			ans = append(ans, strings.Join(this_line, ""))

			this_line = []string{}
			line_length = 0
		}

		this_line = append(this_line, word)
		line_length += len(word)
	}

	last_line := strings.Join(this_line, " ")
	last_line = last_line + strings.Repeat(" ", maxWidth-len(last_line))
	ans = append(ans, last_line)

	return ans
}

type arg_68 struct {
	words    []string
	maxWidth int
}

func Run_68() {
	input := []arg_68{
		{[]string{"This", "is", "an", "example", "of", "text", "justification."}, 16},
		{[]string{"What", "must", "be", "acknowledgment", "shall", "be"}, 16},
		{[]string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"}, 20},
	}

	for _, arg := range input {
		for _, line := range fullJustify(arg.words, arg.maxWidth) {
			fmt.Printf("\"%v\"\n", line)
		}
		fmt.Println()
	}
}
