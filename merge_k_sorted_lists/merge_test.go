package merge

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_mergeKLists(t *testing.T) {
	lists := make([]*ListNode, 3)

	lists[0] = &ListNode{5, nil}
	lists[0] = &ListNode{4, lists[0]}
	lists[0] = &ListNode{1, lists[0]}

	lists[1] = &ListNode{4, nil}
	lists[1] = &ListNode{3, lists[1]}
	lists[1] = &ListNode{1, lists[1]}

	lists[2] = &ListNode{6, nil}
	lists[2] = &ListNode{2, lists[2]}

	l := mergeKLists(lists)
	assert.Equal(t, 1, l.Val)
	l = l.Next
	assert.Equal(t, 1, l.Val)
	l = l.Next
	assert.Equal(t, 2, l.Val)
	l = l.Next
	assert.Equal(t, 3, l.Val)
	l = l.Next
	assert.Equal(t, 4, l.Val)
	l = l.Next
	assert.Equal(t, 4, l.Val)
	l = l.Next
	assert.Equal(t, 5, l.Val)
	l = l.Next
	assert.Equal(t, 6, l.Val)
	l = l.Next
	assert.Nil(t, l)
	//1->1->2->3->4->4->5->6
}
