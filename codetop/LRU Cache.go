package codetop

type LRUCache struct {
	size, capacity int
	head, tail     *Node
	cache          map[int]*Node
}

type Node struct {
	key, val   int
	next, prev *Node
}

func Constructor(capacity int) LRUCache {
	head := Node{}
	tail := Node{}
	head.next = &tail
	tail.prev = &head
	return LRUCache{
		capacity: capacity,
		head:     &head,
		tail:     &tail,
		cache:    make(map[int]*Node),
	}
}

func (this *LRUCache) Get(key int) int {
	if pos, ok := this.cache[key]; ok {
		this.removeNode(pos)
		this.addToHead(pos)
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if pos, ok := this.cache[key]; ok {
		this.removeNode(pos)
		pos.val = value
		this.addToHead(pos)
	} else {
		node := &Node{key: key, val: value}
		this.cache[key] = node
		this.addToHead(node)
		
		if this.size++; this.size > this.capacity {
			node = this.tail.prev
			delete(this.cache, node.key)
			this.removeNode(node)
			this.size--
		}
	}
}

func (this *LRUCache) removeNode(pos *Node) {
	pre, nex := pos.prev, pos.next
	pre.next = nex
	nex.prev = pre
}

func (this *LRUCache) addToHead(pos *Node) {
	pos.next = this.head.next
	pos.prev = this.head
	this.head.next = pos
	pos.next.prev = pos
}
