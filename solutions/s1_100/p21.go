package s1_100

import (
	"fmt"
)

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	hair := &ListNode{Val: -1}
	ptr := hair
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			ptr.Next = list1
			list1 = list1.Next
		} else {
			ptr.Next = list2
			list2 = list2.Next
		}
		ptr = ptr.Next
	}

	if list1 != nil {
		ptr.Next = list1
	}

	if list2 != nil {
		ptr.Next = list2
	}

	return hair.Next
}

type arg_21 struct {
	list1 *ListNode
	list2 *ListNode
}

func Run_21() {
	input := []arg_21{
		{NewListNode([]int{1, 2, 4}), NewListNode([]int{1, 3, 4})},
		{NewListNode([]int{}), NewListNode([]int{})},
		{NewListNode([]int{}), NewListNode([]int{0})},
	}

	for _, arg := range input {
		head := mergeTwoLists(arg.list1, arg.list2)
		head.Print()
		fmt.Println()
	}
}
