package list

type ListNode struct {
	Val  int
	Next *ListNode
}

func BuildList(s []int) *ListNode {
	var head *ListNode
	var last *ListNode
	for i, v := range s {
		Node := ListNode{v, nil}
		if last != nil {
			last.Next = &Node
		}
		last = &Node
		if i == 0 {
			head = &Node
		}
	}
	return head
}

func ToSlice(head *ListNode) []int {
	var s []int
	for head != nil {
		s = append(s, head.Val)
		head = head.Next
	}
	return s
}

/* in place solution:slow/fast forward
slow:one step
fast:two step
reverse first part of list, then compare with last part
*/
func IsPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	s := make([]int, 0)
	for head != nil {
		s = append(s, head.Val)
		head = head.Next
	}
	l := len(s)
	for i := 0; i <= l/2; i++ {
		if s[i] != s[l-i-1] {
			return false
		}
	}
	return true
}

func IsPalindrome2(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	slow := head
	next := slow.Next
	fast := head.Next
	var last *ListNode
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		last = slow
		slow = next
		next = next.Next
		slow.Next = last
	}
	if fast.Next != nil {
		next = next.Next
	}
	for slow != nil && next != nil {
		if slow.Val != next.Val {
			return false
		}
		slow = slow.Next
		next = next.Next
	}
	return true
}

// param is the node to be delete
// node is not the tail
func deleteNode(node *ListNode) {
	node.Next = nextNode.Next
	return
}
