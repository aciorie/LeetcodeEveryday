package codetop4th

type Node struct {
	key, val   int
	prev, next *Node
}

type LRUCache struct {
	head, tail     *Node
	size, capacity int
	cache          map[int]*Node
}

func Constructor(capacity int) LRUCache {
	head, tail := &Node{}, &Node{}
	head.next = tail
	tail.prev = head
	return LRUCache{
		head:     head,
		tail:     tail,
		capacity: capacity,
		cache:    make(map[int]*Node),
	}
}

func (this *LRUCache) RemoveNode(node *Node) {
	tmp := node.next
	node.prev.next = tmp
	tmp.prev = node.prev
}

func (this *LRUCache) AddToHead(node *Node) {
	tmp := this.head.next
	node.next = tmp
	node.prev = this.head
	this.head.next = node
	tmp.prev = node
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.cache[key]; ok {
		this.RemoveNode(v)
		this.AddToHead(v)
		return v.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if v, ok := this.cache[key]; ok {
		this.RemoveNode(v)
		v.val = value
		this.AddToHead(v)
	} else {
		node := &Node{key: key, val: value}
		this.cache[key] = node
		this.AddToHead(node)
		this.size++

		if this.size > this.capacity {
			tmp := this.tail.prev
			delete(this.cache, tmp.key)
			this.RemoveNode(tmp)
			this.size--
		}
	}
}
