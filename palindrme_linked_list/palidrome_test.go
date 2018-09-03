package palindrome

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isPalindrome(t *testing.T) {
	var head *ListNode

	/*l1_1 := &ListNode{Val: 1}
	l1_2 := &ListNode{Val: 2}
	l1_1.Next = l1_2
	head = l1_1
	assert.False(t, isPalindrome(head))

	l2_1 := &ListNode{Val: 1}
	l2_2 := &ListNode{Val: 2}
	l2_3 := &ListNode{Val: 2}
	l2_4 := &ListNode{Val: 1}
	l2_1.Next = l2_2
	l2_2.Next = l2_3
	l2_3.Next = l2_4
	head = l2_1
	assert.True(t, isPalindrome(head))*/

	l3_1 := &ListNode{Val: 1}
	l3_2 := &ListNode{Val: 2}
	l3_3 := &ListNode{Val: 5}
	l3_1.Next = l3_2
	l3_2.Next = l3_3
	head = l3_1
	assert.False(t, isPalindrome(head))
}
