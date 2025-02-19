package s1_100

import (
	"fmt"
)

func swapPairs(head *ListNode) *ListNode {
	if head != nil {
		if head.Next != nil {
			a := head
			b := head.Next
			a.Next = swapPairs(b.Next)
			b.Next = a
			return b
		}
		return head
	}
	return nil
}

// func swapPairs(head *ListNode) *ListNode {
// 	hair := &ListNode{Val: -1, Next: head}
// 	left := hair
// 	right := hair.Next
// 	for right != nil && right.Next != nil {
// 		left.Next = right.Next
// 		right.Next = left.Next.Next
// 		left.Next.Next = right

// 		left = right
// 		right = right.Next
// 	}
// 	return hair.Next
// }

type arg_24 struct {
	head *ListNode
}

func Run_24() {
	input := []arg_24{
		{NewListNode([]int{1, 2, 3, 4})},
		{NewListNode([]int{})},
		{NewListNode([]int{1})},
		{NewListNode([]int{1, 2, 3})},
	}

	for _, arg := range input {
		head := swapPairs(arg.head)
		head.Print()
		fmt.Println()
	}
}
