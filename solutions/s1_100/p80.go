package s1_100

import "fmt"

func removeDuplicatesII(nums []int) int {
	if len(nums) < 3 {
		return len(nums)
	}

	l := 2
	r := 2
	for r < len(nums) {
		nums[l] = nums[r]
		if nums[l-2] != nums[r] {
			l += 1
		}
		r += 1
	}
	return l
}

type arg_80 struct {
	nums []int
}

func Run_80() {
	input := []arg_80{
		{[]int{1, 1, 1, 2, 2, 3}},
		{[]int{0, 0, 1, 1, 1, 1, 2, 3, 3}},
		{[]int{0, 0, 1, 1, 2, 3, 3}},
		{[]int{1}},
	}

	for _, arg := range input {
		k := removeDuplicatesII(arg.nums)
		fmt.Printf("k: %v, nums: ", k)
		for i := 0; i < k; i += 1 {
			fmt.Printf("%v ", arg.nums[i])
		}
		fmt.Println()
	}
}
