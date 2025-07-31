package linkedlists

type Node struct {
	next, prev *Node
	key, val   int
}
type LRUCache struct {
	head, tail     *Node
	size, capacity int
	cache          map[int]*Node
}

func Constructor(capacity int) LRUCache {
	head, tail := Node{}, Node{}
	head.next = &tail
	tail.prev = &head
	return LRUCache{
		head:     &head,
		tail:     &tail,
		capacity: capacity,
		cache:    make(map[int]*Node),
	}
}

func (this *LRUCache) Get(key int) int {
	if pos, exist := this.cache[key]; exist {
		this.removeNode(pos)
		this.addToHead(pos)
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if pos, exist := this.cache[key]; exist {
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
	pos.prev.next = pos.next
	pos.next.prev = pos.prev
}

func (this *LRUCache) addToHead(pos *Node) {
	pos.next = this.head.next
	pos.prev = this.head
	this.head.next = pos
	pos.next.prev = pos
}
