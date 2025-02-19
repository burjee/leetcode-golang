package s1_100

import (
	"fmt"
)

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	hair := &ListNode{Val: -1, Next: head}
	left := hair
	right := hair

	for i := 0; i < n; i += 1 {
		right = right.Next
	}

	for right.Next != nil {
		left = left.Next
		right = right.Next
	}

	left.Next = left.Next.Next

	return hair.Next
}

type arg_19 struct {
	head *ListNode
	n    int
}

func Run_19() {
	input := []arg_19{
		{NewListNode([]int{1, 2, 3, 4, 5}), 2},
		{NewListNode([]int{1}), 1},
		{NewListNode([]int{1, 2}), 1},
	}

	for _, arg := range input {
		head := removeNthFromEnd(arg.head, arg.n)
		head.Print()
		fmt.Println()
	}
}
