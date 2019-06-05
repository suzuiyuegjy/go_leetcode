package list_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/suzuiyuegjy/go_leetcode/list"
)

var _ = Describe("List", func() {
	Describe("Build List", func(){
		Context("normal", func(){
			It("should be", func(){
				s := []int{2, 3, 1, 2, 3}
				L := BuildList(s)
				Expect(ToSlice(L)).To(Equal(s))
			})
		})
	})

	Describe("Is List Palindrome", func(){
		Context("When List is Normal", func(){
			It("it is palindrome", func(){
				s := []int{1,2,3,2,1}
				L := BuildList(s)
				Expect(IsPalindrome2(L)).To(Equal(true))
			})
		})
	})
})
