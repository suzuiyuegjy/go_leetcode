package binary_tree

import (
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/discuss/65226/My-Java-Solution-which-is-easy-to-understand
another way:
if both p and q exist in Tree rooted at root, then return their LCA
if neither p and q exist in Tree rooted at root, then return null
if only one of p or q (NOT both of them), exists in Tree rooted at root, return it)
*/
func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	pa, qa := findAncestor(root, p, q)
	if pa == qa {
		return pa
	}
	return nil
}

func findAncestor(root, p, q *TreeNode) (pa, qa *TreeNode) {
	if root == nil {
		return nil, nil
	}
	lpa, lqa := findAncestor(root.Left, p, q)
	if lpa != nil && lqa != nil {
		return lpa, lqa
	}

	rpa, rqa := findAncestor(root.Right, p, q)
	if rpa != nil && rqa != nil {
		return rpa, rqa
	}

	if lpa != nil || rpa != nil || root.Val == p.Val {
		pa = root
	}
	if lqa != nil || rqa != nil || root.Val == q.Val {
		qa = root
	}
	return pa, qa
}

func BuildTree(s []string) *TreeNode {
	if len(s) == 0 || s[0] == "nil" {
		return nil
	}
	val, _ := strconv.Atoi(s[0])
	root := &TreeNode{val, nil, nil}
	parents := []*TreeNode{root}
	s = s[1:]
	for i, v := range s {
		var Node *TreeNode
		if v == "nil" {
			Node = nil
		} else {
			val, _ = strconv.Atoi(v)
			Node = &TreeNode{val, nil, nil}
		}
		parents = append(parents, Node)
		p := parents[0]
		if i%2 == 0 {
			p.Left = Node
		} else {
			p.Right = Node
			parents = parents[1:]
		}

	}
	return root
}

func ToSlice(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	parents := []*TreeNode{root}
	val := strconv.Itoa(root.Val)
	s := []string{val}
	for len(parents) != 0 {
		p := parents[0]
		parents = parents[1:]
		if p.Left == nil {
			s = append(s, "nil")
		} else {
			val = strconv.Itoa(p.Left.Val)
			s = append(s, val)
			parents = append(parents, p.Left)
		}
		if p.Right == nil {
			s = append(s, "nil")
		} else {
			val = strconv.Itoa(p.Right.Val)
			s = append(s, val)
			parents = append(parents, p.Right)
		}
	}

	return s
}
