package binary_tree_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/suzuiyuegjy/go_leetcode/binary_tree"
)

func node(val int) *TreeNode {
	return &TreeNode{val, nil, nil}
}

var _ = Describe("BinaryTree", func() {
	var s []string

	BeforeEach(func() {
		s = []string{"3", "5", "1", "6", "2", "0", "8", "nil", "nil", "7", "4"}
	})

	Describe("Build Binary Tree", func() {
		Context("with nil node", func() {
			It("should equal to", func() {
				root := BuildTree(s)
				s1 := ToSlice(root)
				Expect(s).To(Equal(s1[:len(s)]))
			})
		})
	})

	Describe("LowestCommonAncestor", func() {
		Context("when q not p's parnet", func() {
			It("should equal to", func() {
				root := BuildTree(s)
				Expect((LowestCommonAncestor(root, node(5), node(1))).Val).To(Equal(3))
			})
		})
		Context("when q is up's parnet", func() {
			It("should equal to", func() {
				root := BuildTree(s)
				Expect((LowestCommonAncestor(root, node(5), node(4))).Val).To(Equal(5))
			})
		})
	})

})
