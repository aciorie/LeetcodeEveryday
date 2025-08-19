package codetop2nd

type node struct {
	key, val   int
	prev, next *node
}

type LRUCache struct {
	cache          map[int]*node
	size, capacity int
	head, tail     *node
}

func Constructor(capacity int) LRUCache {
	head, tail := &node{}, &node{}
	head.next = tail
	tail.prev = head
	return LRUCache{
		head:     head,
		tail:     tail,
		capacity: capacity,
		cache:    make(map[int]*node),
	}
}

func (this *LRUCache) removeNode(node *node) {
	tmp := node.next
	node.prev.next = tmp
	tmp.prev = node.prev
}

func (this *LRUCache) addToHead(node *node) {
	tmp := this.head.next
	this.head.next = node
	node.prev = this.head
	node.next = tmp
	tmp.prev = node
}
func (this *LRUCache) Get(key int) int {
	if v, ok := this.cache[key]; ok {
		this.removeNode(v)
		this.addToHead(v)
		return v.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if v, ok := this.cache[key]; ok {
		this.removeNode(v)
		v.val = value
		this.addToHead(v)
	} else {
		node := &node{key: key, val: value}
		this.cache[key] = node
		this.addToHead(node)
		this.size++
		if this.size > this.capacity {
			nod := this.tail.prev
			delete(this.cache, nod.key)
			this.removeNode(nod)
			this.size--
		}
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
