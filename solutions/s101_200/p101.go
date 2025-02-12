package s101_200

import "fmt"

func isSymmetric(root *TreeNode) bool {
	return helper(root.Left, root.Right)
}

func helper(l *TreeNode, r *TreeNode) bool {
	switch {
	case l != nil && r != nil:
		if l.Val != r.Val {
			return false
		}
		return helper(l.Left, r.Right) && helper(l.Right, r.Left)
	case l != nil && r == nil:
		return false
	case l == nil && r != nil:
		return false
	case l == nil && r == nil:
		return true
	}
	panic("unreachable")
}

type arg_101 struct {
	root *TreeNode
}

func Run_101() {
	input := []arg_101{
		{NewTreeNone([]*int{Ptr(1), Ptr(2), Ptr(2), Ptr(3), Ptr(4), Ptr(4), Ptr(3)})},
		{NewTreeNone([]*int{Ptr(1), Ptr(2), Ptr(2), nil, Ptr(3), nil, Ptr(3)})},
	}

	for _, arg := range input {
		fmt.Println(isSymmetric(arg.root))
	}
}
