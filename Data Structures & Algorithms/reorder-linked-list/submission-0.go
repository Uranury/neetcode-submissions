/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	second := slow.Next
	slow.Next = nil

	var prev *ListNode = nil
	for second != nil {
		tmp := second.Next
		second.Next = prev
		prev = second
		second = tmp
	}

	first, second := head, prev
	for second != nil {
		tmp1, tmp2 := first.Next, second.Next
		first.Next = second
		second.Next = tmp1
		first = tmp1
		second = tmp2
	}
}
