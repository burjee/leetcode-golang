package s1_100

import (
	"fmt"
	"math"
)

func merge_88(nums1 []int, m int, nums2 []int, n int) {
	i := m + n - 1
	m -= 1
	n -= 1
	for i >= 0 {
		a := math.MinInt32
		b := math.MinInt32
		if m >= 0 {
			a = nums1[m]
		}
		if n >= 0 {
			b = nums2[n]
		}

		if a >= b {
			nums1[i] = a
			m -= 1
		} else {
			nums1[i] = b
			n -= 1
		}
		i -= 1
	}
}

type arg_88 struct {
	nums1 []int
	m     int
	nums2 []int
	n     int
}

func Run_88() {
	input := []arg_88{
		{[]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3},
		{[]int{1}, 1, []int{}, 0},
		{[]int{0}, 0, []int{1}, 1},
	}

	for _, arg := range input {
		merge_88(arg.nums1, arg.m, arg.nums2, arg.n)
		fmt.Println(arg.nums1)
	}
}
