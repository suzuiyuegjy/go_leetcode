package list

import "testing"

func TestBuildList(t *testing.T) {
	Head := buildList([]int{2,3,1,2,3})
	t.Log(Head, to_slice(Head))
}

func TestIsPalindrome(t *testing.T) {
	Node4 := ListNode{1, nil}
	Node3 := ListNode{2, &Node4}
	Node2 := ListNode{2, &Node3}
	Node1 := ListNode{1, &Node2}
	Node := ListNode{1, &Node1}
	if isPalindrome2(&Node) == true {
		t.Error("error")
	}
}
