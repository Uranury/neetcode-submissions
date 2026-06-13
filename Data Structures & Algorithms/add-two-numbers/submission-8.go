/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    length1, length2 := 0, 0 
	var longer, shorter *ListNode
	i, j := l1, l2
	for i != nil {
		length1++
		i = i.Next
	}
	for j != nil {
		length2++
		j = j.Next
	}

	var head *ListNode
	if length1 >= length2 {
		head, longer = l1, l1
		shorter = l2
	} else {
		head, longer = l2, l2
		shorter = l1
	}

	carry := 0
	var prev *ListNode
	for longer != nil || carry != 0 {
		if longer == nil {
			prev.Next = &ListNode{Val: carry}
			break
		}
		shortVal := 0
		if shorter != nil {
			shortVal = shorter.Val
			shorter = shorter.Next
		}

		value := longer.Val + shortVal + carry
		longer.Val = value % 10
		carry = value / 10

		prev = longer
		longer = longer.Next
	}

	return head
}
