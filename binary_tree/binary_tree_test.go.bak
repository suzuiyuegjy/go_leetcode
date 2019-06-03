package binary_tree

import "testing"

func TestBuildTree(t *testing.T) {
	root := buildTree([]string{"3", "5", "1", "6", "2", "0", "8", "nil", "nil", "7", "4"})
	t.Log(toSlice(root))
}

func node(val int) *TreeNode {
	return &TreeNode{val, nil, nil}
}

func TestLowestCommonAncestor(t *testing.T) {
	root := buildTree([]string{"3", "5", "1", "6", "2", "0", "8", "nil", "nil", "7", "4"})
	r := lowestCommonAncestor(root, node(5), node(1))
	if r.Val != 3 {
		t.Error(r)
	}
	r = lowestCommonAncestor(root, node(5), node(4))
	if r.Val != 5 {
		t.Error(r)
	}
}
