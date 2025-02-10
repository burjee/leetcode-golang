package s0

import "log"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i, num := range nums {
		half := target - num
		if j, ok := m[half]; ok {
			return []int{j, i}
		}
		m[num] = i
	}

	panic("unreachable")
}

type args_1 struct {
	nums   []int
	target int
}

func Run_1() {
	input := []args_1{
		{[]int{2, 7, 11, 15}, 9},
		{[]int{2, 7, 11, 15}, 26},
		{[]int{2, 7, 11, 15}, 13},
	}

	for _, arg := range input {
		log.Println(twoSum(arg.nums, arg.target))
	}
}
