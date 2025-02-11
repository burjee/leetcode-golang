package s1_100

import (
	"log"
)

func isSameTree(p *TreeNode, q *TreeNode) bool {
	switch {
	case p != nil && q != nil:
		if p.Val != q.Val {
			return false
		}
		return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	case p != nil && q == nil:
		return false
	case p == nil && q != nil:
		return false
	case p == nil && q == nil:
		return true
	}
	panic("unreachable")
}

type arg_100 struct {
	p *TreeNode
	q *TreeNode
}

func Run_100() {
	input := []arg_100{
		{NewTreeNone([]*int{Ptr(1), Ptr(2), Ptr(3)}),
			NewTreeNone([]*int{Ptr(1), Ptr(2), Ptr(3)})},
		{NewTreeNone([]*int{Ptr(1), Ptr(2)}),
			NewTreeNone([]*int{Ptr(1), nil, Ptr(2)})},
	}

	for _, arg := range input {
		log.Println(isSameTree(arg.p, arg.q))
	}
}
