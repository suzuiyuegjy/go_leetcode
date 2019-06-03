package binary_tree_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestBinaryTree(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "BinaryTree Suite")
}
