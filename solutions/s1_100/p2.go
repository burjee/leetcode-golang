package s1_100

import "fmt"

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	hold := 0
	hair := &ListNode{Val: -1}
	ptr := hair

	for l1 != nil || l2 != nil || hold > 0 {
		if l1 != nil {
			hold += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			hold += l2.Val
			l2 = l2.Next
		}
		node := &ListNode{Val: hold % 10}
		hold = hold / 10

		ptr.Next = node
		ptr = ptr.Next
	}

	return hair.Next
}

type arg_2 struct {
	l1 *ListNode
	l2 *ListNode
}

func Run_2() {
	input := []arg_2{
		{NewListNode([]int{2, 4, 3}), NewListNode([]int{5, 6, 4})},
		{NewListNode([]int{0}), NewListNode([]int{0})},
		{NewListNode([]int{9, 9, 9, 9, 9, 9, 9}), NewListNode([]int{9, 9, 9, 9})},
	}

	for _, arg := range input {
		head := addTwoNumbers(arg.l1, arg.l2)
		head.Print()
		fmt.Println()
	}
}
