package collections

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) Print() {
	if t != nil {
		fmt.Printf("%v ", t.Val)
		t.Left.Print()
		t.Right.Print()
	} else {
		fmt.Print("nil ")
	}
}

func NewTreeNone(nums []*int) *TreeNode {
	if nums == nil || len(nums) < 1 {
		return nil
	}

	root := &TreeNode{Val: *nums[0]}
	nodes := []*TreeNode{root}
	nums = nums[1:]

	for i, num := range nums {
		if num != nil {
			node := &TreeNode{Val: *num}
			if i%2 == 0 {
				nodes[0].Left = node
			} else {
				nodes[0].Right = node
				nodes = nodes[1:]
			}
			nodes = append(nodes, node)
		}
	}

	return root
}
