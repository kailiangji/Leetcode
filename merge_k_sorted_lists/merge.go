package merge

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	var head, curr, tmp *ListNode
	var currIndex int
	length := len(lists)
	if length == 0 {
		return nil
	}

	for i, node := range lists {
		if node == nil {
			continue
		}
		if head == nil {
			head = node
			currIndex = i
			continue
		}
		if node.Val < head.Val {
			head = node
			currIndex = i
		}
	}

	curr = head

	if curr == nil {
		return nil
	}

	lists[currIndex] = lists[currIndex].Next

	for curr != nil {
		tmp = nil
		for i, node := range lists {
			if node == nil {
				continue
			}

			if tmp != nil {
				if node.Val < tmp.Val {
					tmp = node
					currIndex = i
				}
				continue
			}

			tmp = node
			currIndex = i
		}
		curr.Next = tmp
		curr = tmp

		if lists[currIndex] != nil {
			lists[currIndex] = lists[currIndex].Next
		}
	}

	return head
}
