package reverse

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reverseKGroup(t *testing.T) {
	var head *ListNode

	for i := 5; i > 0; i-- {
		head = &ListNode{Val: i, Next: head}
	}

	head = reverseKGroup(head, 1)
	fmt.Println(head.Val)
	assert.Equal(t, 1, head.Val)
	head = head.Next
	fmt.Println(head.Val)
	assert.Equal(t, 2, head.Val)
	head = head.Next
	fmt.Println(head.Val)
	assert.Equal(t, 3, head.Val)
	head = head.Next
	fmt.Println(head.Val)
	assert.Equal(t, 4, head.Val)
	head = head.Next
	fmt.Println(head.Val)
	assert.Equal(t, 5, head.Val)
}
