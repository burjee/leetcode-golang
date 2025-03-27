package s1_100

import (
	"fmt"
)

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == right {
		return head
	}

	hair := &ListNode{Val: -1, Next: head}
	head = hair

	for i := 1; i < left; i += 1 {
		head = head.Next
	}

	ptr1 := head.Next
	ptr2 := head.Next.Next
	ptr3 := head.Next.Next.Next
	for left < right-1 {
		ptr2.Next = ptr1
		ptr1 = ptr2
		ptr2 = ptr3
		ptr3 = ptr3.Next
		left += 1
	}

	ptr2.Next = ptr1
	head.Next.Next = ptr3
	head.Next = ptr2

	return hair.Next
}

type arg_92 struct {
	head  *ListNode
	left  int
	right int
}

func Run_92() {
	input := []arg_92{
		{NewListNode([]int{1, 2, 3, 4, 5}), 2, 4},
		{NewListNode([]int{5}), 1, 1},
	}

	for _, arg := range input {
		head := reverseBetween(arg.head, arg.left, arg.right)
		head.Print()
		fmt.Println()
	}
}
