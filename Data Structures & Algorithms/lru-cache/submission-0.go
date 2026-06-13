type listNode struct {
	key int
	val int
	prev *listNode
	next *listNode
}

type LRUCache struct {
    internal map[int]*listNode
	maxCap int
	head *listNode
	tail *listNode
}

func Constructor(capacity int) LRUCache {
    internal := make(map[int]*listNode, capacity)
	head, tail := &listNode{}, &listNode{}
	head.prev, head.next = nil, tail
	tail.prev, tail.next = head, nil
	return LRUCache{
		internal: internal,
		maxCap: capacity,
		head: head,
		tail: tail,
	}
}

func (this *LRUCache) Get(key int) int {
    node, ok := this.internal[key]
	if !ok {
		return -1
	}
	RemoveNode(node)
	InsertFront(this, node)
	return node.val
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.internal[key]; ok {
		node.val = value
		RemoveNode(node)
		InsertFront(this, node)
		return
	}
	node := &listNode{key: key, val: value}
	if len(this.internal) >= this.maxCap {
		evicted := this.tail.prev.key
		RemoveNode(this.tail.prev)
		delete(this.internal, evicted)
	}
	InsertFront(this, node)
	this.internal[key] = node
}

func RemoveNode(node *listNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func InsertFront(this *LRUCache, node *listNode) {
	node.next = this.head.next
	node.prev = this.head
	this.head.next.prev = node
	this.head.next = node
}

