package s1_100

import (
	"fmt"
)

func isValidBST(root *TreeNode) bool {
	return helper_98(root, nil, nil)
}

func helper_98(root *TreeNode, min, max *int) bool {
	if root == nil {
		return true
	}

	if min != nil && *min >= root.Val {
		return false
	}

	if max != nil && root.Val >= *max {
		return false
	}

	return helper_98(root.Left, min, &root.Val) && helper_98(root.Right, &root.Val, max)
}

type arg_98 struct {
	root *TreeNode
}

func Run_98() {
	input := []arg_98{
		{NewTreeNone([]*int{Ptr(2), Ptr(1), Ptr(3)})},
		{NewTreeNone([]*int{Ptr(5), Ptr(1), Ptr(4), nil, nil, Ptr(3), Ptr(6)})},
	}

	for _, arg := range input {
		fmt.Println(isValidBST(arg.root))
	}
}
