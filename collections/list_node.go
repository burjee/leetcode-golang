package collections

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) ToArray() []int {
	arr := []int{}
	ptr := l
	for ptr != nil {
		arr = append(arr, ptr.Val)
		ptr = ptr.Next
	}
	return arr
}

func (l *ListNode) Print() {
	if l != nil {
		fmt.Printf("%v ", l.Val)
		l.Next.Print()
	}
}

func NewListNode(nums []int) *ListNode {
	if nums == nil || len(nums) < 1 {
		return nil
	}

	head := &ListNode{Val: nums[0]}
	ptr := head
	nums = nums[1:]
	for _, num := range nums {
		node := &ListNode{Val: num}
		ptr.Next = node
		ptr = ptr.Next
	}

	return head
}
