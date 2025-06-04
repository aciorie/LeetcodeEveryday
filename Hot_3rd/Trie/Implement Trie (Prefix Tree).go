package trie

type Trie struct {
	children map[rune]*Trie
	isEnd    bool
}

func Constructor() Trie {
	return Trie{
		children: map[rune]*Trie{},
		// 根节点本身不代表任何单词的结束
		isEnd: false,
	}
}

func (this *Trie) Insert(word string) {
	// 从当前节点（通常是根节点）开始遍历
	cur := this

	for _, char := range word {
		// 检查当前节点的 children map 中是否已经有这个字符对应的子节点
		if _, ok := cur.children[char]; !ok {
			// 如果没有，就创建一个新的 Trie 节点作为子节点
			cur.children[char] = &Trie{
				children: make(map[rune]*Trie),
				isEnd:    false,
			}
		}
		cur = cur.children[char]
	}
	cur.isEnd = true
}

func (this *Trie) Search(word string) bool {
	cur := this

	for _, char := range word {
		if _, ok := cur.children[char]; !ok {
			return false
		}
		cur = cur.children[char]
	}

	// 搜完整单词的，必须原模原样
	return cur.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	cur := this

	for _, char := range prefix {
		if _, ok := cur.children[char]; !ok {
			return false
		}
		cur = cur.children[char]
	}

	// 搜前缀的，至少要有前缀，后面还有没有东西无所谓
	return true
}
