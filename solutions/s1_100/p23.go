package s1_100

import (
	"fmt"
	"slices"
)

func mergeKLists(lists []*ListNode) *ListNode {
	nums := []int{}
	for _, list := range lists {
		ptr := list
		for ptr != nil {
			nums = append(nums, ptr.Val)
			ptr = ptr.Next
		}
	}
	slices.Sort(nums)

	hair := &ListNode{Val: -1}
	ptr := hair
	for _, num := range nums {
		node := &ListNode{Val: num}
		ptr.Next = node
		ptr = ptr.Next
	}

	return hair.Next
}

type arg_23 struct {
	lists []*ListNode
}

func Run_23() {
	input := []arg_23{
		{[]*ListNode{NewListNode([]int{1, 4, 5}), NewListNode([]int{1, 3, 4}), NewListNode([]int{2, 6})}},
		{[]*ListNode{}},
		{[]*ListNode{NewListNode([]int{})}},
	}

	for _, arg := range input {
		head := mergeKLists(arg.lists)
		head.Print()
		fmt.Println()
	}
}
