/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    length := 0
	cur := head
	for cur != nil {
		cur = cur.Next
		length++ 
	}

	steps := length - n
	if steps == 0 {
		return head.Next
	}

	i := 0
	cur = head
	for cur != nil {
		if i == steps - 1 {
			cur.Next = cur.Next.Next
			break
		}
		cur = cur.Next
		i++ 
	}

	return head
}
