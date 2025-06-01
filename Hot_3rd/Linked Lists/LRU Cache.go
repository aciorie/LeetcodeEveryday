package main

type LRUCache struct {
	cache    map[int]*node
	capacity int
	list     *doublyLinkedList
}

type node struct {
	key   int
	value int
	prev  *node
	next  *node
}

type doublyLinkedList struct {
	head *node
	tail *node
}

func Constructor(capacity int) LRUCache {
	list := &doublyLinkedList{
		head: &node{},
		tail: &node{},
	}
	list.head.next = list.tail
	list.tail.prev = list.head

	return LRUCache{
		cache:    make(map[int]*node),
		capacity: capacity,
		list:     list}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.cache[key]; ok {
		this.moveToHead(node)
		return node.value
	}
	return -1
}

func (this *LRUCache) moveToHead(node *node) {
	// 从当前位置移除节点
	this.removeNode(node)
	// 将节点添加到头部
	this.addNodeToHead(node)
}

func (this *LRUCache) removeNode(node *node) {
	prev := node.prev
	next := node.next
	prev.next = next
	next.prev = prev
}

func (this *LRUCache) addNodeToHead(node *node) {
	// this.list.head 是哨兵头节点，它本身不存储实际的缓存数据
	node.prev = this.list.head
	node.next = this.list.head.next
	this.list.head.next.prev = node
	this.list.head.next = node
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.cache[key]; ok {
		node.value = value
		this.moveToHead(node)
		return
	}

	newNode := &node{key: key, value: value}
	this.addNodeToHead(newNode)
	this.cache[key] = newNode

	if len(this.cache) > this.capacity {
		// 移除尾部节点
		tail := this.removeTail()
		if tail != nil {
			// 从 cache 中移除对应的 key
			delete(this.cache, tail.key)
		}
	}
}

func (this *LRUCache) removeTail() *node {
	if this.list.tail.prev == this.list.head {
		return nil
	}
	tail := this.list.tail
	this.removeNode(tail)
	return tail
}
