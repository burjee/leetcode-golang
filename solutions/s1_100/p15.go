package s1_100

import (
	"fmt"
	"slices"
)

func threeSum(nums []int) [][]int {
	slices.Sort(nums)
	ans := [][]int{}

	i := 0
	for i < len(nums)-2 && nums[i] < 1 {
		l := i + 1
		r := len(nums) - 1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum == 0 {
				ans = append(ans, []int{nums[i], nums[l], nums[r]})
				for l < r {
					l += 1
					if nums[l-1] != nums[l] {
						break
					}
				}
				for l < r && nums[r-1] == nums[r] {
					r -= 1
					if nums[r] != nums[r+1] {
						break
					}
				}
			} else if sum < 0 {
				l += 1
			} else if sum > 0 {
				r -= 1
			}
		}
		for i < len(nums)-2 {
			i += 1
			if nums[i-1] != nums[i] {
				break
			}
		}
	}

	return ans
}

type arg_15 struct {
	nums []int
}

func Run_15() {
	input := []arg_15{
		{[]int{-1, 0, 1, 2, -1, -4}},
		{[]int{0, 1, 1}},
		{[]int{0, 0, 0}},
	}

	for _, arg := range input {
		fmt.Println(threeSum(arg.nums))
	}
}
