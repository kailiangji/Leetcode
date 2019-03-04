package reverse

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	length := func(lst *ListNode) int {
		i := 0
		h := lst
		for ; h != nil; i++ {
			h = h.Next
		}
		return i
	}(head)

	if length < k {
		return head
	}

	tail := head
	tail1 := head
	head1 := head
	head2 := head
	k1 := k

	if k <= length {
		head2 = head2.Next
		head1 = head2
		for ; k1 > 1; k1-- {
			head2 = head2.Next
			head1.Next = tail1
			tail1 = head1
			head1 = head2
		}
		head = tail1
		tail.Next = tail1
		length -= k
	}

	for ; k <= length; length -= k {
		tail0 := head2
		tail1 = head2
		head1 = head2
		k1 = k
		head2 = head2.Next
		head1 = head2
		for ; k1 > 1; k1-- {
			head2 = head2.Next
			head1.Next = tail1
			tail1 = head1
			head1 = head2
		}
		tail.Next = tail1
		tail = tail0
	}

	tail.Next = head2

	return head
}
