package s1_100

import (
	"fmt"
)

func uniquePaths(m int, n int) int {
	cache := make(map[[2]int]int)
	return helper_62(m, n, &cache)
}

func helper_62(m int, n int, cache *map[[2]int]int) int {
	if m == 1 || n == 1 {
		return 1
	}

	if m == 2 {
		return n
	}

	if n == 2 {
		return m
	}

	if m > n {
		m, n = n, m
	}

	key := [2]int{m, n}
	if value, ok := (*cache)[key]; ok {
		return value
	}

	(*cache)[key] = helper_62(m-1, n, cache) + helper_62(m, n-1, cache)
	return (*cache)[key]
}

type arg_62 struct {
	m int
	n int
}

func Run_62() {
	input := []arg_62{
		{3, 7},
		{3, 2},
	}

	for _, arg := range input {
		fmt.Println(uniquePaths(arg.m, arg.n))
	}
}
