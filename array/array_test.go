package array_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/suzuiyuegjy/go_leetcode/array"
)

var _ = Describe("Array", func() {
	Describe("Product of Array Except Self", func() {
		Context("When array is nil", func() {
			It("product array is []", func() {
				Expect(ProductExceptSelf([]int{})).To(Equal([]int{}))
			})
		})
		Context("when array is normal", func() {
			It("retun product array", func() {
				Expect(ProductExceptSelf([]int{1, 2, 3, 4})).To(Equal([]int{24, 12, 8, 6}))
			})
		})
	})
})
