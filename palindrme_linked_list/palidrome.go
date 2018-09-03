package palindrome

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	var tmp, tmp1, tmp2, tmp3 *ListNode

	if head == nil {
		return true
	}

	tmp1 = nil
	tmp2 = head
	tmp3 = head
	for i := 0; head != nil; i++ {
		head = head.Next
		if i%2 == 0 {
			tmp3 = tmp3.Next
		} else {
			tmp = tmp2
			tmp2 = tmp2.Next
			tmp.Next = tmp1
			tmp1 = tmp
		}
	}

	for tmp1 != nil {
		if tmp1.Val != tmp3.Val {
			return false
		}
		tmp1 = tmp1.Next
		tmp3 = tmp3.Next
	}
	return true
}
