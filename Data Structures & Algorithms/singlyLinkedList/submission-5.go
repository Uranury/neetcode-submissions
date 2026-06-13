type Node struct {
    val int
    next *Node
}

func NewNode(val int) *Node {
    return &Node{val: val}
}

type LinkedList struct {
    head *Node
    size int
}

func NewLinkedList() *LinkedList {
    return &LinkedList{head: nil, size: 0}
}

func (ll *LinkedList) Get(index int) int {
    if index >= ll.size {
        return -1
    }
    node := ll.head
    for i := 0; i < index; i++ {
        node = node.next
    }
    return node.val
}

func (ll *LinkedList) InsertHead(val int) {
    node := NewNode(val)
    node.next = ll.head
    ll.head = node
    ll.size++
}

func (ll *LinkedList) InsertTail(val int) {
    node := NewNode(val)
    if ll.head == nil {
        ll.head = node
        ll.size++
        return
    }
    cur := ll.head
    for cur.next != nil {
        cur = cur.next
    }
    cur.next = node
    ll.size++
}

func (ll *LinkedList) Remove(index int) bool {
    if index >= ll.size {
        return false
    }
    if index == 0 {
        ll.head = ll.head.next
        ll.size--
        return true
    }
    cur := ll.head
    for i := 0; i < index-1; i++ {
        cur = cur.next
    }
    cur.next = cur.next.next
    ll.size--
    return true
}

func (ll *LinkedList) GetValues() []int {
    var arr []int
    cur := ll.head
    for cur != nil {
        arr = append(arr, cur.val)
        cur = cur.next
    }
    return arr
}
