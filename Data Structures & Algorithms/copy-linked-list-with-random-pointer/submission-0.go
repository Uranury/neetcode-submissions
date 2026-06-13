/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	cur := head
	pointers := make(map[*Node]*Node)

	for cur != nil {
		pointers[cur] = &Node{Val: cur.Val}
		cur = cur.Next
	}

	cur = head
	for cur != nil {
		if cur.Next != nil {
			pointers[cur].Next = pointers[cur.Next]
		}
		
		pointers[cur].Random = pointers[cur.Random]
		cur = cur.Next
	}

	return pointers[head]
}
