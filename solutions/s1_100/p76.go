package s1_100

import "fmt"

func minWindow(s string, t string) string {
	total_count := make([]int, 52)
	window_count := make([]int, 52)
	for i := range t {
		char_index := ToZeroIndex(t[i])
		total_count[char_index] += 1
	}

	ans := ""
	start := 0
	window_found := 0
	for end := range s {
		end_index := ToZeroIndex(s[end])

		if total_count[end_index] > 0 {
			window_count[end_index] += 1
			if total_count[end_index] >= window_count[end_index] {
				window_found += 1
			}

			start_index := ToZeroIndex(s[start])
			for total_count[start_index] == 0 || total_count[start_index] < window_count[start_index] {
				if total_count[start_index] < window_count[start_index] {
					window_count[start_index] -= 1
				}
				start += 1
				start_index = ToZeroIndex(s[start])
			}

			if window_found == len(t) {
				if ans == "" || end-start+1 < len(ans) {
					ans = s[start : end+1]
				}
			}
		}
	}

	return ans
}

type arg_76 struct {
	s string
	t string
}

func Run_76() {
	input := []arg_76{
		{"ADOBECODEBANC", "ABC"},
		{"a", "a"},
		{"a", "aa"},
		{"ab", "A"},
	}

	for _, arg := range input {
		fmt.Println(minWindow(arg.s, arg.t))
	}
}
