package reverse

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reverseKGroup(t *testing.T) {
	var head *ListNode

	for i := 2; i > 0; i-- {
		head = &ListNode{Val: i, Next: head}
	}

	head2 := reverseKGroup(head, 3)
	fmt.Println(head2.Val)
	assert.Equal(t, 1, head2.Val)
	head2 = head2.Next
	fmt.Println(head2.Val)
	assert.Equal(t, 2, head2.Val)
}
