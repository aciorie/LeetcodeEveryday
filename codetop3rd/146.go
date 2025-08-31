package codetop3rd

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

func (this *LRUCache) deleteNode(node *Node) {
	tmp := node.next
	node.prev.next = tmp
	tmp.prev = node.prev
}

func (this *LRUCache) addToHead(node *Node) {
	tmp := this.head.next
	this.head.next = node
	node.prev = this.head
	tmp.prev = node
	node.next = tmp
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.cache[key]; ok {
		this.deleteNode(v)
		this.addToHead(v)
		return v.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if v, ok := this.cache[key]; ok {
		this.deleteNode(v)
		v.val = value
		this.addToHead(v)
	} else {
		node := &Node{key: key, val: value}
		this.cache[key] = node
		this.addToHead(node)
		this.size++

		if this.size > this.capacity {
			tmp := this.tail.prev
			delete(this.cache, tmp.key)
			this.deleteNode(tmp)
			this.size--
		}
	}
}
