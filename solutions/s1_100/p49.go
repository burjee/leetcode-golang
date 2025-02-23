package s1_100

import (
	"fmt"
	"slices"
)

func groupAnagrams(strs []string) [][]string {
	ans := [][]string{}
	m := make(map[[26]rune]*[]string)

	for _, str := range strs {
		count := [26]rune{}
		for _, ch := range str {
			count[ch-97] += 1
		}
		if v, ok := m[count]; ok {
			*v = append(*v, str)
		} else {
			m[count] = &[]string{str}
		}
	}

	for _, v := range m {
		ans = append(ans, *v)
	}

	return ans
}

type arg_49 struct {
	strs []string
}

func Run_49() {
	input := []arg_49{
		{[]string{"eat", "tea", "tan", "ate", "nat", "bat"}},
		{[]string{""}},
		{[]string{"a"}},
	}

	for _, arg := range input {
		ans := groupAnagrams(arg.strs)
		for _, strs := range ans {
			slices.Sort(strs)
		}
		slices.SortFunc(ans, func(a, b []string) int { return len(a) - len(b) })
		fmt.Printf("%v\n", ans)
	}
}
